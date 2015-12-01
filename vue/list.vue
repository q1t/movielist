<template>
<div v-if="show" class="list">
    <div class="title">
        <a v-link="{name: 'list', params: { title: list.title }}">{{ list.title }}</a>
    </div>
    <div class="watched-n-all">
        <label> <span>all/watched</span> </label>
        <div class="all">
            {{ list.titles }}
        </div>
        <div class="watched">
            {{ list.watched }}
        </div>
    </div>
    <div class="labels-n-lastfive">
        <div class="labels">
            <div class="label bg-blue" v-show="list.public">
                <i class="fa fa-globe"></i>
            </div>
            <div class="label bg-green" v-for="label in list.labels">
                {{ label }}
            </div>
        </div>

        <div class="lastfive">
            <label for="lastfive">
                <span>Last added movies:</span>
            </label>
            <ul>
                <li class="movie-title">
                    <span> Test movie </span>
                </li>
                <li class="movie-title">
                    <span> Another Test movie </span>
                </li>
                <li class="movie-title">
                    <span> Test movie </span>
                </li>
                <li class="movie-title">
                    <span> Test movie </span>
                </li>
            </ul>
        </div>
    </div>
    <div @click="deleteList()" class="settings">
        <i class="fa fa-trash-o"></i>
    </div>
</div>
</template>
<script>
  import API from '../javascript/api.js'
 export default {
     name: 'List',
     props: {
         list: Object
     },
     data () {
       return {
         show: true
       }
     },
     methods: {
       deleteList() {
         if (confirm(`Delete list ${this.list.title}?`)) {
           API.DeleteList(this.list.id).then((r) => {
            this.show = false
           })
         }
       }
     }

 }
</script>
<style lang="postcss">
 $action-color: color(#3D9970 shade(9%));
 @lost gutter 0px;
 .list {
     lost-utility: clearfix;
     color: silver;
     border-radius: 2px;
     lost-row: 1 20px 20px;
     background: olive;
     font-size: 16px;
     height: 50px;
 }

 .list
 a {
     color: silver;
     text-decoration: none;
     font-size: 1em;
     font-style: italic;
     text-transform: capitalize;
     display: block;
     height: 100%;
 }

 .list
 > div {
     text-align: center;
     height: 100%;
     line-height: 3;
 }

 .list
 > div:last-child {
     border-right: none;
 }

 .list
 .title {
     lost-column: 3/12;
     background: $action-color;
 }

 .list
 .title:hover {
     background: olive;
 }

 .list
 .watched-n-all {
     lost-column: 1/12;
     line-height: 1;
     font-size: 12px;
 }


 .list
 .watched-n-all
 label {
     margin-top: 4px;
     display: inline-block;
     font-size: 8px;
 }

 .list
 .watched-n-all
 .all {
     margin-top: 4px;
     border-bottom: 1px solid;
     border-color: olive;
 }

 .list
 .watched-n-all
 .all:after {
     content: ' ';
     display:block;
     width: 8px;
     border-bottom: 1px solid color(silver shade(15%));
     text-align: center;
     margin-left: 45%;
     margin-top: 2px;
 }

 .list
 .watched-n-all
 .watched {
     margin-top:1px;
 }




 .list
 .labels-n-lastfive {
     background: silver;
     lost-column: 7/12;
 }

 .list
 .labels-n-lastfive
 .labels {
     float: left;
     line-height: 0.8;
     width: 100%;
 }

 .list
 .labels-n-lastfive
 .labels
 .label {
     border-bottom-right-radius: 2px;
     border-bottom-left-radius: 2px;
     display: inline;
     padding: 1px 3px 2px 3px;
     font-size: 76%;
     font-weight: bold;
     line-height: 1;
     color: #fff;
     text-align: center;
     white-space: nowrap;
     vertical-align: baseline;
     margin-left: 10px;
     float: left;
 }

 .list
 .labels-n-lastfive
 .lastfive {
     color: #333;
     font-size: 14px;
     font-style: italic;
     height: 70%;
     margin-top: 16px;
     line-height: 1.7;
     text-align: left;
 }

 .list
 .labels-n-lastfive
 .lastfive
 ul {
     display: inline;
     margin: 0;
     padding: 0;
     margin-left: 10px;
 }

 .list
 .labels-n-lastfive
 .lastfive
 label {
     display: block;
     height:10px;
     margin-left: 10px;
     font-size: 10px;
 }

 .list
 .labels-n-lastfive
 .lastfive
 .movie-title {
     margin-right: 2px;
     display: inline-block;
     font-size: 12px;
     margin-top: 3px;
     list-style: none;
 }

 .list
 .labels-n-lastfive
 .lastfive
 .movie-title:after {
     content: ", ";
 }


 .list
 .labels-n-lastfive
 .lastfive
 .movie-title:last-child:after {
     content: "";
 }

 .list
 .settings {
     background: $action-color;
     float: right;
     lost-column: 1/12;
 }

 .list
 .settings:hover {
     background: olive;
     cursor: pointer;
 }
 .bg-blue {
     background: blue;
 }
 .bg-green {
     background: green;
 }
</style>
