import { Container } from 'unstated'
import { Cookies } from 'react-cookie'

const cookies = new Cookies();

const defaultState = {
  jwt: '',
  user: {
    ID: 0
  }
};

function parseToken (token) {
  const base64Url = token.split('.')[1];
  const base64 = base64Url.replace('-', '+').replace('_', '/');
  return JSON.parse(window.atob(base64));
};

export class APIContainer extends Container {
  constructor(props) {
    super(props);

    const jwt = cookies.get('jwt');

    this.state = Object.assign(defaultState, {jwt: jwt});

    if (jwt) {
      const claims = parseToken(jwt);
      if (claims && claims.user) {
        this.state.user = claims.user;
      }
    }
  }

  signup = (body) => {
    return this._post('/api/signup', body);
  }

  logout = () => {
    cookies.remove('jwt');
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
    // TODO: setup httpOnly, secure, and related cookie options
    // TODO: setup additional CSRF protection for this cookie
    
    cookies.set('jwt', res.JWT);
    const claims = parseToken(res.JWT);

    this.setState({
      jwt: res.JWT, 
      user: claims.user
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