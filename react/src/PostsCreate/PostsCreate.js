import React, { Component } from 'react'
import { Container, Form, Message, Button } from 'semantic-ui-react'
import { Redirect } from 'react-router'
import API from '../api'

const defaultState = {
  loading: false,
  title: '',
  body: '',
  error: '',
  redirectTo: ''
};

export default class Posts extends Component {
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

    const {title, body} = this.state;

    API.createPost({
      title: title,
      body: body
    })
    .then((post) => {
      this.setState({
        loading: false,
        redirectTo: `/dashboard`
        // redirectTo: `/posts/${post.id}`
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
    const { loading, error, title, body, redirectTo} = this.state;
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
