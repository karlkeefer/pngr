import { useState, useCallback, ChangeEvent } from 'react';
import _ from 'lodash'
import { InputOnChangeData, TextAreaProps } from 'semantic-ui-react';

export const useRequest = (initData: Object | undefined): [boolean, string, Function, any] => {
  const [data, setData] = useState(initData);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  // we could just return the promise from run(), but using onSuccess and onFailure callbacks 
  // allows us to react before the loading/errors states change - this is mostly useful if 
  // we want to redirect before the not-loading layout can show itself
  const run = useCallback((promise: Promise<any>, onSuccess: Function, onFailure: Function) => {
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

export type InputChangeHandler = (event: React.ChangeEvent<HTMLInputElement>, data: InputOnChangeData) => void
export type TextAreaChangeHandler = (event: React.ChangeEvent<HTMLTextAreaElement>, data: TextAreaProps) => void

export const useFields = (initFields: Object): [Object, InputChangeHandler | TextAreaChangeHandler, Function] => {
  const [fields, setFields] = useState(initFields)
  const handleChange = useCallback((e: ChangeEvent<HTMLInputElement>, { name, type, value, checked }: InputOnChangeData) => {
    setFields(f => {
      let out = _.cloneDeep(f)
      _.set(out, name, type === 'checkbox' ? checked : value);
      return out
    });
  }, [setFields])

  return [fields, handleChange, setFields];
}
