import React, { Component } from 'react'
import { Container, Grid, Segment } from 'semantic-ui-react'

export default class Verify extends Component {

  state = {
    loading: false
  }

  componentDidMount = () => {
    const { verification } = this.props.match.params
    
    this.setState({loading: true});

    fetch(`/api/verify/${verification}`)
      .then(resp => resp.json())
      .then((result)=> {
        console.log(result);
        // Redirect!
      });
  }

  render() {
    return (
      <Container className="page">
        <Grid centered>
          <Grid.Column textAlign="center" mobile={16} tablet={8} computer={6}>
            <h1>Verifying your account...</h1>
            <Segment.Group>
              <Segment>
                { /*progress bar here, replaced with success message and redirect*/ }
                Loading...
              </Segment>
            </Segment.Group>
          </Grid.Column>
        </Grid>
      </Container>
    );
  }
}