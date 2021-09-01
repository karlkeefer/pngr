import React, { useContext } from 'react'
import { Redirect } from 'react-router'
import { Route } from 'react-router-dom'
import { Loader, Container, Dimmer } from 'semantic-ui-react'

import SimplePage from 'Shared/SimplePage'
import { User } from 'Shared/Context'

// if unauth'd, check the jwt and then redirect to login screen if still not auth'd
// otherwise, behave like a Route while attempting to respect both "render" and "component" properties
export const PrivateRoute = ({ component: C, render: R, ...rest }) => {
  const {user} = useContext(User)

  return (
    <Route {...rest} render={(props) => {
      if (user.id > 0) {
        if (R) {
          return R(props);
        }
        return <C {...props}/>
      } else {
        return <CheckAndRedirect location={props.location}/>
      }
    }} />
  );
}

// check valid cookie, if invalid, redirect to login
const CheckAndRedirect = ({location}) => {
  const {loading} = useContext(User)

  if (loading) {
    return (
      <Container>
        <Dimmer active inverted>
          <Loader size="big"/>
        </Dimmer>
      </Container>
    );
  }

  return (
    <Redirect to={{
      pathname: '/login',
      state: { from: location }
    }} />
  );
}

export const NoMatch = () => (
  <SimplePage title='Page not found'>
    <p>The page you are trying to view does not exist!</p>
  </SimplePage>
);
