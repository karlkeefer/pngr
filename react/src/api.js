import { Container } from 'unstated'

const defaultState = {
  jwt: '',
  user: {
    id: 0
  }
};

export class APIContainer extends Container {
  constructor(props) {
    super(props)
    this.whoami();
  }

  state = defaultState

  // User stuff
  signup = (body) => {
    return this._post('/api/user', body);
  }

  verify = (body) => {
    return this._post('/api/user/verify', body)
      .then(this._setUser);
  }

  whoami = () => {
    // validates existing jwt from cookies
    // and sends back parsed user data from that token
    return this._get('/api/user')
      .then(this._setUser);
  }

  logout = () => {
    this.setState(defaultState);
    return this._delete('/api/session');
  }

  login = (body) => {
    return this._post('/api/session', body)
      .then(this._setUser);
  }

  _setUser = (user) => {
    this.setState({
      user: user
    });
    return Promise.resolve(user);
  }

  // Post stuff
  getPosts = () => {
    return this._get('/api/posts');
  }
  createPost = (body) => {
    return this._post('/api/posts', body); 
  }

  // internal utils
  _get = (url, body) => {
    return this._fetch('GET', url, body);
  }

  _post = (url, body) => {
    return this._fetch('POST', url, body);
  }

  _delete = (url, body) => {
    return this._fetch('DELETE', url, body);
  }


  _fetch = (method, url, body) => {
    return fetch(url, {
      method: method,
      body: JSON.stringify(body),
      headers: {
        // CSRF prevention
        // https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)_Prevention_Cheat_Sheet#Use_of_Custom_Request_Headers
        'X-Requested-With': 'XMLHttpRequest'
      }
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