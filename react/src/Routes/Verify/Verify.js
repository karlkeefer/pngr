import React, { Component } from 'react'
import { Message } from 'semantic-ui-react'
import { Redirect } from 'react-router-dom'
import SimplePage from 'Shared/SimplePage'

export default class Verify extends Component {

  state = {
    loading: false,
    error: false,
    redirect: false,
  }

  componentDidMount = () => {
    const { verification } = this.props.match.params;
    this.setState({loading: true});
    this.props.userContainer.verify({code: verification})
      .then((res) => {
        this.setState({loading: false});
        setTimeout(() => {
          this.setState({redirect: true});
        }, 2500);
      })
      .catch((error) => {
        this.setState({error, loading: false});
      });
  }

  render() {
    const { loading, error, redirect } = this.state;
    if (redirect) {
      return (
        <Redirect to="/posts"/>
      );
    }

    return (
      <SimplePage title='Account Verification' centered loading={loading} error={error}>
        <Message positive>Success! You have verified your email!</Message>
      </SimplePage>
    );
  }
}
