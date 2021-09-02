import React, { useState, useEffect, useContext } from 'react'
import { Message } from 'semantic-ui-react'
import { Redirect } from 'react-router-dom'

import API from 'Api'
import { User } from 'Shared/Context'
import { useRequest } from 'Shared/Hooks';
import SimplePage from 'Shared/SimplePage'

const Verify = ({match}) => {
  const { code } = match.params;
  const [loading, error, run, result] = useRequest({})
  const [redirect, setRedirect] = useState(false)
  const { setUser } = useContext(User)

  useEffect(() => {
    run(API.verify({code}), user => {
      setUser(user);
      setTimeout(() => {
        setRedirect(true);
      }, 2500);
    })
  }, [run, setUser, setRedirect, code])

  if (redirect) {
    return <Redirect to="/posts"/>
  }

  return (
    <SimplePage title='Account Verification' centered loading={loading} error={error}>
      {result && result.id ? <Message positive>Success! You have verified your email!</Message> : false}
    </SimplePage>
  );
}

export default Verify;
