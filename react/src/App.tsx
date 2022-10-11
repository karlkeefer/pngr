import React from 'react'

import { Helmet } from 'react-helmet'
import { BrowserRouter as Router } from 'react-router-dom'

import Nav from 'Nav/Nav'
import Routes from 'Routes/Routes'
import { WithUser } from 'Shared/UserContainer'

const App = () => (
  <WithUser>
    <Helmet
      defaultTitle="PNGR"
      titleTemplate="%s | PNGR"
    >
      {/* put meta tags here for opengraph and stuff */}
    </Helmet>
    <Router>
      <div id="wrapper">
        <Nav />
        <Routes />
      </div>
    </Router>
  </WithUser>
);

export default App;
