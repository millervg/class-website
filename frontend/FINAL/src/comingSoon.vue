<script setup>
import { ref } from 'vue';

import watchlistFetch from './components/watchlistFetch.vue';

const movieName = ref("")
const movieRating = ref("")
const movieComment = ref("")

function addMovie() {
    const movies = {
        Title: movieName.value.trim(),
        Rating: movieRating.value.trim(),
        Comment: movieComment.value.trim()
    };

    fetch('http://localhost:8000/addwatchedmovie', {
        method:'POST',
        body:JSON.stringify(movies)

    })
    .then(response => response.json())
    .then(data => {
        console.log(data)
    })

    movieName.value ='';
    movieRating.value ='';   
    movieComment.value ='';
}

</script>

<template>
  <v-carousel show-arrows="hover" hide-delimiters class="banner" >
    <v-carousel-item src="../images/Oppenheimer.png" cover>
      <div class="text">
      <h1>Watchlist</h1>
    </div>
    </v-carousel-item>
    </v-carousel>




<div class="grid">
    <v-sheet class="mx-auto" width="300">
    <v-form @submit.prevent>
      <v-text-field v-model="movieName" label="Movie Title"></v-text-field>
      <v-text-field v-model="movieRating" label="Rating"></v-text-field>
      <v-text-field v-model="movieComment" label="Comment"></v-text-field>
      <v-btn @click="addMovie" >Add Movie</v-btn>
    </v-form>

  </v-sheet>

    <watchlistFetch></watchlistFetch>

</div>


</template>





<style scoped>

*{
    margin: 0;
    padding: 0;
}

.grid{
    margin-bottom: 50px;
    margin-top: 50px;
    padding: 10px;
    display: flex;
    gap: 20px;
}

h1{
    color: white;
    text-align: center;
    margin-top: 200px;
    text-transform: uppercase;
    font-size: 60px;
}





</style>