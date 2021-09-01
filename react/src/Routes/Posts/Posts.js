import React, { useEffect, useContext } from 'react'
import { Link } from 'react-router-dom'
import { Segment, Message, Header, Button } from 'semantic-ui-react'

import API from 'Api'
import { useRequest } from 'Shared/Hooks'
import { User } from 'Shared/Context'
import SimplePage from 'Shared/SimplePage'

const Posts = () => {
  const [loading, error, run, posts] = useRequest([])
  const { user } = useContext(User)

  useEffect(()=>{
    run(API.getPosts())
  }, [run])

  return (
    <SimplePage icon='file' title='My Posts' loading={loading} error={error}>
      <p>This page fetches some protected data that only the logged in user ({user.email}) can see!</p>
      {posts.length === 0 ? <Message warning>No data to show</Message> : false }
      {posts.map(Post)}
      <Button as={Link} to='/post/create' primary icon='plus' content='New post'/>
    </SimplePage>
  )
}

export default Posts;

const Post = ({id, title, body}, i) => (
  <Segment.Group key={id}>
    <Header attached='top' as='h3'>
      {title} <Button compact basic as={Link} to={`/post/${id}/edit`} content='Edit' style={{marginLeft:'1em'}}/>
    </Header>
    <Segment attached='bottom' content={body}/>
  </Segment.Group>
)
