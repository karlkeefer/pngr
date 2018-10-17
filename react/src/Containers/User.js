import { Container } from 'unstated'
import API from '../Api'

function defaultState() {
  return {
    loading: false,
    user: {
      id: 0
    }
  };
}

export default class UserContainer extends Container {
  constructor(props) {
    super(props);
    this.whoami()
      .finally(() => {
        this.setState({loading: false});
      })
  }

  state = Object.assign(defaultState(), {loading: true})

  verify = (body) => {
    return API.verify(body)
      .then(this._setCurrentUser);
  }

  whoami = () => {
    return API.whoami()
      .then(this._setCurrentUser);
  }

  login = (body) => {
    return API.login(body)
      .then(this._setCurrentUser);
  }

  logout = () => {
    this.setState(defaultState());
    return API.logout();
  }

  _setCurrentUser = (user) => {
    this.setState({user});
    return Promise.resolve(user);
  }
}