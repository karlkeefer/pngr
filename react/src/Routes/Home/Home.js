import React from 'react'
import { Link } from 'react-router-dom'

import SimplePage from 'Shared/SimplePage';

const Home = () => (
  <SimplePage icon='rocket' title='Welcome to PNGR!'>
    <p>This is a boilerplate app using React for the front-end, and Golang + Postgres for the backend.</p>
    <p>The only things implemented are...</p>
    <ul>
      <li>Account Creation</li>
      <li>Session Management</li>
      <li><b>CRUD</b> for simple "posts" content type.</li>
    </ul>
    <p><Link to="/signup">Sign Up</Link> to see how sessions work and create/view secured posts.</p>
  </SimplePage>
)

export default Home;
