<template>
  <div class="watch-list">
    <div class="spinner" v-show="$loadingAsyncData">
      <i class="fa fa-spinner fa-spin fa-lg"></i>
    </div>
    <!--         here is the list of movies with title <b>{{ $route.params.title }}</b>
    -->

    <div class="upper-bar" v-show="!$loadingAsyncData">
      <div class="watch-list-settings">
          <input type="radio"> Sort by adding date </input>
      </div>
      <div class="add-movie-form">
          <input type="text" name="movie-title"
          placeholder="start typing movie title"
          v-model="search"
          >
          <button @click="suggestMovie()"> Add movie </button>
      </div>
    </div>
    <div v-if="!$loadingAsyncData" class="movies">
      <movie
      v-for="movie in movies"
      transition="fade"
      :movie="movie"
      >
      </movie>
    </div>
  </div>
</template>
<script>
  import Movie from './movie.vue'
  import omdb from '../javascript/omdb.js'
  import request from 'superagent'

  // Some mockup data
  var data = [
    {
      imdb: 'tt2379713',
      user_rating: 8.7
    },
    {
      imdb: 'tt2249007',
      user_rating: 8.7
    },
    {
      imdb: 'tt3659388',
      user_rating: 1.0
    },
    {
      imdb: 'tt0185906',
      user_rating: 8.7
    },
    {
      imdb: 'tt0903747',
      user_rating: 8.7
    },
    {
      imdb: 'tt0944947',
      user_rating: 1.0
    }
  ]

  function suggestMovie(sq) {
    let resource = '/proxy/http://sg.media-imdb.com/suggests'
    if (sq.length > 0) {
        let fl = sq.charAt(0)
        let path = `${resource}/${fl}/${sq}.json`
        r.get(path)
        .jsonp()
        .end((res, err) => {
            console.log(res,err)
        })
    }
    return {}
  }

  export default {
    name: 'ListView',
    data() {
      return {
        search: '',
        movies: [],
        mokeData: data,
        show: true
      }
    },
    asyncData() {
      var requests = []
      this.mokeData.forEach(um => {
        requests.push(omdb.get(omdb.newOptions({ imdb: um.imdb })))
      })
      return Promise.all(requests)
      .then(omdb.parseMovie)
      .then(data => {this.movies = data})
    },
    methods: {
      fetchUserMovies(){
        return data
      },
      suggestMovie () {
        suggestMovie(this.search)
      }
    },
    components: {
      Movie
    }
  }
</script>
<style lang="postcss">
  .fade-transition {
    transition: opacity .3s ease-in;
  }
  .fade-enter, .fade-leave {
    opacity: 0;
  }
  .spinner {
    lost-center: 720px;
    text-align: center;
    margin-bottom: 20px;
    margin-top: 20px;
  }
  .movies {
    margin-top: 20px;
    height: 100%;
    lost-center: 720px;
    padding-bottom: 20px;
  }
  .watch-list
  .upper-bar {
    lost-utility: clearfix;
    lost-center: 720px;
    padding-top:10px;
  }
  .watch-list-settings {
    lost-column: 1/2;
    font-size: 14px;
  }
  .add-movie-form {
    lost-column: 1/2;
    & > input {
            padding: 3px 10px;
            margin-left: 5px;
            width: 65%;
      box-sizing: border-box;
    }
    & > button {
            padding: 5px 10px;
            background: olive;
            border: none;
            margin-left: 5px;
            color: white;
            width: 30%;
    }
    & > button:hover {
         cursor: pointer;
    }
  }
</style>
