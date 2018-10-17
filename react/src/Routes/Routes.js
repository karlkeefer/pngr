import React from 'react'
import { Subscribe } from 'unstated'
import { Switch, Route } from 'react-router-dom'
import PrivateRoute from './Helpers'

import UserContainer from '../Containers/User'

import Home from './Home/Home'
import LogIn from './LogIn/LogIn'
import SignUp from './SignUp/SignUp'
import NoMatch from './NoMatch/NoMatch'
import Verify from './Verify/Verify'

import Dashboard from './Dashboard/Dashboard'
import PostsCreate from './PostsCreate/PostsCreate'

const Routes = () => (
  <Switch>
    <Route exact path="/" component={Home} />

    <Route path="/signup" component={SignUp} />
    <Route path="/login" render={(props) => 
      <Subscribe to={[UserContainer]}>
        {userContainer => (
          <LogIn {...props} userContainer={userContainer}/>
        )}
      </Subscribe>
    } />
    <Route path="/verify/:verification" render={(props) =>
      <Subscribe to={[UserContainer]}>
        {userContainer => (
          <Verify {...props} userContainer={userContainer}/>
        )}
      </Subscribe>
    } />

    <PrivateRoute path="/dashboard" component={Dashboard}/>
    <PrivateRoute path="/posts/create" component={PostsCreate}/>

    <Route component={NoMatch} />
  </Switch>
);

export default Routes;