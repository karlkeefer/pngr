import React from 'react'

import { Link } from 'react-router-dom'
import { Segment } from 'semantic-ui-react'

import SimplePage from 'Shared/SimplePage';

import SignUpForm from './SignUpForm'

const SignUp = () => (
  <SimplePage centered title='Create a new account'>
    <Segment.Group>
      <Segment>
        <SignUpForm/>
      </Segment>
      <Segment>
        Already have an account? <Link to="/login">Log In</Link>.
      </Segment>
    </Segment.Group>
  </SimplePage>
);

export default SignUp;
