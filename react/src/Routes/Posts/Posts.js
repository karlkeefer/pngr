import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import { Container, Segment, Message, Header, Button, Icon } from 'semantic-ui-react'

import { Subscribe } from 'unstated'
import UserContainer from 'Containers/User'
import API from 'Api'

export default class Posts extends Component {
  state = {
    loading: true,
    posts: [],
    error: ''
  }
  
  componentDidMount() {
    API.getPosts()
      .then(posts => {
        this.setState({posts: posts, loading: false});
      })
      .catch(error => {
        this.setState({error: error, loading: false});
      });
  }

  render() {
    const { loading, posts, error } = this.state;
    return (
      <Subscribe to={[UserContainer]}>
        {(userContainer) => (
          <Container className="page">
            <Header as='h1'>
              <Icon name='file' />
              My Posts
            </Header>
            <p>This page fetches some protected data that only the logged in user ({userContainer.state.user.email}) can see!</p>
            {error ? <Message negative>{error}</Message> : false }
            {posts.length === 0 && !loading ? <Message warning>No posts to show</Message> : false }
            {posts.map(({id, title, body}, i) => (
              <Segment.Group key={id}>
                <Header attached='top' as='h3'>
                  {title}
                  &nbsp;&nbsp;&nbsp;
                  <Button compact basic as={Link} to={`/post/${id}/edit`} content='Edit'/>
                </Header>
                <Segment attached='bottom'>
                  {body}
                </Segment>
              </Segment.Group>
            ))}
            
            <Button as={Link} to='/post/create' primary icon='plus' content='New post'/>
          </Container>
        )}
      </Subscribe>
    );
  }
}
