import React, { useContext } from 'react'
import { Menu, Container } from 'semantic-ui-react'
import { NavLink } from 'react-router-dom'

import { User } from 'Shared/Context'

// helper for semanticUI + react-router
const Link = props => (
  <NavLink
    {...props}
    activeClassName="active"
  />
);

const Nav = () => {
  const {user, handleLogout} = useContext(User)

  return <Menu fixed="top" inverted>
    <Container>
      <Menu.Item as={Link} exact to="/" name="Home" />
      { user.status === 'active' ? <Menu.Item as={Link} to="/posts" name="Posts" /> : false }
      <Menu.Menu position="right">
        { user.status !== 'active' ? <Menu.Item as={Link} exact to="/login" name="Log In" /> : false }
        { user.status !== 'active' ? <Menu.Item as={Link} exact to="/signup" name="Sign Up" /> : false }
        { user.status === 'active' ? <Menu.Item link={true} onClick={handleLogout} content="Log Out"/> : false }
      </Menu.Menu> 
    </Container>
  </Menu>
}

export default Nav;
