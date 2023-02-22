import React, { useCallback } from 'react'

import { Form, Button, Message } from 'semantic-ui-react'

import API from 'Api'
import { useRequest, useFields } from 'Shared/Hooks';


const SignUpForm = () => {
  const [loading, error, run, result] = useRequest({ success: false })
  const {fields, handleChange} = useFields({ email: '', pass: '' })

  const handleSubmit = useCallback(() => {
    run(API.signup(fields))
  }, [run, fields])

  if (result && result.success) {
    return <Message positive>Check the developer console to see your verification link!</Message>
  }

  const { email, pass } = fields;

  return (
    <Form error name="signup" loading={loading} onSubmit={handleSubmit}>
      <Message error>{error}</Message>
      <Form.Input
        autoFocus
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
      <Button positive fluid size="huge" type="submit">Create Account</Button>
    </Form>
  )
}

export default SignUpForm;
