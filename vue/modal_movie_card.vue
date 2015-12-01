<template>
  <div class="modal-mask" v-show="show" transition="modal">
    <div class="modal-wrapper">
      <div class="modal-container">
        <div class="modal-header">
          <slot name="header">
            {{ movie.title }}
          </slot>
        </div>

        <div class="modal-body">
          <slot name="body">
            {{ movie.plot }}
          </slot>
        </div>

        <div class="modal-footer">
          <slot name="footer">
            {{ movie.production }}
            |
            <a :href="movie.website">Movie Website</a>
            <button class="modal-default-button"
            @click="show = false">
            Close
          </button>
        </slot>
      </div>
    </div>
  </div>
</div>
</template>
<script type="text/javascript">
export default {
  name: 'MovieCard',
  props: {
    show: {
      type: Boolean,
      required: true,
      twoWay: true
    },
    movie: {
      type: Object
    }
  }
}
</script>
<style lang="postcss">
.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(10, 10, 10, .97);
  display: table;
  transition: opacity .3s ease;
}

.modal-wrapper {
  display: table-cell;
  vertical-align: middle;
}

.modal-container {

  width: 300px;
  margin: 0px auto;
  margin-top: -100px;
  padding: 20px 30px;
  background-color: #fff;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, .33);
  transition: all .3s ease;
  font-family: Helvetica, Arial, sans-serif;
}

.modal-header h3 {
  margin-top: 0;
  color: #42b983;
}

.modal-body {
  margin: 20px 0;
}

.modal-default-button {
  float: right;
}
.modal-enter, .modal-leave {
  opacity: 0;
}

.modal-enter .modal-container,
.modal-leave .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}
</style>
