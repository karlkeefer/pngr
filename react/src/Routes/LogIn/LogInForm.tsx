import React, { useCallback, useContext } from 'react'

import { Redirect, useLocation } from 'react-router-dom'
import { Form, Button, Message } from 'semantic-ui-react'

import API from 'Api'
import { useRequest, useFields } from 'Shared/Hooks';
import { User } from 'Shared/Models'
import { UserContainer } from 'Shared/UserContainer'

const LogInForm = () => {
  const location = useLocation<{ from: string }>()
  const { user, setUser } = useContext(UserContainer)
  const [loading, error, run] = useRequest({} as User)
  const { fields, handleInputChange, setFields } = useFields({} as User)

  const handleSubmit = useCallback(() => {
    run(API.login(fields), user => {
      setUser(user);
      setFields({} as User)
    });
  }, [run, fields, setFields, setUser])

  if (user.id && user.id > 0 && !loading) {
    const { from } = location.state || { from: { pathname: '/posts' } };
    return <Redirect to={from} />;
  }

  const { email, pass } = fields;

  return (
    <Form error name="login" loading={loading} onSubmit={handleSubmit}>
      <Message error>{error}</Message>
      <Form.Input
        autoFocus
        size="big"
        name="email"
        type="email"
        placeholder="Email"
        required
        value={email}
        onChange={handleInputChange} />
      <Form.Input
        size="big"
        name="pass"
        type="password"
        placeholder="Password"
        required
        value={pass}
        onChange={handleInputChange} />
      <Button primary fluid size="huge" type="submit">Log In</Button>
    </Form>
  )
}

export default LogInForm;
