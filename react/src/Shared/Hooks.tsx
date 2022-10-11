import { useState, useCallback, ChangeEvent } from 'react';

import _ from 'lodash'
import { InputOnChangeData, TextAreaProps } from 'semantic-ui-react';

export type RunFunc<T> = (promise: Promise<any>, onSuccess?: (data: T) => void, onFailure?: Function) => void

export const useRequest = <T extends Object>(initData: T): [boolean, string, RunFunc<T>, T] => {
  const [data, setData] = useState(initData);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

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
        setData(data);
        setLoading(false);
      })
      .catch(error => {
        if (onFailure) {
          onFailure(error);
        }
        setError(error);
        setLoading(false);
      });
  }, [])

  return [loading, error, run, data];
}

// useFields gives us a simple handleChange that works for most form inputs.
// this hook also supports nested properties!
// You just have to set your input field's name attr appropriately
// e.g. w/ a schema like {person:{first_name:''}} you can do <input name="person.first_name"/>


export type InputChangeHandler = (event: ChangeEvent<HTMLInputElement>, data: InputOnChangeData) => void
export type TextAreaChangeHandler = (event: ChangeEvent<HTMLTextAreaElement>, data: TextAreaProps) => void

type ChangeData = {
  name: string
  type: string
  value: string
  checked?: boolean
}

export const useFields = <T extends Object>(initFields: T): {fields: T, handleInputChange: InputChangeHandler, handleTextAreaChange: TextAreaChangeHandler, setFields: (newFieldValues: T) => void} => {
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

  const handleInputChange = useCallback((e: ChangeEvent<HTMLInputElement>, cd: InputOnChangeData) => {
    changeHandler(cd as ChangeData)
  }, [changeHandler])

  const handleTextAreaChange = useCallback((e: ChangeEvent<HTMLTextAreaElement>, cd: TextAreaProps) => {
    changeHandler(cd as ChangeData)
  }, [changeHandler])

  return {fields, handleInputChange, handleTextAreaChange, setFields};
}
