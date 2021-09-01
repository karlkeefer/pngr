import React from 'react'

import { Switch, Route } from 'react-router-dom'
import { PrivateRoute, NoMatch } from 'Routes/Helpers'

import Home from 'Routes/Home/Home'
import SignUp from 'Routes/SignUp/SignUp'
import LogIn from 'Routes/LogIn/LogIn'
import Verify from 'Routes/Verify/Verify'
import Posts from 'Routes/Posts/Posts'
import PostForm from 'Routes/Posts/PostForm'

const Routes = () => (
  <Switch>
    <Route exact path="/" component={Home} />
    <Route exact path="/signup" component={SignUp}/>
    <Route exact path="/login" component={LogIn}/>
    <Route exact path="/verify/:verification" component={Verify}/>
    <PrivateRoute exact path="/posts" component={Posts}/>
    <PrivateRoute exact path="/post/create" component={PostForm}/>
    <PrivateRoute exact path="/post/:id/edit" component={PostForm}/>
    <Route component={NoMatch} />
  </Switch>
)

export default Routes;
