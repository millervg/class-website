<script setup>
import { ref } from 'vue';
import watchlistFetch from './components/watchlistFetch.vue';

const movieName = ref('');
const movieRating = ref('');
const movieComment = ref('');


function addMovie() {
    const movies = {
        Title: movieName.value.trim(),
        Rating: parseInt(movieRating.value.trim()),
        Comment: movieComment.value.trim()
    };

    fetch('http://localhost:8000/addwatchedmovie', {
        method: 'POST',
        body: JSON.stringify(movies)
    })
    .then(response => response.json())
    .then(data => {
    })
    .catch(error => {
        console.error('Error:', error);
    });

    movieName.value = '';
    movieRating.value = '';
    movieComment.value = '';
}
</script>

<template>
   <div>
    <div class="btn-wrapper">
      <btn>
        <h3 class="btn-text">
          <v-icon>mdi-movie-filter</v-icon>
          Add to your watchlist!
        </h3>
      </btn>
    </div>    
    <div class="grid">
      <aside class="form1">
        <v-sheet class="sheet">
          <v-form class="form" @submit.prevent>
            <h2>
              <v-icon>mdi-list-box</v-icon>
              Add To Watchlist
            </h2>
            <v-text-field v-model="movieName" label="Movie Title"></v-text-field>
            <v-text-field v-model="movieRating" label="Rating(0/10)"></v-text-field>
            <v-text-field v-model="movieComment" label="Comment"></v-text-field>
            <v-btn class="btn" style="background-color:darkslategray;" @click="addMovie" block>Add Movie</v-btn>
          </v-form>
        </v-sheet>
      </aside>
      <aside>
        <watchlistFetch></watchlistFetch>
      </aside>
    </div>
  </div>

</template>

<style scoped>
.btn-text:hover{
  font-size: larger; 
  text-transform: uppercase; 
  color: black; 
  background-color: #feb47b; 
  border: 2px solid #ff7e5f; 
  border-radius: 10px; 
  padding: 12px 24px; 
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;

}

.btn-wrapper {
  display: flex; 
  justify-content: center;
  align-items: center;
  padding: 10px;
  border-radius: 5px;
  margin-bottom: 10px;
  margin-top: 35px;
}

.btn-text {
  background-color: white; 
  padding: 10px; 
  border-radius: 5px; 
  text-align: center;
}


h2 {
  text-align: center;
  margin-bottom: 10px;
}

h1 {
  color: white;
  text-align: center;
  margin-top: 100px;
  text-transform: uppercase;
  font-size: 60px;
}

.background {
  background-color: white;
  text-align: center;
  border-radius: 10px;
  margin: 20px auto;
  width: fit-content;
}

.sheet {
  width: 50%;
  margin-left: 200px;
  color: black;
  border-radius: 30px;
  margin-bottom: 20px;
}

.form:hover {
  box-shadow: 0px 0px 10px 0px rgba(0,0,0,0.75);
}

.form {
  margin-top: 90px;
  padding: 20px;
  border-radius: 30px;
  border: 2px solid white;
}

.grid {
  display: flex;
}

aside {
  flex: 1;
}
</style>
