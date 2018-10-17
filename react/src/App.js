import React, { Component } from 'react'
import { Provider, Subscribe } from 'unstated'
import { BrowserRouter as Router, Switch, Route, Redirect } from 'react-router-dom'
import { Loader, Container, Dimmer } from 'semantic-ui-react'

import UserContainer from './Containers/User'

import Nav from './Nav/Nav'

import Home from './Routes/Home/Home'
import LogIn from './Routes/LogIn/LogIn'
import SignUp from './Routes/SignUp/SignUp'
import NoMatch from './Routes/NoMatch/NoMatch'
import Verify from './Routes/Verify/Verify'

import Dashboard from './Routes/Dashboard/Dashboard'
import PostsCreate from './Routes/PostsCreate/PostsCreate'

export default class App extends Component {
  render() {
    return (
      <Provider inject={[UserContainer]}>
        <Router>
          <div className="wrapper">
            <Nav/>

            <section>
              <Subscribe to={[UserContainer]}>
                {userContainer => (
                  <Switch>
                    <Route exact path="/" component={Home} />

                    <Route path="/signup" component={SignUp} />
                    <Route path="/login" render={(props) => <LogIn {...props} userContainer={userContainer}/>} />
                    <Route path="/verify/:verification" render={(props) => <Verify {...props} userContainer={userContainer}/>} />

                    <PrivateRoute path="/dashboard" component={Dashboard}/>
                    <PrivateRoute path="/posts/create" component={PostsCreate}/>

                    <Route component={NoMatch} />
                  </Switch>
                )}
              </Subscribe>
            </section>
          </div>
        </Router>
      </Provider>
    );
  }
}

class PrivateRoute extends Component {
  render = () => {
    const { component: C, render: R, ...rest } = this.props;
    return (
      <Route {...rest} render={(props) => (
        <Subscribe to={[UserContainer]}>
          {userContainer => {
            if (userContainer.state.user.id > 0) {
              if (R) {
                return R();
              }
              return <C {...props} />
            } else {
              return <CheckAndRedirect location={props.location} userContainer={userContainer}/>
            }
          }}
        </Subscribe>
      )} />
    );
  }
}

// check valid cookie, if invalid, redirect to login
class CheckAndRedirect extends Component {
  state = {
    checkingCookie: true
  }

  componentDidMount = () => {
    this.props.userContainer.whoami()
      .finally(() => {
        if (this.props.userContainer.state.user.id === 0) {
          this.setState({checkingCookie: false});
        }
      })
  }

  render = () => {
    if (this.state.checkingCookie) {
      return (
        <Container>
          <Dimmer active inverted>
            <Loader size="big">Loading</Loader>
          </Dimmer>
        </Container>
      );
    }

    return (
      <Redirect to={{
        pathname: '/login',
        state: { from: this.props.location }
      }} />
    );
  }
}