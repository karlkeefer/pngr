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

function _put(endpoint, body) {
  return _fetch('PUT', endpoint, body);
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

export default class API {
  // User stuff
  static signup = (body) => {
    return _post('/api/user', body);
  }

  static verify = (body) => {
    return _post('/api/user/verify', body);
  }

  static whoami = () => {
    // validates existing jwt from cookies
    // and sends back parsed user data from that token
    return _get('/api/user');
  }

  static logout = () => {
    return _delete('/api/session');
  }

  static login = (body) => {
    return _post('/api/session', body);
  }

  // Post stuff
  static getPosts = () => {
    return _get('/api/posts');
  }
  static createPost = (body) => {
    return _post('/api/posts', body); 
  }
  static getPost = (id) => {
    return _get(`/api/posts/${id}`);
  }
  static updatePost = (body) => {
    return _put('/api/posts', body);
  }
  static deletePost = (id) => {
    return _delete(`/api/posts/${id}`);
  }
}
