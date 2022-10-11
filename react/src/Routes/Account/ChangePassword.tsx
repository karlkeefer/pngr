import React from 'react'

import { Segment } from 'semantic-ui-react'

import SimplePage from 'Shared/SimplePage'

import ChangePasswordForm from './ChangePasswordForm'

const ChangePassword = () => (
  <SimplePage title='Update your password' centered>
    <Segment.Group>
      <Segment>
        <ChangePasswordForm />
      </Segment>
    </Segment.Group>
  </SimplePage>
);

export default ChangePassword;
