import React, { Component } from 'react'
import { Provider } from 'unstated'
import { BrowserRouter as Router } from 'react-router-dom'
import { Helmet } from 'react-helmet'

import UserContainer from 'Containers/User'

import Nav from 'Nav/Nav'
import Routes from 'Routes/Routes'

export default class App extends Component {
  render() {
    return (
      <Provider inject={[UserContainer]}>
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
      </Provider>
    );
  }
}
