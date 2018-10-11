import { Container } from 'unstated'

// TODO: store/retrive token and user in cookie
// TODO: can we read the user props from the token on the client-side?
const defaultState = {
  token: '',
  user: {
    ID: 0
  }
};

export class APIContainer extends Container {
  state = defaultState

  signup = (body) => {
    return this._post('/api/signup', body);
  }

  logout = () => {
    // TODO: delete cookie
    this.setState(defaultState);
  }

  login = (body) => {
    return this._post('/api/login', body)
    .then(this._handleAuth);
  }

  verify = (code) => {
    return this._get(`/api/verify/${code}`)
    .then(this._handleAuth);
  }

  _handleAuth = (res) => {
    // TODO: set a cookie and read from the cookie on init
    // TODO: setup CSRF protection for said cookie
    this.setState({
      token: res.JWT, 
      user: res.User
    });
    return Promise.resolve(res);
  }

  _get = (url, body) => {
    return this._fetch('GET', url, body);
  }

  _post = (url, body) => {
    return this._fetch('POST', url, body);
  }

  _fetch = (method, url, body) => {
    return fetch(url, {
      method: method, 
      body: JSON.stringify(body)
    })
    .then(resp => resp.json())
    .then(result => {
      if (result.Error) {
        return Promise.reject(result.Error);
      }
      return Promise.resolve(result);
    });
  }
}

const API = new APIContainer();

export default API;