<script setup>
import { ref } from 'vue';

const movies = ref([]);  

    fetch('http://localhost:8000/watchedmovies',{
        method: 'POST',
    })
    .then(response => response.json())
    .then(data => {
        movies.value = data
        console.log(data)
    })
    .catch(error => {
        console.error('Error:', error);
    });
</script>




<template>

    <v-container class="sheet">
        <v-list>
            <h1>Your Watchlist</h1>
            <v-list-item v-for="movie in movies" :key="movie.MovieId"> 
                <v-list-title>1.<span>{{ movie.Title }}</span></v-list-title>
                <br></br>
                <v-list-subtitle>{{ movie.Comment }}</v-list-subtitle>
            </v-list-item>
        </v-list>
    </v-container>


</template>



<style scoped>
.sheet{
    max-width: 50%;
    text-align: center;
}

span{
    margin-left: 10px;
}


</style>