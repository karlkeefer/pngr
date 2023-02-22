import React, { useCallback } from 'react'

import { Form, Button, Message } from 'semantic-ui-react'

import API from 'Api'
import { useRequest, useFields } from 'Shared/Hooks';

const ChangePasswordForm = () => {
  const [loading, error, run, result] = useRequest({ success: false })
  const {fields, handleChange} = useFields({ pass: '' })

  const handleSubmit = useCallback(() => {
    run(API.updatePassword(fields))
  }, [run, fields])

  const { pass } = fields;

  if (result && result.success) {
    return <Message success>Password Updated!</Message>
  }

  return (
    <Form error name="login" loading={loading} onSubmit={handleSubmit}>
      <Message error>{error}</Message>
      <Form.Input
        size="big"
        name="pass"
        type="password"
        placeholder="Password"
        required
        value={pass}
        onChange={handleChange} />
      <Button primary fluid size="huge" type="submit">Update Password</Button>
    </Form>
  )
}

export default ChangePasswordForm;
