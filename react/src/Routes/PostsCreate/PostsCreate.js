import React, { Component } from 'react'
import { Container, Form, Message, Button } from 'semantic-ui-react'
import { Redirect } from 'react-router'
import API from '../../Api'

function defaultState(){
  return {
    loading: false,
    error: '',
    redirectTo: '',
    fields: {
      title: '',
      body: '',
    }
  };
}

export default class Posts extends Component {
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

    API.createPost(this.state.fields)
    .then((post) => {
      this.setState({
        loading: false,
        redirectTo: `/posts`
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
    const { loading, error, redirectTo } = this.state;
    const { title, body } = this.state.fields;
    if (redirectTo) {
      return <Redirect to={redirectTo}/>
    }

    return (
      <Container className="page">
        <h1>Create a Post</h1>
        <Form name="createPost" loading={loading} onSubmit={this.handleSubmit}>
          {error ? <Message negative>{error}</Message> : ''}
          <Form.Input
            size="big"
            name="title"
            type="text"
            placeholder="Post Title"
            required
            value={title}
            onChange={this.handleChange} />
          <Form.TextArea
            name="body"
            rows={4}
            placeholder="Post content"
            required
            value={body}
            onChange={this.handleChange} />
          <Button primary fluid size="huge" type="submit">Create Post</Button>
        </Form>
      </Container>
    );
  }
}
