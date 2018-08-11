import React, { Component } from 'react'
import { Menu, Container } from 'semantic-ui-react'
import { NavLink } from 'react-router-dom'

// helper for semanticUI + react-router
const Link = props => (
  <NavLink
    exact
    {...props}
    activeClassName="active"
  />
);

export default class Home extends Component {
  handleLogOut() {
    // TODO: move login/logout/sessionHandler stuff into a single spot

    // TODO: optimistically clear out local 
    // session info in case the logout API call fails

    fetch('/api/logout').then(resp => {
      if (resp.status !== 200) {
        // TODO: Warn about API failures in some generic way
      }
    });
  }

  render() {
    return (
      <Menu fixed="top" inverted>
        <Container>
          <Menu.Item as={Link} to="/" name="Home" />
          <Menu.Menu position="right">
            <Menu.Item as={Link} to="/login" name="Log In" />
            <Menu.Item link={true} onClick={this.handleLogOut} content="Log Out"/>
          </Menu.Menu>
        </Container>
      </Menu>
    );
  }
}