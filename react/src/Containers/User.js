import { Container } from 'unstated'

function defaultState() {
  return {
    user: {
      id: 0
    }
  };
}

class UserContainer extends Container {
  state = defaultState()

  clearCurrentUser = () => {
    this.setState(defaultState());
  }

  setCurrentUser = (user) => {
    this.setState({
      user: user
    });
  }
}

// TODO: eww, exporting a singleton(!)
// this is setup so that we can modify state after API calls
// maybe there is a way to avoid that by making our API helper more closely tied to REACT
const u = new UserContainer();
export default u;