<script setup>
import { ref } from 'vue';

const names = ref([
    'Bob',
    'Angie',
    'Monica'
]);

const cardTitle = ref("");
const cardGenre = ref("");
const cardYear = ref("");
const cardImg = ref("");
</script>

<template>
  <div class="container">
    <v-card class="card" variant="outlined">
      <v-card-title>{{ names[0] }}</v-card-title>
      <v-card-text>He is a builder</v-card-text>
    </v-card>
    <v-card class="card" variant="outlined">
      <v-card-title>{{ names[1] }}</v-card-title>
      <v-card-text>She loves to cook</v-card-text>
    </v-card>
    <v-card class="card" variant="outlined">
      <v-card-title>{{ names[2] }}</v-card-title>
      <v-card-text>She loves to knit</v-card-text>
    </v-card>
  </div>
</template>

<style scoped>
.container {
  margin-top: 200px;
  display: flex;
  justify-content: center;
}

.card:hover {
  box-shadow: 4px 8px black;
}
.card {
  text-align: center;
  margin-left: 10px;
  gap: 20px;
  width: 200px;
  height: 100px;
  background-color: white; 
}
</style>
