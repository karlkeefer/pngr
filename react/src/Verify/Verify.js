import React, { Component } from 'react'
import { Container, Grid, Message } from 'semantic-ui-react'
import { Redirect } from 'react-router-dom'

import API from '../api'

export default class Verify extends Component {

  state = {
    success: false,
    error: null
  }

  componentDidMount = () => {
    const { verification } = this.props.match.params

    API.verify(verification)
      .then((res) => {
        this.setState({success: true});
      })
      .catch((e) => {
        this.setState({error: e});
      });
  }

  render() {
    const { success, error } = this.state;
     if (success) {
      return (
        <Redirect to="/dashboard"/>
      );
    }

    return (
      <Container className="page">
        <Grid centered>
          <Grid.Column textAlign="center" mobile={16} tablet={8} computer={6}>
            <h1>Verifying your account...</h1>
            <div>
              {error ? <Message negative>Error: {error}</Message> : <Message>Loading...</Message>}
            </div>
          </Grid.Column>
        </Grid>
      </Container>
    );
  }
}