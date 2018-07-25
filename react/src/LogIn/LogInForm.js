import React, { Component } from 'react'
import { Form, Button } from 'semantic-ui-react'

const defaultState = {
  loading: false,
  email: '',
  password: ''
};

export default class LogInForm extends Component {
  state = defaultState

  handleChange = (e, {name, value}) => {
    this.setState({ [name]: value });
  }

  handleSubmit = (e, val) => {
    e.preventDefault();
    this.setState({loading: true});

    const {email, password} = this.state;

    fetch('/api/login', {
      method: 'POST', 
      body: JSON.stringify({
        id: email,
        password: password
      })
    }).then(resp => resp.json())
    .then((result)=> {
      console.log(result);
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
        <Button primary fluid size="huge" type="submit">Log In</Button>
      </Form>
    );
  }
}