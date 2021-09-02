// internal utils
const _get = (url, body) => {
  return _fetch('GET', url, body);
}

const _post = (url, body) => {
  return _fetch('POST', url, body);
}

const _delete = (url, body) => {
  return _fetch('DELETE', url, body);
}

const _put = (url, body) => {
  return _fetch('PUT', url, body);
}

const _fetch = (method, url, body) => {
  return fetch(`/api${url}`, {
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
    if (result.error) {
      return Promise.reject(result.error);
    }
    return Promise.resolve(result);
  })
  .catch(error => {
    return Promise.reject(error.toString());
  });
}

export default class API {
  // User stuff
  static signup = (body) => {
    return _post('/user', body);
  }
  static verify = (body) => {
    return _post('/user/verify', body);
  }
  static updatePassword = (body) => {
    return _put('/user/password', body)
  }

  // Session stuff
  static whoami = () => {
    // validates existing jwt from cookies
    // and sends back parsed user data from that token
    return _get('/user');
  }
  static logout = () => {
    return _delete('/session');
  }
  static login = (body) => {
    return _post('/session', body);
  }
  static reset = (body) => {
    return _post('/reset', body);
  }
  static checkReset = (code) => {
    return _get(`/reset/${code}`);
  }

  // Post stuff
  static getPosts = () => {
    return _get('/posts');
  }
  static createPost = (body) => {
    return _post('/posts', body);
  }
  static getPost = (id) => {
    return _get(`/posts/${id}`);
  }
  static updatePost = (body) => {
    return _put('/posts', body);
  }
  static deletePost = (id) => {
    return _delete(`/posts/${id}`);
  }
}
