import React from 'react'
import { BrowserRouter as Router } from 'react-router-dom'
import { Helmet } from 'react-helmet'

import Nav from 'Nav/Nav'
import Routes from 'Routes/Routes'

import { WithUser } from 'Shared/Context'

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
        <Nav/>
        <Routes/>
      </div>
    </Router>
  </WithUser>
);

export default App;
