import React, { Component } from 'react'
import { Form, Button } from 'semantic-ui-react'

const defaultState = {
  loading: false,
  email: '',
  password: ''
};

export default class SignUpForm extends Component {
  state = defaultState

  handleChange = (e, {name, value}) => {
    this.setState({ [name]: value });
  }

  handleSubmit = (e, val) => {
    e.preventDefault();
    this.setState({loading: true});

    const {email, password} = this.state;

    fetch('/api/signup', {
      method: 'POST', 
      body: JSON.stringify({
        id: email,
        password: password
      })
    }).then(resp => resp.json())
    .then((result)=> {
      // TODO: redirect somewhere on successful login
      this.setState(defaultState);
    });
  }

  render() {
    const {loading, email, password} = this.state;

    return (
      <Form name="login" loading={loading} onSubmit={this.handleSubmit}>
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
          name="password"
          type="password"
          placeholder="Password"
          required
          value={password}
          onChange={this.handleChange} />
        <Button positive fluid size="huge" type="submit">Create Account</Button>
      </Form>
    );
  }
}