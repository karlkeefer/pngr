export default class API {
  // SESSION
  static logout = () => (
    _delete('/session')
  )
  static login = body => (
    _post('/session', body)
  )

  // RESETS
  static reset = body => (
    _post('/reset', body)
  )
  static checkReset = (code) => (
    _get(`/reset/${code}`)
  )

  // USER
  static signup = body => (
    _post('/user', body)
  )
  static whoami = () => (
    _get('/user')
  )
  static verify = body => (
    _post('/user/verify', body)
  )
  static updatePassword = body => (
    _put('/user/password', body)
  )

  // POSTS
  static getPosts = () => (
    _get('/post')
  )
  static getPost = (id) => (
    _get(`/post/${id}`)
  )
  static createPost = body => (
    _post('/post', body)
  )
  static updatePost = body => (
    _put('/post', body)
  )
  static deletePost = id => (
    _delete(`/post/${id}`)
  )
}

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
