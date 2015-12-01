// TODO FIXME handle all those errors :)
import Storage from './storage.js'
import request from 'superagent'
const API = '/api'
const PrivateAPI = '/api/private'

export default {
  LoginUser (payload) {
    return Public.methods.loginUser(payload)
  },
  RegisterUser (payload) {
    return Public.methods.registerUser(payload)
  },
  UserInfo () {
    return Private.methods.userInfo()
  },
  UserLists () {
    return Private.methods.userLists()
  },
  GetList (title) {
    return Private.methods.getList(title)
  },
  AddList (payload) {
    return Private.methods.newList(payload)
  },
  DeleteList (id) {
    return Private.methods.deleteList(id)
  },
  CheckForBearer
}
let Public = {
  methods: {
    loginUser (data) {
      let params = Object.assign({}, Public.params.loginUser)
      params.send = data
      return api.reqAPI(params)
    },
    registerUser (data) {
      let params = Object.assign({}, Public.params.registerUser)
      params.send = data
      return api.reqAPI(params)
    }
  },
  params: {
    loginUser: {
      method: 'POST',
      path: '/login',
      send: {}
    },
    registerUser: {
      method: 'POST',
      path: '/register',
      send: {}
    }
  }
}

let Private = {
  methods: {
    userInfo () {
      return api.reqPrivateAPI(Private.params.userInfo)
    },
    userLists () {
      return api.reqPrivateAPI(Private.params.userLists)
    },
    getList (title) {
      let params = Object.assign({}, Private.params.List)
      params.path = params.path + title
      return api.reqPrivateAPI(params)
    },
    deleteList (id) {
      let params = Object.assign({}, Private.params.deleteList)
      params.path = params.path + id
      return api.reqPrivateAPI(params)
    },
    newList (payload) {
      let params = Object.assign({}, Private.params.newList)
      params.send = payload
      return api.reqPrivateAPI(params)
    }
  },
  params: {
    userInfo: {
      method: 'GET',
      path: '/user'
    },
    userLists: {
      method: 'GET',
      path: '/user/lists'
    },
    List: {
      method: 'GET',
      path: '/user/lists/'
    },
    newList: {
      method: 'POST',
      path: '/user/lists/new',
      send: {}
    },
    deleteList: {
      method: 'DELETE',
      path: '/user/lists/'
    }
  }
}
function CheckForBearer (resp) {
  return new Promise((resolve, reject) => {
    if (resp.body.bearer && resp.body.bearer.length > 0) {
      resolve({ resp: resp, bearer: resp.body.bearer })
    } else {
      // maybe a reason why a token would not been issued
      let error = resp.body.error
      let errors = resp.body.errors
      reject({ error: error, errors: errors})
    }
  })
}
const api = {
  authHeader: 'Authorization',
  reqAPI (params) {
    let req = this.req(params)
    .use(this.prefix(API))
    return new Promise((resolve, reject) => {
      req.end((err, resp) => {
        if (err) {
          reject(err)
        }
        resolve(resp)
      })
    })
  },
  reqPrivateAPI (params) {
    let req = this.req(params)
    .use(this.prefix(PrivateAPI))
    .use(this.authReq())
    return new Promise((resolve, reject) => {
      req.end((err, resp) => {
        if (err) {
          reject(err)
        }
        resolve(resp)
      })
    })
  },
  req (params) {
    let req = request(params.method, params.path)
    .use(this.defaults)
    // FIXME
    if (params.send) {
      req.send(params.send)
    } else if (params.set) {
      req.set(params.set)
    } else if (params.query) {
      req.query(params.query)
    }
    return req
  },
  defaults (req) {
    // default request settings
    return function (req) {
      req.set('Accept', 'application/json')
      return req
    }
  },
  prefix (prefix) {
    // copied from superagent-prefix
    return function (req) {
      if (req.url[0] === '/') {
        req.url = prefix + req.url
      }
      return req
    }
  },
  authReq (req) {
    return function (req) {
      let header = api.authHeader
      let bearer = Storage.retriveToken()
      return req.set(`${header}`, `Bearer ${bearer}`)
    }
  }
}
