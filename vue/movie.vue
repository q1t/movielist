<template>
    <div class="movie">
    	<modal :show.sync="showCard" :movie="movie">
		</modal>
    	<div class="poster" @click="showCard = true" :style="{
    		backgroundImage: bgUrl(movie.poster) 
    	}">
	 		<div class="info">
		    	<div class="title">
		    		{{ movie.title }}	
		    	</div>
		    	<div class="year">
		    		{{ movie.year }}	
		    	</div>
		    	<div class="genre">
		    		{{ movie.genre }}	
		    	</div>
<!-- 		    	<div id="delimeter"></div>
 -->		    	<div class="ratings">
		    		<div class="metascore source" 
		    			>
		    			<label for="rating">Metascore:</label>	
		    			<div class="rating">{{ movie.metascore }}</div>
		    		</div>
		    		<div class="imdb source" 
		    			>
		    			<label for="rating">IMDb:</label>	
		    			<div class="rating"><span>{{ movie.imdbRating }}</span> </div>
		    		</div>		
					<!-- TODO:FIXME:WTF -->
		    		<div class="tomatoes source" 
		    			>
		    			<label for="rating">Tomatoes:</label>	
		    			<div class="rating" 
		    				>
		    				{{ movie.tomatoRating }}
		    			</div>
		    			<div class="rating" 
		    				v-show="showTomatoesUserRating"
		    				>
		    				User: {{ movie.tomatoUserRating }}
		    			</div>
		    		</div>
		    	</div>
	    	</div>
    	</div>
    </div>
</template>

<script>
import Modal from './modal_movie_card.vue'

export default {
    name: 'Movie',
    props: {
        movie: Object
    },
    data() {
    	return {
    			showCard: false,
    			showTomatoesUserRating: false,
    			RatingSources: {
	    			imdb: 'imdbRating',
	    			metascore: 'metascore',
	    			tomatoes: {
	    				rating: 'tomatoRating',
	    				userRating: 'tomatoUserRating'
	    			}
    			}
    		}
    },
    methods: {
    	bgUrl(url) {
    		// use proxy prefix to retrive IMDB poster image
    		// from cache on the server 
    		return `url(/proxy/${url})` 
    	},
    	showMovieCard() {
    		this.showCard = true
    	}
    },
    components: {
    	Modal
    }
}
</script>

<style lang="postcss">
	.movie {
		lost-waffle: 1/3 auto 60px;
		color: black;
		cursor: pointer;
	}

	.movie
	.poster {
		border-radius: 2px;
		background-position: center;
		background-repeat: no-repeat;
		background-size: cover;
		max-height: 300px;
		min-height: 280px;
		position: relative;
    border: 1px solid #C5A8A8;
	}

	.movie
	.info {
		border-bottom-right-radius: 2px;
		border-bottom-left-radius: 2px;
		text-align: center;
		font-size: 14px;
		position: absolute;
		left: 0;
		bottom: 0;
		height: auto;
		width: 100%;
		transition: all 0.8s;
		margin-top: 10px;
		.title {
			margin-top: 5px;
		}
		color: silver;
		background: rgba(0, 1, 2, 0.72);
	}
	.movie
	.info
	.genre {
    /* -- */
	}

	.movie
	.info
	.ratings {
		lost-utility: clearfix;
		font-size: 10px;	
		padding-bottom: 4px;
	}
	.movie
	.info
	.ratings
	.source  {
		font-style: italic;
		text-align: center;
		margin-top: 3px;
		lost-utility: clearfix;
		lost-column: 1/3 0;
		.rating {
			font-style: normal;
			font-weight: 800;
		}
	}
</style>
