import React, { Component } from 'react'
import { Redirect } from 'react-router'
import { Route } from 'react-router-dom'
import { Subscribe } from 'unstated'
import { Loader, Container, Dimmer } from 'semantic-ui-react'

import UserContainer from '../Containers/User'

// if unauth'd, check the jwt and then redirect to login screen if still not auth'd
// otherwise, behave like a Route while attempting to respect both "render" and "component" properties
export class PrivateRoute extends Component {
  render = () => {
    const { component: C, render: R, ...rest } = this.props;
    return (
      <Route {...rest} render={(props) => (
        <Subscribe to={[UserContainer]}>
          {userContainer => {
            if (userContainer.state.user.id > 0) {
              if (R) {
                return R();
              }
              return <C {...props} />
            } else {
              return <CheckAndRedirect location={props.location} userContainer={userContainer}/>
            }
          }}
        </Subscribe>
      )} />
    );
  }
}

// check valid cookie, if invalid, redirect to login
class CheckAndRedirect extends Component {
  render = () => {
    if (this.props.userContainer.state.loading !== false) {
      return (
        <Container>
          <Dimmer active inverted>
            <Loader size="big">Loading</Loader>
          </Dimmer>
        </Container>
      );
    }

    return (
      <Redirect to={{
        pathname: '/login',
        state: { from: this.props.location }
      }} />
    );
  }
}

export class NoMatch extends Component {
  render() {
    return (
      <Container className="page">
        <h1>Page not found</h1>
        <p>The page you are trying to view does not exist!</p>
      </Container>
    );
  }
}