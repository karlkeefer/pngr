import React, { useState, useEffect, useContext } from 'react'
import { useParams } from 'react-router'
import { Redirect } from 'react-router-dom'

import API from 'Api'
import { User } from 'Shared/Context'
import { useRequest } from 'Shared/Hooks';
import SimplePage from 'Shared/SimplePage';

const CheckReset = () => {
  const { code } = useParams();
  const [loading, error, run] = useRequest({})
  const [redirect, setRedirect] = useState('/posts')
  const { user, userLoading, setUser } = useContext(User)

  useEffect(() => {
    if (!userLoading) {
      // wait until defailt whoami returns before attempting reset
      run(API.checkReset(code), user => {
        setRedirect('/account/password');
        setUser(user);
      });
    }
  }, [run, userLoading, setUser, setRedirect, code])

  if (user.id > 0 && redirect) {
    return <Redirect to={redirect}/>
  }

  return (
    <SimplePage title='Logging you in...' loading={userLoading || loading} error={error} centered/>
  )
}

export default CheckReset;
