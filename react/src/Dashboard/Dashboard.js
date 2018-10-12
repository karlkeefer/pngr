import React, { Component } from 'react'
import { Container } from 'semantic-ui-react'
import API from '../api'
import { Subscribe } from 'unstated'
import { Redirect } from 'react-router'

import Posts from './Posts/Posts'

export default class Dashboard extends Component {
  render() {
    return (
      <Subscribe to={[API]}>
        {(api) => {
          if (api.state.user.id === 0) {
            return <Redirect to="login"/>;
          }

          return (
            <Container className="page">
              <h1>Dashboard</h1>
              <p>This page fetches some protected data that only the logged in user ({api.state.user.email}) can see!</p>
              <Posts/>
            </Container>
          );
        }}
      </Subscribe>
    );
  }
}