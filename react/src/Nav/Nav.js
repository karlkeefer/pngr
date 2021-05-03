import React from 'react'
import { Menu, Container } from 'semantic-ui-react'
import { Subscribe } from 'unstated'
import { NavLink } from 'react-router-dom'

import UserContainer from 'Containers/User'

// helper for semanticUI + react-router
const Link = props => (
  <NavLink
    {...props}
    activeClassName="active"
  />
);

const Nav = () => (
  <Subscribe to={[UserContainer]}>
    {userContainer => (
      <Menu fixed="top" inverted>
        <Container>
          <Menu.Item as={Link} exact to="/" name="Home" />
          { userContainer.isLoggedIn() ? <Menu.Item as={Link} to="/posts" name="Posts" /> : false }
          <Menu.Menu position="right">
            { !userContainer.isLoggedIn() ? <Menu.Item as={Link} exact to="/login" name="Log In" /> : false }
            { !userContainer.isLoggedIn() ? <Menu.Item as={Link} exact to="/signup" name="Sign Up" /> : false }
            { userContainer.isLoggedIn() ? <Menu.Item link={true} onClick={userContainer.logout} content="Log Out"/> : false }
          </Menu.Menu> 
        </Container>
      </Menu>
    )}
  </Subscribe>
)

export default Nav;
