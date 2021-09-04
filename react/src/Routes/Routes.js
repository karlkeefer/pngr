import React from 'react'
import { Switch, Route } from 'react-router-dom'

import { PrivateRoute, NoMatch } from 'Routes/Helpers'

import Home from 'Routes/Home/Home'
import SignUp from 'Routes/SignUp/SignUp'
import LogIn from 'Routes/LogIn/LogIn'
import Reset from 'Routes/Reset/Reset'
import CheckReset from 'Routes/Reset/CheckReset'
import ChangePassword from 'Routes/Account/ChangePassword'
import Verify from 'Routes/Verify/Verify'
import Posts from 'Routes/Posts/Posts'
import Post from 'Routes/Posts/Post'
import PostForm from 'Routes/Posts/PostForm'

const Routes = (props) => (
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
