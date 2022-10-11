import { useContext, useEffect, useState } from 'react'

import { useLocation } from 'react-router'
import { NavLink, NavLinkProps } from 'react-router-dom'
import { Button, Container, Menu } from 'semantic-ui-react'

import { Anon, LoggedIn } from 'Shared/Roles'
import { UserContainer } from 'Shared/UserContainer'

import './responsive.css'

// helper for semanticUI + react-router
const MenuLink = (props: NavLinkProps) => (
  <NavLink
    {...props}
    activeClassName="active"
  />
);

const Nav = () => {
  const location = useLocation();
  const [open, setOpen] = useState(false)
  const { handleLogout } = useContext(UserContainer)

  useEffect(() => {
    setOpen(false);
  }, [location])

  const menuClass = open ? '' : 'hidden';

  return <Menu stackable fixed="top" inverted>
    <Container>
      <Button id="toggler" fluid color='black' icon='sidebar' onClick={() => setOpen(!open)} />
      <Menu.Menu className={menuClass} position="left" id="override">
        <Menu.Item as={MenuLink} exact to="/" name="Home" />
        <LoggedIn>
          <Menu.Item as={MenuLink} to="/posts" name="Posts" />
        </LoggedIn>
      </Menu.Menu>
      <Menu.Menu className={menuClass} position="right">
        <Anon>
          <Menu.Item as={MenuLink} exact to={{ pathname: "/login", state: { from: location } }} name="Log In" />
          <Menu.Item as={MenuLink} exact to="/signup" name="Sign Up" />
        </Anon>
        <LoggedIn>
          <Menu.Item link={true} onClick={handleLogout} content="Log Out" />
        </LoggedIn>
      </Menu.Menu>
    </Container>
  </Menu>
}

export default Nav;
