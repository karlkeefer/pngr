import React, { Component } from 'react'
import { Container } from 'semantic-ui-react'
import API from '../api'
import { Subscribe } from 'unstated'
import { Redirect } from 'react-router'

export default class Dashboard extends Component {
  render() {
    if (API.state.user.ID === 0) {
        return <Redirect to="login"/>;
    }

    return (
      <Container className="page">
        <h1>Dashboard</h1>
        <Subscribe to={[API]}>
          {api => (
            <p>This page fetches some protected data that only the logged in user ({api.state.user.Email}) can see!</p>
          )}
        </Subscribe>
      </Container>
    );
  }
}