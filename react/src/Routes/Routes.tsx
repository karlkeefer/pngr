import React from 'react'

import { Switch, Route } from 'react-router-dom'

import ChangePassword from 'Routes/Account/ChangePassword'
import { NoMatch, RequireAuth } from "Routes/Helpers";
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
    
    {/* crud post routes */}
    <Route
        exact
        path="/account/password"
        render={() => (
          <RequireAuth redirectTo="/login">
            <ChangePassword />
          </RequireAuth>
        )}
      />
    <Route
      exact
      path="/posts"
      render={() => {
        return (
          <RequireAuth redirectTo="/login">
            <Posts />
          </RequireAuth>
        );
      }}
    />
    <Route
      exact
      path="/post/create"
      render={() => (
        <RequireAuth redirectTo="/login">
          <PostForm />
        </RequireAuth>
      )}
    />
    <Route
      exact
      path="/post/:id/edit"
      render={() => (
        <RequireAuth redirectTo="/login">
          <PostForm />
        </RequireAuth>
      )}
    />
    <Route
      exact
      path="/post/:id"
      render={() => (
        <RequireAuth redirectTo="/login">
          <Post />
        </RequireAuth>
      )}
    />
    <Route component={NoMatch} />
  </Switch>
)

export default Routes;
