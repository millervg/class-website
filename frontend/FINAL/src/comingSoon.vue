<script setup>
import { ref } from 'vue';

import watchlistFetch from './components/watchlistFetch.vue';

const movieName = ref("")
const movieRating = ref("")
const movieComment = ref("")

function addMovie() {
    const movies = {
        Title: movieName.value.trim(),
        Rating: parseInt(movieRating.value.trim()),
        Comment: movieComment.value.trim()
    };

    fetch('http://localhost:8000/addwatchedmovie', {
        method:'POST',
        body:JSON.stringify(movies)

    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
    })
    .catch(error => {
        console.error('Error:', error);
    });

        movieName.value ='';
        movieRating.value ='';
        movieComment.value ='';

        location.reload()


}

</script>

<template>
    <v-carousel hide-delimiters hide-controls class="banner">
    <v-carousel-item src="../images/Oppenheimer.png" cover>
      <div class="text">
        <h1>Watchlist</h1>
      </div>
    </v-carousel-item>

  </v-carousel>
  <div>
    <btn><a href="#/home">Click me</a></btn>
  </div>
<div class="grid">
    <aside class="form1">

        <v-sheet class="sheet">
             <v-form class="form" @submit.prevent>
                <h2>
                    <v-icon>mdi-list-box</v-icon>
                    Add To Watchlist
                </h2>
                <v-text-field v-model="movieName" label="Movie Title" ></v-text-field>
                <v-text-field v-model="movieRating" label="Rating(0/10)" ></v-text-field>
                <v-text-field v-model="movieComment" label="Comment"></v-text-field>
                <v-btn class="btn" style="background-color:darkslategray ;" @click="addMovie"  block>Add Movie</v-btn>
            </v-form>
        </v-sheet>
    </aside>
    <aside >
        <watchlistFetch></watchlistFetch>
    </aside>
</div>


</template>
<style scoped>
*{
    margin: 0;
    padding: 0;
}

.background {
  background-color: white;
  text-align: center;
  border-radius: 10px;
  margin: 20px auto;
  width: fit-content;
}

.background:hover{
    color: white;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

.fade-in {
  animation: fadeIn 5s ease-in-out;
  text-align: center;
  padding: 20px;
  
}


.sheet{
    width: 50%;
    margin-left: 200px;
    color: black;
    border-radius: 30px;
    margin-bottom: 20px;

}

.form:hover{
    box-shadow: 0px 0px 10px 0px rgba(0,0,0,0.75);
}

.form{
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

h2{
    text-align: center;
    margin-bottom: 10px;
}

h1{
    color: white;
    text-align: center;
    margin-top: 200px;
    text-transform: uppercase;
    font-size: 60px;
}

</style>