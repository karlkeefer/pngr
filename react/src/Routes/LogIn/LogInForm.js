import React, { useCallback, useContext } from 'react'
import { Form, Button, Message } from 'semantic-ui-react'
import { Redirect, useLocation } from 'react-router-dom'

import API from 'Api'
import { User } from 'Shared/Context'
import { useRequest, useFields } from 'Shared/Hooks';

const empty = {email: '', pass: ''};

const LogInForm = () => {
  const location = useLocation()
  const {user, setUser} = useContext(User)
  const [loading, error, run] = useRequest({})
  const [fields, handleChange, setFields] = useFields(empty)

  const handleSubmit = useCallback(() => {
    run(API.login(fields), user=>{
      setUser(user);
      setFields(empty)
    });
  }, [run, fields, setFields, setUser])

  if (user.id > 0 && !loading) {
    const { from } = location.state || { from: { pathname: '/posts' } };
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

export default LogInForm;
