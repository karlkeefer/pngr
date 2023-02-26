import React, { useEffect } from 'react'

import { useParams } from 'react-router'
import { Link } from 'react-router-dom'
import { Button } from 'semantic-ui-react'

import API from 'Api'
import { useRequest } from 'Shared/Hooks'
import { Post } from 'Shared/Models'
import SimplePage from 'Shared/SimplePage'

const ViewPost = () => {
  const params = useParams<{ id: string }>();
  const postID = Number(params.id);
  const [loading, error, run, post] = useRequest({} as Post);

  // if we have a post ID, fetch it
  useEffect(() => {
    if (postID) {
      run(API.getPost(postID))
    }
  }, [run, postID])

  const { id, title, body } = post;

  return (
    <SimplePage icon='file alternate outline' title={title} loading={loading} error={error}>
      <p style={{whiteSpace: 'pre'}}>{body}</p>
      {id && id > 0 && 
        <Button as={Link} to={`/post/${id}/edit`} content='Edit' />}
    </SimplePage>
  )
}

export default ViewPost;
