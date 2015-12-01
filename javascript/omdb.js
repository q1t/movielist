import request from 'superagent'
import prefix from 'superagent-prefix'
import camel from 'to-camel-case'
import qs from 'query-string'

// https://esdiscuss.org/topic/es6-iteration-over-object-values
var entries = function* (obj) {
  for (let key of Object.keys(obj)) {
    yield [key, obj[key]];
  }
}
const ApiParameters = {
  imdb: 'i',
  title: 't',
  type: 'type', // movie, series, episode
  year: 'y',
  plot: 'plot', // short, full
  return_type: 'r', // json, xml
  tomatoes: 'tomatoes', // true, false
  callback: 'callback', // JSONP callback
  version: 'v' // API version, default 1
}
const DefaultParams = {
  imdb: '',
  title: '',
  type: '',
  year: '',
  plot: 'short',
  return_type: 'json',
  tomatoes: 'true',
  callback: '',
  version: '1'
}
class Omdb {
  constructor (apiurl = 'http://www.omdbapi.com/') {
    this.apiurl = apiurl
  }
  parseMovie (response) {
    return new Promise((resolve, reject) => {
      resolve(response.map(field => {
        let movie = {}
        for (let [prop, value] of entries(JSON.parse(field.text))) {
          movie[camel(prop)] = value
        }
        return movie
      }))
    })
  }
  newOptions (params) {
    let p = params || DefaultParams

    var opts = {}
    for (let param in ApiParameters) {
      opts[ApiParameters[param]] = p[param] || DefaultParams[param]
    }
    return opts
  }
  get (params) {
    return new Promise((resolve, reject) => {
      // dirty hack
      request.get(this.apiurl + '?' + qs.stringify(params))
      .end((err, resp) => {
        if (!err) {
          resolve(resp)
        } else {
          reject(err)
        }
      })
    })
  }
}
export default new Omdb()
