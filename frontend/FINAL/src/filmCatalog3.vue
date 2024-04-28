<script setup>
import { ref } from 'vue';
import movieFetch from './components/movieFetch.vue';


</script>

<template>
        <movie-fetch></movie-fetch>
      <div class="pages">
            <ul>
                  <li><a href="#/filmCatalog">1</a></li>
                  <li><a href="#/filmCatalog2">2</a></li>
                  <li><a href="#/filmCatalog3">3</a></li>
            </ul>
      </div>
</template>


<style scoped>
ul {
      list-style: none;
}

li {
      text-align: center;
      display: inline-block;
      margin-bottom: 10px;
      padding: 15px;
      border: 2px solid white;
      margin-left: 10px;
      border-radius: 30px;
      background-color: white;
}

li a {
  padding: 5px 10px;
  color: black; 
  text-decoration: solid;
  text-transform: uppercase;
  font-weight: bolder;
  font-size: larger;
  
}

.pages{
      text-align: center;
}
</style>
  


  