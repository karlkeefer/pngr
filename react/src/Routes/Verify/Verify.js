import React, { useState, useEffect } from 'react'
import { Message } from 'semantic-ui-react'
import { Redirect } from 'react-router-dom'

import { useRequest } from 'Shared/Hooks';
import SimplePage from 'Shared/SimplePage'

const Verify = ({match, userContainer}) => {
  const { verification } = match.params;
  const [loading, error, run, result] = useRequest({})
  const [redirect, setRedirect] = useState(false)

  useEffect(() => {
    run(userContainer.verify({code: verification}))
      .then(() => {
        setTimeout(() => {
          setRedirect(true);
        }, 2500);
      })
  }, [run, userContainer, setRedirect, verification])

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
