import { useState, useCallback } from 'react';
import API from 'Api'

export const useAPI = (init) => {
  const [data, setData] = useState(init);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const run = useCallback((promiseFn, args) => {
    setLoading(true);
    setError('');

    return promiseFn(args)
      .then(data => {
        setData(data);
        setLoading(false);
        return Promise.resolve(data);
      })
      .catch(error => {
        setError(error);
        setLoading(false);
        return Promise.reject(error);
      });
  }, [])

  return [data, loading, error, run, API];
}

export const useFields = (init) => {
  const [fields, setFields] = useState(init)
  const handleChange = useCallback((e, {name, type, value, checked}) => {
      setFields(f => {
        f = {...f, [name]: type === 'checkbox' ? checked : value };
        return f
      });
    }, [setFields])

  return [fields, setFields, handleChange];
}
