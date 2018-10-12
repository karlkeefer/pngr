import React, { Component } from 'react'
import { Provider, Subscribe } from 'unstated'

import API from './api'

import Nav from './Nav/Nav'

import Home from './Home/Home'
import LogIn from './LogIn/LogIn'
import SignUp from './SignUp/SignUp'
import NoMatch from './NoMatch/NoMatch'
import Verify from './Verify/Verify'
import PostsCreate from './PostsCreate/PostsCreate'

import Dashboard from './Dashboard/Dashboard'

import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'

export default class App extends Component {
  render() {
    return (
      <Provider inject={[API]}>
        <Router>
          <div className="wrapper">
            <Subscribe to={[API]}>
              {api => (
                <Nav api={api}/>
              )}
            </Subscribe>

            <section className="page">
              <Switch>
                <Route exact path="/" component={Home} />

                <Route path="/signup" component={SignUp} />
                <Route path="/login" component={LogIn} />
                <Route path="/verify/:verification" component={Verify}/>

                <Route path="/dashboard" component={Dashboard}/>
                <Route path="/posts/create" component={PostsCreate}/>

                <Route component={NoMatch} />
              </Switch>
            </section>
          </div>
        </Router>
      </Provider>
    );
  }
}
