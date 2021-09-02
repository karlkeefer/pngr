import React, { useEffect, useContext } from 'react'
import { Link } from 'react-router-dom'
import { Segment, Message, Header, Button, Placeholder } from 'semantic-ui-react'

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
    <SimplePage icon='copy' title='My Posts' error={error}>
      <p>This page fetches some protected data that only the logged in user ({user.email}) can see!</p>
      {loading ? <Placeholder style={{marginBottom:'1em'}}>
        <Placeholder.Paragraph>
          <Placeholder.Line />
          <Placeholder.Line />
          <Placeholder.Line />
          <Placeholder.Line />
        </Placeholder.Paragraph>
      </Placeholder> : false }
      {posts.length === 0 && !loading ? <Message warning>No posts found...</Message> : false }
      {posts.map(Post)}
      <Button as={Link} to='/post/create' primary icon='plus' content='New post'/>
    </SimplePage>
  )
}

export default Posts;

const Post = ({id, title, body}, i) => (
  <Segment.Group key={id}>
    <Header attached='top' as='h3'>
      <Link to={`/post/${id}`}>{title}</Link>
    </Header>
    <Segment attached='bottom' content={body}/>
  </Segment.Group>
)
