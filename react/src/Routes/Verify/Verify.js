import React, { Component } from 'react'
import { Container, Grid, Message } from 'semantic-ui-react'
import { Redirect } from 'react-router-dom'

export default class Verify extends Component {

  state = {
    success: false,
    redirect: false,
    error: ''
  }

  componentDidMount = () => {
    const { verification } = this.props.match.params;
    
    this.props.userContainer.verify({code: verification})
      .then((res) => {
        this.setState({
          success: true
        });
        setTimeout(() => {
          this.setState({redirect: true});
        }, 2500);
      })
      .catch((e) => {
        this.setState({error: e});
      });
  }

  render() {
    const { success, error, redirect } = this.state;
    if (redirect) {
      return (
        <Redirect to="/posts"/>
      );
    }

    return (
      <Container className="page">
        <Grid centered>
          <Grid.Column textAlign="center" mobile={16} tablet={8} computer={6}>
            <h1>Verifying your account...</h1>
            <div>
              {!success && !error ? <Message>Loading...</Message> : ''}
              {error ? <Message negative>Error: {error}</Message> : ''}
              {success ? <Message positive>Success! You have verified your email!</Message>: ''}
            </div>
          </Grid.Column>
        </Grid>
      </Container>
    );
  }
}