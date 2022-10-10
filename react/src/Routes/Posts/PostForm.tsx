import React, { useState, useEffect, useCallback } from 'react'
import { Form, Message, Button } from 'semantic-ui-react'
import { useParams } from 'react-router'
import { Redirect } from 'react-router'

import API from 'Api'
import SimplePage from 'Shared/SimplePage';
import { useRequest, useFields } from 'Shared/Hooks';
import { InputChangeHandler, Post, TextAreaChangeHandler } from 'Shared/Types'

const PostForm = () => {
  const params = useParams<{ id: string }>();
  const post = {
    id: Number(params.id),
    title: '',
    body: ''
  };
  const [loading, error, run] = useRequest<Post>(post)
  const [fields, handleChange, setFields] = useFields<Post>(post)
  const [redirectTo, setRedirectTo] = useState('');

  const handleSubmit = useCallback(() => {
    const action = post.id ? API.updatePost(fields) : API.createPost(fields);
    run(action, () => {
      setRedirectTo('/posts')
    })
  }, [post.id, fields, run])

  const handleDelete = useCallback(() => {
    run(API.deletePost(post.id), () => {
      setRedirectTo('/posts')
    })
  }, [run, post.id])

  // if we have a post ID, fetch it
  useEffect(() => {
    if (post.id) {
      run(API.getPost(post.id), (post) => {
        setFields(post);
      });
    }
  }, [post.id, run, setFields])

  if (redirectTo) {
    return <Redirect to={redirectTo} />
  }

  const { title, body } = fields;

  return (
    <SimplePage icon='file alternate outline' title={post.id ? `Edit Post #${post.id}` : 'Create a Post'}>
      <Form error name="createPost" loading={loading} onSubmit={handleSubmit}>
        <Message error>{error}</Message>
        <Form.Input
          autoFocus
          size="big"
          name="title"
          type="text"
          placeholder="Post Title"
          required
          value={title}
          onChange={handleChange as InputChangeHandler} />
        <Form.TextArea
          name="body"
          rows={4}
          placeholder="Post content"
          required
          value={body}
          onChange={handleChange as TextAreaChangeHandler} />
        <Button primary size="huge" type="submit">Save</Button>
        {post.id ? <Button negative size="huge" type="button" onClick={handleDelete}>Delete</Button> : false}
      </Form>
    </SimplePage>
  )
}

export default PostForm;
