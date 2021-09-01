import { useState, useCallback } from 'react';

export const useRequest = (initData) => {
  const [data, setData] = useState(initData);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const run = useCallback((promise) => {
    setLoading(true);
    setError('');

    return promise
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

  return [loading, error, run, data];
}

export const useFields = (initFields) => {
  const [fields, setFields] = useState(initFields)
  const handleChange = useCallback((e, {name, type, value, checked}) => {
      setFields(f => {
        f = {...f, [name]: type === 'checkbox' ? checked : value };
        return f
      });
    }, [setFields])

  return [fields, handleChange, setFields];
}