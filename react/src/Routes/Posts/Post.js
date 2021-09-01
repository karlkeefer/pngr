import React, { useEffect } from 'react'
import { Link } from 'react-router-dom'
import { Button } from 'semantic-ui-react'

import API from 'Api'
import { useRequest } from 'Shared/Hooks'
import SimplePage from 'Shared/SimplePage'

const Post = ({match}) => {
  const postID = Number(match.params.id);
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
      <Button as={Link} to={`/post/${id}/edit`} content='Edit'/>
    </SimplePage>
  )
}

export default Post;
