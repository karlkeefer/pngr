import React, { Component } from 'react'
import { Form, Button, Message } from 'semantic-ui-react'
import API from '../api'

const defaultState = {
  loading: false,
  email: '',
  pass: '',
  error: '',
  success: false,
  verifyURL: ''
};

export default class SignUpForm extends Component {
  state = defaultState

  handleChange = (e, {name, value}) => {
    this.setState({ [name]: value });
  }

  handleSubmit = (e, val) => {
    e.preventDefault();
    this.setState({
      loading: true,
      error: ''
    });

    const {email, pass} = this.state;

    API.signup({
      email: email,
      pass: pass
    })
    .then((res) => {
      this.setState({
        loading: false,
        success: true,
        verifyURL: res.URL
      });
    })
    .catch((error) => {
      this.setState({
        loading: false,
        error: error
      });
    });
  }

  render() {
    const {loading, email, pass, error, success, verifyURL} = this.state;

    if (success) {
      // TODO: this should be a message saying to check your email
      // and the verificationURL should be sent via email instead of passed in response
      return <Message positive><a href={verifyURL}>Click to confirm your email</a></Message>
    }

    return (
      <Form name="signup" loading={loading} onSubmit={this.handleSubmit}>
        {error ? <Message negative>{error}</Message> : ''}
        <Form.Input
          size="big"
          name="email"
          type="email"
          placeholder="Email"
          required
          value={email}
          onChange={this.handleChange} />
        <Form.Input
          size="big"
          name="pass"
          type="password"
          placeholder="Password"
          required
          value={pass}
          onChange={this.handleChange} />
        <Button positive fluid size="huge" type="submit">Create Account</Button>
      </Form>
    );
  }
}