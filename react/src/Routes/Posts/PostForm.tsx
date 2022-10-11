import React, { useState, useEffect, useCallback } from 'react'

import { useParams } from 'react-router'
import { Redirect } from 'react-router'
import { Form, Message, Button } from 'semantic-ui-react'

import API from 'Api'
import { useRequest, useFields } from 'Shared/Hooks';
import { Post } from 'Shared/Models'
import SimplePage from 'Shared/SimplePage';

const PostForm = () => {
  const params = useParams<{ id: string }>();
  const postID = Number(params.id);
  const [loading, error, run] = useRequest({} as Post)
  const {fields, handleInputChange, handleTextAreaChange, setFields} = useFields({} as Post)
  const [redirectTo, setRedirectTo] = useState('');

  // if we have a post ID, fetch it
  useEffect(() => {
    if (postID) {
      run(API.getPost(postID), (post) => {
        setFields(post);
      });
    }
  }, [postID, run, setFields])

  // handlers
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

  if (redirectTo) {
    return <Redirect to={redirectTo} />
  }

  const { id, title, body } = fields;

  return (
    <SimplePage icon='file alternate outline' title={id ? `Edit Post #${id}` : 'Create a Post'}>
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
          onChange={handleInputChange} />
        <Form.TextArea
          name="body"
          rows={4}
          placeholder="Post content"
          required
          value={body}
          onChange={handleTextAreaChange} />
        <Button primary size="huge" type="submit">Save</Button>
        {id && id > 0 &&
          <Button negative size="huge" type="button" onClick={handleDelete}>Delete</Button>}
      </Form>
    </SimplePage>
  )
}

export default PostForm;
