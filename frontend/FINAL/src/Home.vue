<script setup>
import { ref } from 'vue';
import movieFetch from './components/movieFetch.vue';

</script>


<template>
    <section class="section1">
            <h1>Welcome to FilmFinder.com</h1>
    </section>
    <div class="app">
        <section>
            <div class="center">
                <v-card class="card">
                    <v-card-title>Hello</v-card-title>
                </v-card>
            </div>
        </section>
    </div>
</template>
  
  
<style scoped>
.section1 {
  text-align: center;
  margin-top: 10px;
  padding: 20px; 
}

.section1 h1:hover{
    box-shadow: 2px 4px black;

}

  .section1 h1{
    display: inline-block;
    border: 5px solid black;  
    padding: 10px;
    border-radius: 10px;
  }
  .app {
    min-height: 100vh;
    justify-content: center;
    align-items: center;
  }
  
  .card{
    max-width: 100px;
    background-color: black;
    color: white;
  }
  
  .center {
    display: flex;
    justify-content: center;
  }
</style>