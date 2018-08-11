import React, { Component } from 'react'
import { Form, Button } from 'semantic-ui-react'

const defaultState = {
  loading: false,
  email: '',
  pass: ''
};

export default class SignUpForm extends Component {
  state = defaultState

  handleChange = (e, {name, value}) => {
    this.setState({ [name]: value });
  }

  handleSubmit = (e, val) => {
    e.preventDefault();
    this.setState({loading: true});

    const {email, pass} = this.state;

    fetch('/api/signup', {
      method: 'POST', 
      body: JSON.stringify({
        email: email,
        pass: pass
      })
    }).then(resp => resp.json())
    .then((result)=> {
      console.log(result.URL);
      this.setState(defaultState);
    });
  }

  render() {
    const {loading, email, pass} = this.state;

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