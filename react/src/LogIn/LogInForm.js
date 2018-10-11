import React, { Component } from 'react'
import { Form, Button, Message } from 'semantic-ui-react'
import { Redirect } from 'react-router'

import { Subscribe } from 'unstated'
import API from '../api'

const defaultState = {
  loading: false,
  email: '',
  pass: '',
  error: ''
};

export default class LogInForm extends Component {
  state = defaultState

  handleChange = (e, {name, value}) => {
    this.setState({[name]: value });
  }

  handleSubmit = (e, val) => {
    e.preventDefault();
    this.setState({
      loading: true,
      error: ''
    });

    const {email, pass} = this.state;

    API.login({
      email: email,
      pass: pass
    })
    .then((success) => {
      // redirects immediately from api.state.user.ID change
    })
    .catch((error) => {
      this.setState({
        loading: false,
        error: error
      });
    });
  }

  render() {
    const {loading, email, pass, error} = this.state;

    return (
      <Subscribe to={[API]}>
        {api => {
          if (api.state.user.ID > 0) {
            return <Redirect to='/dashboard' />;
          }
          return (
            <Form name="login" loading={loading} onSubmit={this.handleSubmit}>
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
              <Button primary fluid size="huge" type="submit">Log In</Button>
            </Form>
          );
        }}
      </Subscribe>
    );
  }
}