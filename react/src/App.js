import React, { Component } from 'react'
import { Provider } from 'unstated'

import UserContainer from './Containers/User'
import API from './Api'

import Nav from './Nav/Nav'

import Home from './Routes/Home/Home'
import LogIn from './Routes/LogIn/LogIn'
import SignUp from './Routes/SignUp/SignUp'
import NoMatch from './Routes/NoMatch/NoMatch'
import Verify from './Routes/Verify/Verify'

import Dashboard from './Routes/Dashboard/Dashboard'
import PostsCreate from './Routes/PostsCreate/PostsCreate'

import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'

export default class App extends Component {
  componentDidMount = () => {
    API.whoami();
  }

  render() {
    return (
      <Provider inject={[UserContainer]}>
        <Router>
          <div className="wrapper">
            <Nav/>

            <section>
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
