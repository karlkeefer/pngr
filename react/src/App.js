import React, { Component } from 'react'
import { Provider } from 'unstated'
import { BrowserRouter as Router } from 'react-router-dom'

import UserContainer from './Containers/User'

import Nav from './Nav/Nav'
import Routes from './Routes/Routes'

export default class App extends Component {
  render() {
    return (
      <Provider inject={[UserContainer]}>
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
