export default {
  saveToken (token) {
    this.saveItem('token', token)
    return token
  },
  removeToken () {
    return this.deleteItem('token')
  },
  // no input :(
  retriveToken () {
    return this.retriveItem('token')
  },
  // generic interface
  retriveItem (key) {
    return localStorage.getItem(key)
  },
  deleteItem (key) {
    return localStorage.removeItem(key)
  },
  saveItem (key, value) {
    return localStorage.setItem(key, value)
  },
  wipe () {
    return localStorage.clear()
  }
}
