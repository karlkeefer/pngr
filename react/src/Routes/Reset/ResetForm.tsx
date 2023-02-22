import React, { useCallback } from 'react'

import { Form, Button, Message } from 'semantic-ui-react'

import API from 'Api'
import { useRequest, useFields } from 'Shared/Hooks';


const ResetForm = () => {
  const [loading, error, run, result] = useRequest({ success: false })
  const {fields, handleChange} = useFields({ email: '' })

  const handleSubmit = useCallback(() => {
    run(API.reset(fields))
  }, [run, fields])

  const { email } = fields;

  if (result && result.success) {
    return <Message success>Check the developer console to see your password reset link!</Message>
  }

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
      <Button primary fluid size="huge" type="submit">Reset Password</Button>
    </Form>
  )
}

export default ResetForm;
