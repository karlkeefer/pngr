import React, { Component } from 'react'
import { Container, Grid, Segment } from 'semantic-ui-react'
import { Redirect } from 'react-router-dom'

export default class Verify extends Component {

  state = {
    success: false,
    error: null
  }

  componentDidMount = () => {
    const { verification } = this.props.match.params

    fetch(`/api/verify/${verification}`)
      .then(resp => resp.json())
      .then(result => {
        if (result.Error) {
          this.setState({
            error: result.Error
          });
        } else {
          this.setState({
            success: true
          });
        }
      });
  }

  render() {
    const { success, error } = this.state;
     if (success) {
      return (
        <Redirect to="/account"/>
      );
    }

    return (
      <Container className="page">
        <Grid centered>
          <Grid.Column textAlign="center" mobile={16} tablet={8} computer={6}>
            <h1>Verifying your account...</h1>
            <Segment.Group>
              <Segment>
                {error ? 'Error: ' + error : 'Loading...'}
              </Segment>
            </Segment.Group>
          </Grid.Column>
        </Grid>
      </Container>
    );
  }
}