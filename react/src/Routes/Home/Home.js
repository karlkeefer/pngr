import React, { Component } from 'react'
import { Container, Icon, Segment, Header } from 'semantic-ui-react'
import { Link } from 'react-router-dom'

export default class Home extends Component {
  render() {
    return (
      <Container className="page">
        <Header as='h1'>
          <Icon name='rocket'/>
          Welcome to PNGR!
        </Header>
        <Segment>
          <p>This is a boilerplate app using React for the front-end, and Golang + Postgres for the backend.</p>
          <p>The only things implemented are...</p>
          <ul>
            <li>Account Creation</li>
            <li>Session Management</li>
            <li><b>CR</b>eate<b>U</b>pdate<b>D</b>elete for simple "posts"</li>
          </ul>
          <p><Link to="/signup">Sign Up</Link> to see how sessions work and create/view secured posts.</p>
        </Segment>
      </Container>
    );
  }
}