import React, { useState, useEffect, useCallback } from 'react'
import { Form, Message, Button } from 'semantic-ui-react'
import { Redirect } from 'react-router'

import API from 'Api'
import { useRequest, useFields } from 'Shared/Hooks';
import SimplePage from 'Shared/SimplePage';

const PostForm = ({match}) => {
  const postID = Number(match.params.id);
  const [loading, error, run] = useRequest({})
  const [fields, handleChange, setFields] = useFields({title: '', body: ''})
  const [redirectTo, setRedirectTo] = useState('');

  const handleSubmit = useCallback(() => {
    run(postID ? API.updatePost(fields) : API.createPost(fields))
      .then(()=>{
        setRedirectTo('/posts')
      });
  }, [postID, fields, run])

  const handleDelete = useCallback(() => {
    run(API.deletePost(postID))
      .then(()=>{
        setRedirectTo('/posts')
      });
  }, [run, postID])

  // if we have a post ID, fetch it
  useEffect(()=>{
    if (postID) {
      run(API.getPost(postID))
        .then(post => {
          if (post) {
            setFields(post);
          }
        });
    }
  }, [postID, run, setFields])

  if (redirectTo) {
    return <Redirect to={redirectTo}/>
  }

  const {title, body} = fields;

  return (
    <SimplePage icon='edit outline' title={postID ? `Edit Post #${postID}` : 'Create a Post'}>
      <Form error name="createPost" loading={loading} onSubmit={handleSubmit}>
        <Message error>{error}</Message>
        <Form.Input
          size="big"
          name="title"
          type="text"
          placeholder="Post Title"
          required
          value={title}
          onChange={handleChange} />
        <Form.TextArea
          name="body"
          rows={4}
          placeholder="Post content"
          required
          value={body}
          onChange={handleChange} />
        <Button primary size="huge" type="submit">Save</Button>
        {postID ? <Button negative size="huge" type="button" onClick={handleDelete}>Delete</Button> : false }
      </Form>
    </SimplePage>
  )
}

export default PostForm;
