<template>
    <div class="lists" :class="{ loading: !lists.length }">
        <div class="add-list">
            <div class="add-list-button">
                <button @click="showForm = true"> New List </button>
            </div>
            <div v-if="showForm" class="add-list-form">
                <input type="text" placeholder="title" v-model="newList.title"> </input>
                <input type="checkbox" v-model="newList.isPublic"> public </input>
                <button @click="createList"> Add </button>
            </div>
        </div>
        <list
            v-for="list in lists"
            :list="list"
        >
        </list>
    </div>
</template>
<script>
 import List from './list.vue'
 import API from '../javascript/api.js'
 export default {
     name: 'ListsView',
     data(){
         return {
             lists: [],
             showForm: false,
             newList: {
               title: '',
               isPublic: false
             }
         }
     },
     methods: {
       createList () {
         API.AddList(this.newList).then((r) => {
           this.newList.title = ''
           this.newList.isPublic = false
           this.showForm = false
           this.reloadAsyncData()
         })
       }
     },
     asyncData() {
        API.UserLists().catch((e) => {
          console.log(e)
        })
        .then((r) => {
          let lists = JSON.parse(r.text)
          this.lists = lists
        })
     },
     components: {
         List
     }

     //route: {
     //    activate() {
     //        console.log("activated")
     //    },
     //    candeactivate(transition) {
     //        console.log(transition)
     //    }
     //}
 }

 // grab this data from server
 var mock = [
     {
         title: 'blockbasters',
         titles: 113,
         watched: 17,
         labels: ['test', 'mustwatch', 'awesome'],
         public: true
     }
 ]
</script>
<style lang="postcss">
  .add-list {
      height: 40px;
      margin-bottom: 10px;
      lost-utility: clearfix;
      .add-list-button {
        lost-column: 1/8;
        button {
            padding: 5px 10px;
            background: olive;
            border: none;
            font-size: 14px;
            color: white;
        }
        button:hover {
            cursor: pointer;
        }
      }
  }
  .add-list-form {
    height: 100%;
    font-size: 14px;
    lost-column: 5/6;
    & > input {
            font-size: 14px;
            padding: 3px 10px;
            margin-left: 5px;
    }
    & > button {
            padding: 5px 10px;
            background: olive;
            border: none;
            font-size: 14px;
            margin-left: 5px;
            color: white;
    }
    & > button:hover {
         cursor: pointer;
    }
  }

  .fade-transition {
      transition: max-height 0.1s;
  }
  .fade-enter {
    max-height: 0;
  } 
  .fade-leave {
    max-height: auto;
  }
 .lists {
     lost-center: 720px;
     lost-utility: clearfix;
     padding-bottom: 20px;
     padding-top: 20px;
 }
</style>
