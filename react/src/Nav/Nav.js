import React, { useContext } from 'react'
import { Menu, Container } from 'semantic-ui-react'
import { NavLink } from 'react-router-dom'

import { User } from 'Shared/Context'
import { LoggedIn, Anon } from 'Shared/Roles';

// helper for semanticUI + react-router
const Link = props => (
  <NavLink
    {...props}
    activeClassName="active"
  />
);

const Nav = () => {
  const {handleLogout} = useContext(User)

  return <Menu fixed="top" inverted>
    <Container>
      <Menu.Item as={Link} exact to="/" name="Home" />
      <LoggedIn>
        <Menu.Item as={Link} to="/posts" name="Posts" />
      </LoggedIn>
      <Menu.Menu position="right">
        <Anon>
          <Menu.Item as={Link} exact to="/login" name="Log In" />
          <Menu.Item as={Link} exact to="/signup" name="Sign Up" />
        </Anon>
        <LoggedIn>
          <Menu.Item link={true} onClick={handleLogout} content="Log Out"/>
        </LoggedIn>
      </Menu.Menu> 
    </Container>
  </Menu>
}

export default Nav;
