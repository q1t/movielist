<template>
  <div class="login">
    <div>
      <h3> Login into your account </h3>
      <div v-show="showError" class="errors">
        <ul>
          <li v-for="e in errorMessages">
            <span>{{ e }}</span>
            </li>
          </ul>
      </div>
      <div class="username">
        <input type="text" name="username" v-model="username" 
        placeholder="username"
        >
      </div>
      <div class="password">
        <input type="password" name="password" v-model="password"
        placeholder="password"
        >
      </div>
      <div class="submit">
        <button @click.prevent="loginUser" type="submit"> Login </button>
      </div>
    </div>
  </div>
</template>
<script>
import API from '../javascript/api.js'
  export default {
    data () {
      return {
        username: '',
        password: '',
        showError: false,
        errorMessages: []
      }
    },
    methods: {
      loginUser(){
        API.LoginUser({
          username: this.username, 
          password: this.password 
        })
        .then(API.CheckForBearer)
        .then((d) => {
          this.showError = false
          this.$dispatch('recive-auth-token', d.bearer)
        })
        .catch((err) => {
          this.handleError(err)
        })
      },
      handleError(err) {
          this.showError = true
          if (err.error) {
            this.errorMessages = [err.error]
          }
          if (err.errors) {
            this.errorMessages = err.errors;
          }
      }
    }
  }
</script>
<style lang="postcss">
$gkkb: rgba(255, 1, 142, 1);
  .login {
    lost-utility: clearfix;
    height: 100%;
    & > div:first-child {
      padding-right: 50px;
      padding-left: 50px;
      padding-bottom: 50px;
      padding-top: 10px;
    }
  }
  .login
  .errors {
    ul {
      list-style: none;
      padding: 0;
      font-size: 14px;
      li {
        margin-bottom: 2px;
        span {
          border-bottom: 1px solid; 
          border-color: red;
        }
      }
    }
  }
  .login
  .username, 
  .login 
  .password {
    lost-utility: clearfix;
    lost-row: 1/3 10px;
    & > input {
      lost-column: 1/1;
      box-sizing: border-box;
      display: block;
      padding: 10px;
      border-radius: 2px;
      border: none;
    }
  }
  .login
  .submit {
    lost-row: 1/3;
  }

  .login .submit button {
    border-radius: 2px;
    width:100%;
    background:color($gkkb shade(8%));
    border:0;
    padding:10px;
    color:#fff;
    cursor:pointer;
    transition:background .3s;
    -webkit-transition:background .3s;
    font-weight: bold;
  }

  .login .submit button:hover{  
    background: $gkkb;
  }
</style>
