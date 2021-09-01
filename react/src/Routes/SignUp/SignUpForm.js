import React, { useCallback } from 'react'
import { Form, Button, Message } from 'semantic-ui-react'

import { useRequest, useFields } from 'Shared/Hooks';
import API from 'Api'

const SignUpForm = () => {
  const [loading, error, run, result] = useRequest({})
  const [fields, handleChange] = useFields({email: '', pass: ''})

  const handleSubmit = useCallback(() => {
    run(API.signup(fields))
  }, [run, fields])

  if (result && result.URL) {
    // TODO: this should be a message saying to check your email
    // and the verificationURL should be sent via email instead of passed in response
    return (
      <Message positive>
        <a href={result.URL}>Click to confirm your email</a>
      </Message>
    )
  }

  const {email, pass} = fields;

  return (
    <Form error name="signup" loading={loading} onSubmit={handleSubmit}>
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
      <Button positive fluid size="huge" type="submit">Create Account</Button>
    </Form>
  )
}

export default SignUpForm;
