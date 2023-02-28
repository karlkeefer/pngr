import { useState, useCallback, ChangeEvent, useRef, useEffect, SyntheticEvent, FormEvent } from 'react';

import _ from 'lodash'
import { CheckboxProps, DropdownProps, InputOnChangeData, TextAreaProps } from 'semantic-ui-react';

export type RunFunc<T> = (promise: Promise<any>, onSuccess?: (data: T) => void, onFailure?: Function) => void

export const useRequest = <T extends Object>(initData: T): [boolean, string, RunFunc<T>, T] => {
  const [data, setData] = useState(initData);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  // prevent setters if component is unmounted
  // this can happen if e.g. onSuccess callback causes page navigation
  const isMountedRef = useRef(true);
  useEffect(() => {
    return () => void (isMountedRef.current = false);
  }, []);

  // we could just return the promise from run(), but using onSuccess and onFailure callbacks 
  // allows us to react before the loading/errors states change - this is mostly useful if 
  // we want to redirect before the not-loading layout can show itself
  const run = useCallback((promise: Promise<any>, onSuccess?: (data: T) => void, onFailure?: Function) => {
    setLoading(true);
    setError('');

    return promise
      .then(data => {
        if (onSuccess) {
          onSuccess(data);
        }
        if (isMountedRef.current) {
          setData(data);
          setLoading(false);
        }
      })
      .catch(error => {
        if (onFailure) {
          onFailure(error);
        }
        if (isMountedRef.current) {
          setError(error);
          setLoading(false);
        }
      });
  }, [])

  return [loading, error, run, data];
}

// useFields gives us a simple handleChange that works for most form inputs.
// this hook also supports nested properties!
// You just have to set your input field's name attr appropriately
// e.g. w/ a schema like {person:{first_name:''}} you can do <input name="person.first_name"/>
type Evt = SyntheticEvent<HTMLElement, Event> | ChangeEvent<HTMLInputElement | HTMLTextAreaElement> | FormEvent<HTMLInputElement> | null
type Data = InputOnChangeData | TextAreaProps | CheckboxProps | DropdownProps

export type ChangeHandler = (event: Evt, data: Data) => void


type ChangeData = {
  name: string
  type: string
  value: string
  checked?: boolean
}

export const useFields = <T extends Object>(initFields: T): {fields: T, handleChange: ChangeHandler, setFields: (newFieldValues: T) => void} => {
  const [fields, setFields] = useState(initFields)

  // changeHandler works for <Input> and <TextArea>, but the onChange field for semantic-ui form components 
  // has different type signatures so we have to create multiple handlers
  const changeHandler = useCallback(({ name, type, value, checked }: ChangeData)=>{
    setFields(f => {
      let out = _.cloneDeep(f)
      _.set(out, name, type === 'checkbox' ? checked : value);
      return out
    });
  }, [setFields])

  const handleChange = useCallback((e: Evt, cd: Data) => {
    changeHandler(cd as ChangeData)
  }, [changeHandler])

  return {fields, handleChange, setFields};
}
