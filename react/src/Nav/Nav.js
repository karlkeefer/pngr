import React, { useContext, useEffect, useState } from 'react'
import { Menu, Container, Button } from 'semantic-ui-react'
import { useLocation } from 'react-router'
import { NavLink } from 'react-router-dom'

import { User } from 'Shared/Context'
import { LoggedIn, Anon } from 'Shared/Roles';

import './responsive.css'

// helper for semanticUI + react-router
const MenuLink = props => (
  <NavLink
    {...props}
    activeClassName="active"
  />
);

const Nav = () => {
  const location = useLocation();
  const [open, setOpen] = useState(false)
  const {handleLogout} = useContext(User)

  useEffect(()=>{
    setOpen(false);
  }, [location])

  const menuClass = open ? '' : 'hidden';

  return <Menu stackable fixed="top" inverted>
    <Container>
      <Button id="toggler" fluid color='black' icon='sidebar' onClick={() => setOpen(!open)}/>
      <Menu.Menu className={menuClass} position="left" id="override">
        <Menu.Item as={MenuLink} exact to="/" name="Home" />
        <LoggedIn>
          <Menu.Item as={MenuLink} to="/posts" name="Posts" />
        </LoggedIn>
      </Menu.Menu>
      <Menu.Menu className={menuClass} position="right">
        <Anon>
          <Menu.Item as={MenuLink} exact to={{pathname:"/login", state:{from:location}}} name="Log In" />
          <Menu.Item as={MenuLink} exact to="/signup" name="Sign Up" />
        </Anon>
        <LoggedIn>
          <Menu.Item link={true} onClick={handleLogout} content="Log Out"/>
        </LoggedIn>
      </Menu.Menu> 
    </Container>
  </Menu>
}

export default Nav;
