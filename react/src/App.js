import React, { Component } from 'react'

import Nav from './Nav/Nav'

import Home from './Home/Home'
import LogIn from './LogIn/LogIn'
import SignUp from './SignUp/SignUp'
import NoMatch from './NoMatch/NoMatch'
import Verify from './Verify/Verify'

import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'

export default class App extends Component {
  render() {
    return (
      <Router>
        <div className="wrapper">
          <Nav/>

          <section className="page">
            <Switch>
              <Route exact path="/" component={Home} />
              <Route path="/signup" component={SignUp} />
              <Route path="/login" component={LogIn} />
              <Route path="/verify/:verification" component={Verify} />
              <Route component={NoMatch} />
            </Switch>
          </section>
        </div>
      </Router>
    );
  }
}
