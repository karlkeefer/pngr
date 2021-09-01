import React from 'react'
import { BrowserRouter as Router } from 'react-router-dom'
import { Helmet } from 'react-helmet'

import { WithUser } from 'Shared/Context'

import Nav from 'Nav/Nav'
import Routes from 'Routes/Routes'

const App = () => (
  <WithUser>
    <Helmet
      defaultTitle="PNGR"
      titleTemplate="%s | PNGR"
    >
      {/* put meta tags here for opengraph and stuff */}
    </Helmet>
    <Router>
      <div className="wrapper">
        <Nav/>
        <Routes/>
      </div>
    </Router>
  </WithUser>
);

export default App;
