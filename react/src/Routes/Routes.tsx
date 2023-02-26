import { Routes as ReactRouterRoutes, Route } from "react-router-dom";

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
  <ReactRouterRoutes>
    <Route path="/" element={<Home />} />
    <Route path="/signup" element={<SignUp />}/>
    <Route path="/login" element={<LogIn />}/>
    <Route path="/reset" element={<Reset />}/>
    <Route path="/reset/:code" element={<CheckReset />}/>
    <Route path="/verify/:code" element={<Verify />}/>
    
    {/* crud post routes */}
    <Route
      path="/change-password"
      element={
        <RequireAuth>
          <ChangePassword />
        </RequireAuth>
      }
      />
    <Route
      path="/posts"
      element={
        <RequireAuth>
          <Posts />
        </RequireAuth>
      }
    />
    <Route
      path="/post/create"
      element={
        <RequireAuth>
          <PostForm />
        </RequireAuth>
      }
    />
    <Route
      path="/post/:id/edit"
      element={
        <RequireAuth>
          <PostForm />
        </RequireAuth>
      }
    />
    <Route
      path="/post/:id"
      element={
        <RequireAuth>
          <Post />
        </RequireAuth>
      }
    />
    <Route element={<NoMatch />} />
  </ReactRouterRoutes>
)

export default Routes;
