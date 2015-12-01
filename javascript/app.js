import '../styles/app.css' // css is bundled within js file
import 'babel-polyfill'
import 'marked'

import Vue from 'vue'
import jwt_decode from 'jwt-decode'
import Storage from '../javascript/storage.js'
import VueRouter from 'vue-router'
import VueAsyncData from 'vue-async-data'

Vue.use(VueRouter)
Vue.use(VueAsyncData)

// debug mode
Vue.config.debug = true

var router = new VueRouter({
  history: true // enable HTML5 history (pushstate)
})

// Components
import Home from '../vue/home.vue'
import ListsView from '../vue/lists_view.vue'
import ListView from '../vue/list_view.vue'

router.map({
  '/': {
    name: 'home',
    component: Home
  },
  '/lists': {
    name: 'lists',
    component: ListsView
  },
  'lists/:title': {
    name: 'list',
    component: ListView
  }
})

var app = Vue.extend({
  name: 'Main',
  created () {
    let token = Storage.retriveToken()
    if (token) {
      this.loggedin = true
    }
  },
  computed: {
    user () {
      return jwt_decode(Storage.retriveToken())
    }
  },
  data () {
    return {
      loggedin: false
    }
  },
  events: {
    'recive-auth-token': function (bearer) {
      Storage.saveToken(bearer)
      this.loggedin = true
    }
  },
  methods: {
    logoutUser () {
      Storage.removeToken()
      this.loggedin = false
    },
    userInf () {
      return jwt_decode(Storage.retriveToken())
    }
  }
})
router.start(app, '#app')
