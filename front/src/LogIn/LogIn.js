import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import { Container, Grid, Segment } from 'semantic-ui-react'

import LogInForm from './LogInForm'

export default class LogIn extends Component {
  render() {
    return (
      <Container className="page">
        <Grid centered>
          <Grid.Column textAlign="center" mobile={16} tablet={8} computer={6}>
            <h1>Log In to your account</h1>
            <Segment.Group>
              <Segment>
                <LogInForm/>
              </Segment>
              <Segment>
                Don't have an account? <Link to="/signup">Sign Up</Link>.
              </Segment>
            </Segment.Group>
          </Grid.Column>
        </Grid>
      </Container>
    );
  }
}