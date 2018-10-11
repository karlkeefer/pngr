import React, { Component } from 'react'
import { Menu, Container } from 'semantic-ui-react'
import { NavLink } from 'react-router-dom'
import { withRouter } from 'react-router'

// helper for semanticUI + react-router
const Link = props => (
  <NavLink
    exact
    {...props}
    activeClassName="active"
  />
);

class Nav extends Component {
  logout = () => {
    this.props.api.logout();
    this.props.history.push('/');
  }

  render() {
    let userMenu;

    if (this.props.api.state.user.ID > 0) {
      userMenu = (
      <Menu.Menu position="right">
        <Menu.Item as={Link} to="/dashboard" name="Dashboard" />
        <Menu.Item link={true} onClick={this.logout} content="Log Out"/>
      </Menu.Menu>
      );
    } else {
      userMenu = (
      <Menu.Menu position="right">
        <Menu.Item as={Link} to="/login" name="Log In" />
        <Menu.Item as={Link} to="/signup" name="Sign Up" />
      </Menu.Menu>
      );
    }

    return (
      <Menu fixed="top" inverted>
        <Container>
          <Menu.Item as={Link} to="/" name="Home" />
          {userMenu}
        </Container>
      </Menu>
    );
  }
}

export default withRouter(Nav);