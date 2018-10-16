import React, { Component } from 'react'
import { Form, Button, Message, Loader, Container, Grid, Segment, Dimmer } from 'semantic-ui-react'
import { Redirect } from 'react-router'
import { Link } from 'react-router-dom'
import { Subscribe } from 'unstated'

import API from '../../Api'
import UserContainer from '../../Containers/User'

function defaultState() {
  return {
    loading: false,
    checkingCookie: false,
    error: '',
    fields: {
      email: '',
      pass: '',
    }
  };
};

export default class LogIn extends Component {
  state = defaultState()

  componentDidMount() {
    // check existing cookie first
    this.setState({checkingCookie: true});

    API.whoami()
      .then(this._handleSuccess)
      .catch(this._handleError);
  }

  _handleError = (error) => {
    this.setState({
      error,
      loading: false,
      checkingCookie: false
    });
  }

  _handleSuccess = () => {
    this.setState(Object.assign({}, defaultState()));
  }

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

    API.login(this.state.fields)
      .then(this._handleSuccess)
      .catch(this._handleError);
  }

  render() {
    const { loading, error, checkingCookie } = this.state;
    const { email, pass } = this.state.fields;
    const { from } = this.props.location.state || { from: { pathname: '/dashboard' } };

    if (checkingCookie) {
      return (
        <Container>
          <Dimmer active inverted>
            <Loader size="big">Loading</Loader>
          </Dimmer>
        </Container>
      );
    }

    return (
      <Subscribe to={[UserContainer]}>
        {userContainer => {
          if (userContainer.state.user.id > 0 && !checkingCookie && !loading) {
            return <Redirect to={from}/>;
          }
          return (
            <Container className="page">
              <Grid centered>
                <Grid.Column textAlign="center" mobile={16} tablet={8} computer={6}>
                  <h1>Log In to your account</h1>
                  <Segment.Group>
                    <Segment>
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
                    </Segment>
                    <Segment>
                      Don't have an account? <Link to="/signup">Sign Up</Link>.
                    </Segment>
                  </Segment.Group>
                </Grid.Column>
              </Grid>
            </Container>
          );
        }}
      </Subscribe>
    );
  }
}
