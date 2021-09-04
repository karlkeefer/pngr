import React, { useContext } from 'react'
import { Redirect } from 'react-router'
import { Route } from 'react-router-dom'
import { Loader, Container, Dimmer } from 'semantic-ui-react'

import SimplePage from 'Shared/SimplePage'
import { User } from 'Shared/Context'

// check the user is logged in, and redirect to login screen if still not auth'd
export const PrivateRoute = ({ component: C, ...rest }) => {
  const {user, userLoading} = useContext(User)

  return (
    <Route {...rest} render={(props) => {
      if (userLoading) {
        return <BigLoader/>
      }

      if (!user.id) {
        return <Redirect to={{pathname: '/login', state: { from: rest.location }}} />
      }

      return <C/>
    }} />
  );
}

export const NoMatch = () => (
  <SimplePage icon='cancel' title='Not Found'>
    <p>The page you are trying to view does not exist!</p>
  </SimplePage>
);

const BigLoader = () => (
  <Container>
    <Dimmer active inverted>
      <Loader size="big"/>
    </Dimmer>
  </Container>
)
