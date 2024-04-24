<script setup>
import { ref } from 'vue';


const movies = ref([]);

    fetch('http://localhost:8000/movies',{
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
  <div class="header">
    <h1 >Featured Movies</h1>
  </div>
  <div class="container">
    <v-card v-for="movie in movies" :key="movie.MovieId" class="card">
      <div class="card-content">
        <v-img :src="`../../public/images/${movie.ImageName}`" class="card-image"></v-img>
        <div class="text-overlay">
          <v-card-title class="card-title">{{ movie.Title}}</v-card-title>
        </div>
      </div>
    </v-card>

  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css?family=Poppins:100,200,300,400,500,600,700,800,900&display=swap');

*{
    margin: 0;
    padding: 0;
    font-family: 'Poppins', sans-serif;

}

.header h1{
    display: inline-block;
    border: 5px solid black;  
    padding: 10px;
    border-radius: 10px;

}

.header{
  text-align: center;
  margin-top: 10px;
  padding: 20px; 

}

.header h1:hover {
  box-shadow: 6px 6px black;
}

.container {
  margin-top: 10px;
  display: grid;
  grid-template-columns: auto auto auto auto auto auto;
  justify-content: center;
}

.card:hover {
  box-shadow: 6px 6px black;
}

.card {
  position: relative;
  text-align: center;
  width: 190px;
  height: 290px;
  margin-left: 15px;
  padding: 0;
  border-radius: 15px;
  margin-bottom: 10px;

}

.card-content {
  position: relative;
  width: 100%;
  height: 100%;
}

.card-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.text-overlay {
  position: absolute;
  bottom: 0;
  width: 100%;
  color: black;
}
.card-title,
.card-text {
  margin-bottom: 10px;
  text-decoration:solid;
  font-weight: bolder;
  text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.5);
}
</style>
