import React, { Component } from 'react'
import { Container } from 'semantic-ui-react'
import { Subscribe } from 'unstated'
import { Redirect } from 'react-router'

import UserContainer from '../../Containers/User'

import Posts from './Posts/Posts'

export default class Dashboard extends Component {
  render() {
    return (
      <Subscribe to={[UserContainer]}>
        {(userContainer) => {
          if (userContainer.state.user.id === 0) {
            return <Redirect to="login"/>;
          }

          return (
            <Container className="page">
              <h1>Dashboard</h1>
              <p>This page fetches some protected data that only the logged in user ({userContainer.state.user.email}) can see!</p>
              <Posts/>
            </Container>
          );
        }}
      </Subscribe>
    );
  }
}