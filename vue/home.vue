<template>
  <div id="home">
    <div v-show="!loggedin" class="inlined">
      <login></login>
      <register></register>
    </div>
    <div class="home-page markdown-body">
      {{{ page }}}
    </div>
  </div>
</template>
<script>
import Login from './login.vue'
import Register from './register.vue'
import how from '../docs/howtouse.md'
import Storage from '../javascript/storage.js'

export default {
  name: 'home',
  props: {
    loggedin: {
      type: Boolean,
      required: true
    }
  },
  data(){
    return {
      page: null
    }
  },
  created(){
    this.page = how
  },
  components: {
    Login,
    Register
  }
}
</script>

<style lang="postcss">
.inlined {
  lost-utility: clearfix;
  lost-row: 1/1;
  background: blue;
  color: white;
  border-bottom: 2px solid;
  border-color: color(#0074D9 shade(11%));
}

.inlined:first-child {
  & > div {
    lost-column: 1/2;
    &:first-child {
    }
  }
}
</style>
