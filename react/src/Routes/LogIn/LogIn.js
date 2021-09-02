import React from 'react'
import { Segment } from 'semantic-ui-react'
import { Link } from 'react-router-dom'

import SimplePage from 'Shared/SimplePage';
import LogInForm from './LogInForm'

const LogIn = (props) => (
  <SimplePage title='Log In to your account' centered>
    <Segment.Group>
      <Segment>
        <LogInForm {...props}/>
      </Segment>
      <Segment>
        Don't have an account? <Link to="/signup">Sign Up</Link>.<br/>
        <Link to="/reset">I forgot my password</Link>.
      </Segment>
    </Segment.Group>
  </SimplePage>
);

export default LogIn;
