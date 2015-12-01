<template>
  <div class="register">
    <div>
    <h3>Or register a new one</h3>
      <div v-show="showError" class="errors">
        <ul>
          <li v-for="e in errorMessages">
            <span>{{ e }}</span>
            </li>
          </ul>
      </div>
    <div class="username">
      <!-- <label for="email">email</label> -->
      <input type="username" name="username" v-model="username" 
      placeholder="username"
      >
    </div>
    <div class="password">
      <input type="password" name="password" v-model="password"
      placeholder="password"
      >
    </div>
    <div class="submit">
      <button @click.prevent="registerUser" type="submit">Create account</button>
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
        email: '',
        password: '',
        showError: false,
        errorMessages: []
      }
    },
    methods: {
      registerUser () {
        API.RegisterUser({
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
  .register {
    lost-utility: clearfix;
    height: 100%;
    & > div:first-child {
      padding-right: 50px;
      padding-left: 50px;
      padding-bottom: 50px;
      padding-top: 10px;
    }
  }
  .register
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
  .register
  .email , 
  .register
  .username,
  .register
  .password {
    lost-utility: clearfix;
    lost-row: 1/2 10px;
    & > input {
      lost-column: 1/1;
      box-sizing: border-box;
      display: block;
      padding: 10px;
      border-radius: 2px;
      border: none;
    }
  }
  .register
  .submit {
    lost-row: 1/2;
  }
.register .submit button {
  border: 0;
  width:100%;
  background:color($gkkb shade(8%));
  border-radius: 2px;
  padding:10px;
  color:#fff;
  cursor:pointer;
  transition:background .3s;
  -webkit-transition:background .3s;
  font-weight: bold;
}

.register .submit button:hover{
  background: $gkkb;
}
</style>
