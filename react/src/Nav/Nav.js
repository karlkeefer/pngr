import React, { Component } from 'react'
import { Menu, Container } from 'semantic-ui-react'
import { Subscribe } from 'unstated'
import { NavLink } from 'react-router-dom'

import UserContainer from '../Containers/User'

// helper for semanticUI + react-router
const Link = props => (
  <NavLink
    exact
    {...props}
    activeClassName="active"
  />
);

class Nav extends Component {
  render() {
    return (
      <Menu fixed="top" inverted>
        <Subscribe to={[UserContainer]}>
          {userContainer => (
            <NavInner userContainer={userContainer}/>
          )}
        </Subscribe>
      </Menu>
    );
  }
}

class NavInner extends Component {
  render() {
    const { status } = this.props.userContainer.state.user;
    if (status === 0) {
      return (
        <Container>
          <Menu.Menu position="right">
            <Menu.Item as={Link} to="/login" name="Log In" />
            <Menu.Item as={Link} to="/signup" name="Sign Up" />
          </Menu.Menu> 
        </Container>
      );
    }

    return (
      <Container>
        <StatusCheckMenuItem status={status} minStatus={1} as={Link} to="/posts" name="Posts" />
        <Menu.Menu position="right">
          <Menu.Item link={true} onClick={this.props.userContainer.logout} content="Log Out"/>
        </Menu.Menu>
      </Container>
    );
  }
}

// This shows/hides menu items based on the user status
// <StatusCheckMenuItem status={status} minStatus={1} as={Link} to="/posts" name="Posts" />
const StatusCheckMenuItem = (props) => {
  const {status, minStatus, ...rest} = props;
  return status >= minStatus ? <Menu.Item {...rest}/> : false;
}

export default Nav;