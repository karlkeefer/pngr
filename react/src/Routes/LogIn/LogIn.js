import React, { Component } from 'react'
import { Form, Button, Message, Segment } from 'semantic-ui-react'
import { Redirect, Link } from 'react-router-dom'
import SimplePage from 'Shared/SimplePage';

function defaultState() {
  return {
    loading: false,
    error: '',
    fields: {
      email: '',
      pass: '',
    }
  };
};

export default class LogIn extends Component {
  state = defaultState()

  handleChange = (e, {name, value}) => {
    this.setState(state => {
      state.fields = Object.assign(state.fields, {[name]: value });
      return state;
    });
  }

  handleSubmit = (e, val) => {
    e.preventDefault();
    this.setState({
      loading: true,
      error: ''
    });

    this.props.userContainer.login(this.state.fields)
      .then(() => {
        this.setState(Object.assign({}, defaultState()));
      })
      .catch((error) => {
        this.setState({
          error,
          loading: false
        });
      });
  }

  render() {
    const { loading, error } = this.state;
    const { email, pass } = this.state.fields;
    const { from } = this.props.location.state || { from: { pathname: '/posts' } };

    if (this.props.userContainer.state.user.id > 0 && !loading) {
      return <Redirect to={from}/>;
    }

    return (
      <SimplePage title='Log In to your account' centered>
        <Segment.Group>
          <Segment>
            <Form error name="login" loading={loading} onSubmit={this.handleSubmit}>
              <Message error>{error}</Message>
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
          </Segment>
          <Segment>
            Don't have an account? <Link to="/signup">Sign Up</Link>.
          </Segment>
        </Segment.Group>
      </SimplePage>
    );
  }
}
