import { useContext } from 'react'

import { Link } from 'react-router-dom'
import { Segment } from 'semantic-ui-react'

import SimplePage from 'Shared/SimplePage'
import { UserContainer } from 'Shared/UserContainer'

import LogInForm from './LogInForm'

const LogIn = () =>{ 
  const { userLoading } = useContext(UserContainer)
  return (
  <SimplePage title='Log In to your account' centered loading={userLoading}>
    <Segment.Group>
      <Segment>
        <LogInForm />
      </Segment>
      <Segment>
        Don't have an account? <Link to="/signup">Sign Up</Link>.<br />
        <Link to="/reset">I forgot my password</Link>.
      </Segment>
    </Segment.Group>
  </SimplePage>
);}

export default LogIn;
