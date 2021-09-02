import React from 'react'
import { Segment } from 'semantic-ui-react'

import SimplePage from 'Shared/SimplePage';
import ResetForm from './ResetForm'

const Reset = (props) => (
  <SimplePage title='Reset your password' centered>
    <Segment.Group>
      <Segment>
        <ResetForm {...props}/>
      </Segment>
    </Segment.Group>
  </SimplePage>
);

export default Reset;
