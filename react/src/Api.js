import UserContainer from './Containers/User'

// internal utils
function _get(url, body) {
  return _fetch('GET', url, body);
}

function _post(url, body) {
  return _fetch('POST', url, body);
}

function _delete(url, body) {
  return _fetch('DELETE', url, body);
}

function _fetch(method, url, body) {
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

function _setUser(user) {
  UserContainer.setCurrentUser(user);
  return Promise.resolve(user);
}

export default class API {
  // User stuff
  static signup = (body) => {
    return _post('/api/user', body);
  }

  static verify = (body) => {
    return _post('/api/user/verify', body)
      .then(_setUser);
  }

  static whoami = () => {
    // validates existing jwt from cookies
    // and sends back parsed user data from that token
    return _get('/api/user')
      .then(_setUser);
  }

  static logout = () => {
    UserContainer.clearCurrentUser();
    return _delete('/api/session');
  }

  static login = (body) => {
    return _post('/api/session', body)
      .then(_setUser);
  }

  // Post stuff
  static getPosts = () => {
    return _get('/api/posts');
  }
  static createPost = (body) => {
    return _post('/api/posts', body); 
  }
}
