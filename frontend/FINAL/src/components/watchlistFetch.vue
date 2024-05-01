<script setup>
import { ref } from 'vue';

const movies = ref([]);  
    fetch('http://localhost:8000/watchedmovies',{
        method: 'POST',
    })
    .then(response => response.json())
    .then(data => {
        movies.value = data
        console.log(movies.WatchedMovieId)

        console.log(data)
    })
    .catch(error => {
        console.error('Error:', error);
    });

</script>

<template>
    <v-container class="sheet2">
      <v-table class="list">
        <caption>
          <h1>
            <v-icon>mdi-list-box</v-icon>
            Your Watchlist</h1>
        </caption>
        <thead>
          <tr>
            <th style="text-align: center;">Movie Title</th>
            <th style="text-align: center;">Ratings</th>
            <th style="text-align: center;">Comments</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="movie in movies" :key="movie.MovieId">
            <td>{{ movie.Title }}</td>
            <td>{{ movie.Rating }}/10</td>
            <td>{{ movie.Comment }}</td>
          </tr>
        </tbody>
      </v-table>
    </v-container>
  </template>
  
<style scoped>
  .sheet2 {
    max-width: 100%;
    text-align: center;
    margin-bottom: 20px;
    margin-top: 50px;

    
  }

  .list:hover{
    box-shadow: 0px 0px 10px 0px rgba(0,0,0,0.75);

  }
  
  .list {
    text-align: center;
    color: black;
    background-color: white;
    margin-top: 50px;
    margin-left: 50px;
    margin-right: 60px;
    margin-bottom: 20px;

    border: 3px solid white;
    border-radius: 20px;
  }
  
  h1 {
    margin-bottom: 10px; 
    margin-top: 10px;
  }
</style>
  