<script setup>
import { ref } from 'vue';
import movieFetch from './components/movieFetch.vue';
</script>

<template>
      <movie-fetch></movie-fetch>
      <div class="page">
            <ul>
                  <li><a href="/filmcatalog2">1</a></li>
                  <li><a href="#/home">2</a></li>
                  <li><a href="#/home">3</a></li>
                  <li><a href="#/home">4</a></li>

            </ul>
      </div>
</template>

<style scoped>
ul {
      list-style: none;
}

li{
      text-align: center;
      display: inline-block;
      margin-bottom: 10px;
      padding: 15px;
      border: 1px solid black;


}

li a {
  display: block;
  padding: 5px 10px;
  color: black; /* Change text color to black */
  text-decoration: none;
}

.page{
      text-align: center;
}
</style>
  


  