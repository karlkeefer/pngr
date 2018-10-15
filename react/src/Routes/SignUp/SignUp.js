import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import { Container, Grid, Segment } from 'semantic-ui-react'

import SignUpForm from './SignUpForm'

export default class LogIn extends Component {
  render() {
    return (
      <Container className="page">
        <Grid centered>
          <Grid.Column textAlign="center" mobile={16} tablet={8} computer={6}>
            <h1>Create a new account</h1>
            <Segment.Group>
              <Segment>
                <SignUpForm/>
              </Segment>
              <Segment>
                Already have an account? <Link to="/login">Log In</Link>.
              </Segment>
            </Segment.Group>
          </Grid.Column>
        </Grid>
      </Container>
    );
  }
}