import React from 'react'
import { Subscribe } from 'unstated'
import { Switch, Route } from 'react-router-dom'
import { PrivateRoute, NoMatch } from 'Routes/Helpers'

import UserContainer from 'Containers/User'

import Home from 'Routes/Home/Home'

import SignUp from 'Routes/SignUp/SignUp'
import LogIn from 'Routes/LogIn/LogIn'
import Verify from 'Routes/Verify/Verify'

import Posts from 'Routes/Posts/Posts'
import PostForm from 'Routes/Posts/PostForm'

const Routes = () => (
  <Subscribe to={[UserContainer]}>
    {userContainer => (
      <Switch>
        <Route exact path="/" component={Home} />

        <Route exact path="/signup" component={SignUp} />
        <Route exact path="/login" render={(props) => 
          <LogIn {...props} userContainer={userContainer}/>  
        } />
        <Route exact path="/verify/:verification" render={(props) =>
          <Verify {...props} userContainer={userContainer}/>
        } />

        <PrivateRoute exact path="/posts" component={Posts}/>
        <PrivateRoute exact path="/post/create" component={PostForm}/>
        <PrivateRoute exact path="/post/:id/edit" component={PostForm}/>

        <Route component={NoMatch} />
      </Switch>
    )}
  </Subscribe>
);

export default Routes;
