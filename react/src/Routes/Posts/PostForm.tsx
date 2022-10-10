import React, { useState, useEffect, useCallback } from 'react'
import { Form, Message, Button } from 'semantic-ui-react'
import { useParams } from 'react-router'
import { Redirect } from 'react-router'

import API from 'Api'
import { useRequest, useFields, TextAreaChangeHandler, InputChangeHandler } from 'Shared/Hooks';
import SimplePage from 'Shared/SimplePage';

type PostFields = {
  title: string,
  body: string
}

const PostForm = () => {
  const params = useParams<{ id: string }>();
  const postID = Number(params.id);
  const [loading, error, run] = useRequest({})
  const [fields, handleChange, setFields] = useFields({ title: '', body: '' })
  const [redirectTo, setRedirectTo] = useState('');

  const handleSubmit = useCallback(() => {
    const action = postID ? API.updatePost(fields) : API.createPost(fields);
    run(action, () => {
      setRedirectTo('/posts')
    })
  }, [postID, fields, run])

  const handleDelete = useCallback(() => {
    run(API.deletePost(postID), () => {
      setRedirectTo('/posts')
    })
  }, [run, postID])

  // if we have a post ID, fetch it
  useEffect(() => {
    if (postID) {
      run(API.getPost(postID), (post: Object) => {
        setFields(post);
      });
    }
  }, [postID, run, setFields])

  if (redirectTo) {
    return <Redirect to={redirectTo} />
  }

  const { title, body } = fields as PostFields;

  return (
    <SimplePage icon='file alternate outline' title={postID ? `Edit Post #${postID}` : 'Create a Post'}>
      <Form error name="createPost" loading={!!loading} onSubmit={handleSubmit}>
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
        {postID ? <Button negative size="huge" type="button" onClick={handleDelete}>Delete</Button> : false}
      </Form>
    </SimplePage>
  )
}

export default PostForm;
