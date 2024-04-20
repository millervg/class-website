<script setup>
import { ref, computed } from 'vue';
import filmCatalog from './filmCatalog.vue';
import reviews from './reviews.vue';
import comingSoon from './comingSoon.vue';
import home from './home.vue';

const routes = {
    '/filmCatalog': filmCatalog,
    '/reviews': reviews,
    '/comingSoon': comingSoon,
    '/home': home

}

const currentPath = ref(window.location.hash)

window.addEventListener("hashchange", () => {
    currentPath.value = window.location.hash
})

const currentView = computed(() => {
    return routes[currentPath.value.slice(1) || '/'] 
})

const drawer = ref(false);
</script>


<style scoped>


@import url('https://fonts.googleapis.com/css?family=Poppins:100,200,300,400,500,600,700,800,900&display=swap');


*{
    margin: 0;
    padding: 0;
    font-family: 'Poppins', sans-serif;

}

.home{
  text-decoration: none;
  color: white; 
  text-transform: capitalize;

}

.navigation-list {
  list-style-type: none;
  margin: 0; 
  padding: 0; 
}

.navigation-list li {
  display: inline; 
  margin-right: 20px; 
}

.navigation-list li a {
  text-decoration: none; 
  color: white; 
  text-transform: capitalize;
  margin-right: 30px;
}

.navigation-list li a:hover {
  color: blueviolet;
}



</style>


<template>
  <!-- for this section maybe do a side menu for more links and pages. EX. top 100 moveis....sum from IMBd -->
    <v-app>
    <v-navigation-drawer color="black" v-model="drawer">
      <v-list-item prepend-icon="" href="#/" title="bob" @click="drawer = !drawer"></v-list-item>
      <v-list-item prepend-icon="" href="#/" title=""
        @click="drawer = !drawer"></v-list-item>
      <v-list-item prepend-icon="" href="#/" title="" @click="drawer = !drawer"></v-list-item>
    </v-navigation-drawer>
  <!-- finish -->

    <v-app-bar color="black">
      <v-app-bar-nav-icon icon="mdi-movie-roll"  @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-app-bar-title ><a class="home" href="#/home">FilmFinder</a></v-app-bar-title>
      <ul class="navigation-list">
        <li><a href="#/filmCatalog">Film Catalog</a></li>
        <li><a href="#/reviews">Reviews</a></li>
        <li><a href="#/comingSoon">ComingSoon</a></li>
      </ul>
      <v-btn href="#/home" icon="mdi-home-outline"></v-btn>
      <v-btn icon="mdi-magnify"></v-btn>
      <!-- <v-btn icon="mdi-dots-vertical"></v-btn> -->
    </v-app-bar>

    <v-main style="background-color: darkslategray;">
      <component :is="currentView"></component>
    </v-main>
    <v-footer app color="black">Copyright 2024 by FilmFinder.com</v-footer>
  </v-app>
</template>
