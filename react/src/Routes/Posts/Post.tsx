import React, { useEffect } from 'react'
import { useParams } from 'react-router'
import { Link } from 'react-router-dom'
import { Button } from 'semantic-ui-react'

import API from 'Api'
import SimplePage from 'Shared/SimplePage'
import { useRequest } from 'Shared/Hooks'
import { Post } from 'Shared/Models'

const ViewPost = () => {
  const params = useParams<{ id: string }>();
  const [loading, error, run, post] = useRequest<Post>({
    id: Number(params.id),
    title: '',
    body: ''
  })

  // if we have a post ID, fetch it
  useEffect(() => {
    if (post.id) {
      run(API.getPost(post.id))
    }
  }, [run, post.id])

  const { id, title, body } = post;

  return (
    <SimplePage icon='file' title={title} loading={loading} error={error}>
      <p>{body}</p>
      {id ? <Button as={Link} to={`/post/${id}/edit`} content='Edit' /> : false}
    </SimplePage>
  )
}

export default ViewPost;
