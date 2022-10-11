import React, { useState, useEffect, useContext } from 'react'

import { useParams } from 'react-router'
import { Redirect } from 'react-router-dom'
import { Message } from 'semantic-ui-react'

import API from 'Api'
import { useRequest } from 'Shared/Hooks';
import { User } from 'Shared/Models'
import SimplePage from 'Shared/SimplePage'
import { UserContainer } from 'Shared/UserContainer'

const Verify = () => {
  const { code } = useParams<{ code: string }>();
  const [loading, error, run, result] = useRequest<User>({} as User)
  const [redirect, setRedirect] = useState(false)
  const { userLoading, setUser } = useContext(UserContainer)

  useEffect(() => {
    if (!userLoading) {
      // wait until defailt whoami returns before attempting reset
      run(API.verify({ code }), user => {
        setUser(user);
        setTimeout(() => {
          setRedirect(true);
        }, 2500);
      })
    }
  }, [run, setUser, setRedirect, code, userLoading])

  if (redirect) {
    return <Redirect to="/posts" />
  }

  return (
    <SimplePage title='Account Verification' centered loading={loading} error={error}>
      {result && result.id ? <Message positive>Success! You have verified your email!</Message> : false}
    </SimplePage>
  );
}

export default Verify;
