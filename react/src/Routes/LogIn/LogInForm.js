import React, { useCallback } from 'react'
import { Form, Button, Message } from 'semantic-ui-react'
import { Redirect } from 'react-router-dom'

import { useRequest, useFields } from 'Shared/Hooks';

const empty = {email: '', pass: ''};

const LogIn = ({location, userContainer}) => {
  const { from } = location.state || { from: { pathname: '/posts' } };
  const [loading, error, run] = useRequest({})
  const [fields, handleChange, setFields] = useFields(empty)

  const handleSubmit = useCallback(() => {
    run(userContainer.login(fields))
      .then(()=>{
        setFields(empty)
      });
  }, [userContainer, fields, setFields, run])

  if (userContainer.state.user.id > 0 && !loading) {
    return <Redirect to={from}/>;
  }

  const { email, pass } = fields;

  return (
    <Form error name="login" loading={loading} onSubmit={handleSubmit}>
      <Message error>{error}</Message>
      <Form.Input
        size="big"
        name="email"
        type="email"
        placeholder="Email"
        required
        value={email}
        onChange={handleChange} />
      <Form.Input
        size="big"
        name="pass"
        type="password"
        placeholder="Password"
        required
        value={pass}
        onChange={handleChange} />
      <Button primary fluid size="huge" type="submit">Log In</Button>
    </Form>
  )
}

export default LogIn;
