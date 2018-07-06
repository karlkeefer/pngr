import React, { Component } from 'react'

import Nav from './Nav/Nav'

import Home from './Home/Home'
import LogIn from './LogIn/LogIn'
import SignUp from './SignUp/SignUp'

import { BrowserRouter as Router, Route } from 'react-router-dom'

export default class App extends Component {
  render() {
    return (
      <Router>
        <div className="wrapper">
          <Nav/>

          <section className="page">
            <Route exact path="/" component={Home} />
            <Route path="/signup" component={SignUp} />
            <Route path="/login" component={LogIn} />
          </section>
        </div>
      </Router>
    );
  }
}
