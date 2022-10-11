import React from 'react'

import { Switch, Route } from 'react-router-dom'

import ChangePassword from 'Routes/Account/ChangePassword'
import { PrivateRoute, NoMatch } from 'Routes/Helpers'
import Home from 'Routes/Home/Home'
import LogIn from 'Routes/LogIn/LogIn'
import Post from 'Routes/Posts/Post'
import PostForm from 'Routes/Posts/PostForm'
import Posts from 'Routes/Posts/Posts'
import CheckReset from 'Routes/Reset/CheckReset'
import Reset from 'Routes/Reset/Reset'
import SignUp from 'Routes/SignUp/SignUp'
import Verify from 'Routes/Verify/Verify'

const Routes = () => (
  <Switch>
    <Route exact path="/" component={Home} />
    <Route exact path="/signup" component={SignUp}/>
    <Route exact path="/login" component={LogIn}/>
    <Route exact path="/reset" component={Reset}/>
    <Route exact path="/reset/:code" component={CheckReset}/>
    <Route exact path="/verify/:code" component={Verify}/>
    <PrivateRoute exact path="/account/password" component={ChangePassword}/>
    <PrivateRoute exact path="/posts" component={Posts}/>
    <PrivateRoute exact path="/post/create" component={PostForm}/>
    <PrivateRoute exact path="/post/:id/edit" component={PostForm}/>
    <PrivateRoute exact path="/post/:id" component={Post}/>
    <Route component={NoMatch} />
  </Switch>
)

export default Routes;
