import React from 'react'
import { Redirect } from 'react-router'
import { Route } from 'react-router-dom'
import { Subscribe } from 'unstated'
import { Loader, Container, Dimmer } from 'semantic-ui-react'

import UserContainer from 'Containers/User'
import SimplePage from 'Shared/SimplePage'

// if unauth'd, check the jwt and then redirect to login screen if still not auth'd
// otherwise, behave like a Route while attempting to respect both "render" and "component" properties
export const PrivateRoute = ({ component: C, render: R, ...rest }) => (
  <Route {...rest} render={(props) => (
    <Subscribe to={[UserContainer]}>
      {userContainer => {
        if (userContainer.state.user.id > 0) {
          if (R) {
            return R(props);
          }
          return <C {...props} userContainer={userContainer}/>
        } else {
          return <CheckAndRedirect location={props.location} userContainer={userContainer}/>
        }
      }}
    </Subscribe>
  )} />
);

// check valid cookie, if invalid, redirect to login
const CheckAndRedirect = ({userContainer, location}) => (
  userContainer.state.loading ?
    <Container>
      <Dimmer active inverted>
        <Loader size="big">Loading</Loader>
      </Dimmer>
    </Container>
    :
    <Redirect to={{
      pathname: '/login',
      state: { from: location }
    }} />
);

export const NoMatch = () => (
  <SimplePage title='Page not found'>
    <p>The page you are trying to view does not exist!</p>
  </SimplePage>
);
