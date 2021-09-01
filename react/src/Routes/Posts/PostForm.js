import React, {useState, useEffect, useCallback} from 'react'
import { Form, Message, Button } from 'semantic-ui-react'
import { Redirect } from 'react-router'

import { useAPI, useFields } from 'Shared/Hooks';

import SimplePage from 'Shared/SimplePage';

const PostForm = ({match}) => {
  const [post, loading, error, run, API] = useAPI({})
  const [fields, setFields, handleChange] = useFields({title: '', body: ''})
  const [isUpdate, setIsUpdate] = useState(false)
  const [redirectTo, setRedirectTo] = useState('');

  const handleSubmit = useCallback(() => {
    run(isUpdate ? API.updatePost : API.createPost, fields)
      .then(()=>{
        setRedirectTo('/posts')
      });
  }, [API, isUpdate, fields, run])

  const handleDelete = useCallback(() => {
    run(API.deletePost, post.id)
      .then(()=>{
        setRedirectTo('/posts')
      });
  }, [API, run, post.id])

  // if we have a post ID, fetch it
  useEffect(()=>{
    const postID = Number(match.params.id);
    if (postID) {
      run(API.getPost, postID)
        .then(post => {
          if (post) {
            setFields(post);
            setIsUpdate(true);
          }
        });
    }
  }, [match, run, API, setFields])

  if (redirectTo) {
    return <Redirect to={redirectTo}/>
  }

  const {title, body} = fields;

  return (
    <SimplePage icon='edit outline' title={isUpdate ? `Edit Post #${post.id}` : 'Create a Post'}>
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
        {isUpdate ? <Button negative size="huge" type="button" onClick={handleDelete}>Delete</Button> : false }
      </Form>
    </SimplePage>
  )
}

export default PostForm;
