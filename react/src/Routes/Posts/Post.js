import React, { useEffect } from 'react'
import { useParams } from 'react-router'
import { Link } from 'react-router-dom'
import { Button } from 'semantic-ui-react'

import API from 'Api'
import { useRequest } from 'Shared/Hooks'
import SimplePage from 'Shared/SimplePage'

const Post = () => {
  const params = useParams();
  const postID = Number(params.id);
  const [loading, error, run, post] = useRequest({})

  // if we have a post ID, fetch it
  useEffect(()=>{
    if (postID) {
      run(API.getPost(postID))
    }
  }, [run, postID])

  const {id, title, body} = post;

  return (
    <SimplePage icon='file' title={title} loading={loading} error={error}>
      <p>{body}</p>
      {id ? <Button as={Link} to={`/post/${id}/edit`} content='Edit'/> : false}
    </SimplePage>
  )
}

export default Post;
