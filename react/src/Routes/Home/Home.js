import React from 'react'
import { Segment } from 'semantic-ui-react'
import { Link } from 'react-router-dom'

import SimplePage from 'Shared/SimplePage';

const Home = () => (
  <SimplePage icon='rocket' title='Welcome to PNGR!'>
    <Segment>
      <p>This is a boilerplate app using React for the front-end, and Golang + Postgres for the backend.</p>
      <p>The only things implemented are...</p>
      <ul>
        <li>Account Creation</li>
        <li>Session Management</li>
        <li><b>CR</b>eate<b>U</b>pdate<b>D</b>elete for simple "posts"</li>
      </ul>
      <p><Link to="/signup">Sign Up</Link> to see how sessions work and create/view secured posts.</p>
    </Segment>
  </SimplePage>
)

export default Home;
