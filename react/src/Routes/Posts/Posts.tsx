import React, { useContext, useEffect } from 'react'

import { Link } from 'react-router-dom'
import { Segment, Message, Header, Button, Placeholder, SegmentGroup, PlaceholderHeader } from 'semantic-ui-react'

import API from 'Api'
import { useRequest } from 'Shared/Hooks'
import { Post } from 'Shared/Models'
import SimplePage from 'Shared/SimplePage'
import { UserContainer } from 'Shared/UserContainer'

const Posts = () => {
  const [loading, error, run, posts] = useRequest([])
  const { user } = useContext(UserContainer)

  useEffect(() => {
    run(API.getPosts())
  }, [run])

  return (
    <SimplePage icon='copy outline' title='My Posts' error={error}>
      <p>This page fetches some protected data that only the logged in user ({user.email}) can see!</p>
      {loading && <PostsPlaceholder/>}
      {posts.length === 0 && !loading && 
        <Message warning>No posts found...</Message>}
      {posts.map(SinglePost)}
      <Button as={Link} to='/post/create' primary icon='plus' content='New post' />
    </SimplePage>
  )
}

export default Posts;

const PostsPlaceholder = () => (
  <SegmentGroup style={{ marginBottom: '1em' }}>
    <Segment attached='bottom'>
      <Placeholder>
        <Placeholder.Header>
          <Placeholder.Line/>
        </Placeholder.Header>
      </Placeholder>
    </Segment>
    <Segment attached='top'>
      <Placeholder>
        <Placeholder.Paragraph>
          <Placeholder.Line />
          <Placeholder.Line />
        </Placeholder.Paragraph>
      </Placeholder>
    </Segment>
  </SegmentGroup>
)

const SinglePost = ({ id, title, body }: Post) => (
  <Segment.Group key={id}>
    <Header attached='top' as='h3'>
      <Link to={`/post/${id}`}>{title}</Link>
    </Header>
    <Segment attached='bottom' content={body} />
  </Segment.Group>
)
