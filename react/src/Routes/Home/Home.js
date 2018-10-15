import React, { Component } from 'react'
import { Container } from 'semantic-ui-react'

export default class Home extends Component {
  render() {
    return (
      <Container className="page">
        <h1>Welcome!</h1>
        <p>This is a boilerplate app using React for the front-end, and Golang + Postgres for the backend.</p>
        <p>The only things implemented are basic account creation, session management, and super basic 'content' creation.</p>
      </Container>
    );
  }
}