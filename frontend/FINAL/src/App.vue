<script setup>
import { ref, computed } from 'vue';
import filmCatalog from './filmCatalog.vue';
import comingSoon from './comingSoon.vue';
import home from './Home.vue';

const routes = {
    '/filmCatalog': filmCatalog,
    '/comingSoon': comingSoon,
    '/home': home,

}

const socials = [
  'mdi-instagram',
  'mdi-linkedin',
  'mdi-github',
]

const currentPath = ref(window.location.hash)

window.addEventListener("hashchange", () => {
    currentPath.value = window.location.hash
})

const currentView = computed(() => {
    return routes[currentPath.value.slice(1) || '/home'] 
})
</script>

<template>
    <v-app>
    <v-app-bar color="black">
      <v-app-bar-nav-icon icon="mdi-movie-roll"></v-app-bar-nav-icon>
      <v-app-bar-title ><a class="home" href="#/home">FilmFinder</a></v-app-bar-title>
      <ul class="navigation-list">
        <li><a href="#/home">Home</a></li>
        <li><a href="#/filmCatalog">Film Catalog</a></li>
        <li><a href="#/comingSoon">Watchlist</a></li>
      </ul>
    </v-app-bar>

    <v-main style="background-color: darkslategray;">
      <component :is="currentView"></component>
    </v-main>

    <v-footer color="black">
  Copyright 2024 by FilmFinder.com
  <div class="footer">
    <a href="https://www.instagram.com/vickigracee/" class="social-link">
      <v-icon>{{ socials[0] }}</v-icon>
    </a>
    <a href="https://www.linkedin.com/in/abuchiosuala/" class="social-link">
      <v-icon>{{ socials[1] }}</v-icon>
    </a>
    <a href="https://github.com/millervg" class="social-link">
      <v-icon>{{ socials[2] }}</v-icon>
    </a>
  </div>
</v-footer>
  </v-app>
</template>

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

.btn{
  margin-right: 10px;
  width: auto;
}

.social-link{
  padding: 5px;
  color: white;
}

.social-link:hover{
  color:#8a2be2;
}

.footer{
  display: flex;
  margin-left: auto;
  align-items: center;
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
