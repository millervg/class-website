<script setup>
import { ref, computed } from 'vue';
import Home from './Home.vue'; 
import Gallery from './Gallery.vue'; 
import Etc from './Etc.vue';

const routes = {
    '/': Home,
    '/Gallery': Gallery,
    '/Etc': Etc,

}

const currentPath = ref(window.location.hash)

window.addEventListener("hashchange", () => {
    currentPath.value = window.location.hash
})

const currentView = computed(() => {
    return routes[currentPath.value.slice(1) || '/'] || NotFound
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
    <v-app>
    <v-navigation-drawer color="black" v-model="drawer">
      <v-list-item prepend-icon="mdi-home-circle" href="#/" title="Episodes" @click="drawer = !drawer"></v-list-item>
      <v-list-item prepend-icon="mdi-play-box-outline " href="#/Gallery" title="Gallery"
        @click="drawer = !drawer"></v-list-item>
      <v-list-item prepend-icon="mdi-chat" href="#/Etc" title="Etc" @click="drawer = !drawer"></v-list-item>
    </v-navigation-drawer>

    <v-app-bar color="black">
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-app-bar-title>FilmFinder</v-app-bar-title>
      <ul class="navigation-list">
        <li ><a href="#/">Episodes</a></li>
        <li><a href="#/Gallery">Gallery</a></li>
        <li><a href="#/Etc">Etc</a></li>
      </ul>
      <v-btn href="#/" icon="mdi-home-outline"></v-btn>
      <v-btn icon="mdi-magnify"></v-btn>
      <v-btn icon="mdi-dots-vertical"></v-btn>
    </v-app-bar>

    <v-main >
      <component :is="currentView"></component>
    </v-main>
    <v-footer app color="black">Copyright 2024 by FilmFinder.com</v-footer>
  </v-app>
</template>
