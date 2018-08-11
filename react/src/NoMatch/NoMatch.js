import React, { Component } from 'react'
import { Container } from 'semantic-ui-react'

export default class NoMatch extends Component {
  render() {
    return (
      <Container className="page">
        <h1>Page not found</h1>
        <p>The page you are trying to view does not exist!</p>
      </Container>
    );
  }
}