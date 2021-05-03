import React, { Component } from 'react'
import { Container, Form, Message, Button } from 'semantic-ui-react'
import { Redirect } from 'react-router'
import API from 'Api'

function defaultState(){
  return {
    loading: false,
    isUpdate: false,
    error: '',
    redirectTo: '',
    fields: {
      title: '',
      body: '',
    }
  };
}

export default class PostForm extends Component {
  state = defaultState()

  componentDidMount = () => {
    const post_id = Number(this.props.match.params.id);
    if (post_id) {
      this.setState({
        loading: true,
        isUpdate: true
      });

      API.getPost(post_id)
        .then(post => {
          this.setState(state => {
            state.fields = Object.assign({}, state.fields, post);
            state.loading = false;
            return state;
          });
        })
        .catch(error => {
          this.setState({
            error,
            isUpdate: false,
            loading: false
          });
        });
    }
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

    let action = this.state.isUpdate ? API.updatePost : API.createPost;

    action(this.state.fields)
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

  handleDelete = (e, val) => {
    e.preventDefault();
    this.setState({
      loading: true,
      error: ''
    });

    API.deletePost(this.state.fields.id)
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
    API.deletePost(this.state.fields.id);
  }

  render() {
    const { loading, isUpdate, error, redirectTo } = this.state;
    const { id, title, body } = this.state.fields;
    if (redirectTo) {
      return <Redirect to={redirectTo}/>
    }

    return (
      <Container className="page">
        <h1>{isUpdate ? `Edit Post #${id}` : 'Create a Post'}</h1>
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
          <Button primary size="huge" type="submit">Save</Button>
          {isUpdate ? <Button negative size="huge" type="button" onClick={this.handleDelete}>Delete</Button> : ''}
        </Form>
      </Container>
    );
  }
}
