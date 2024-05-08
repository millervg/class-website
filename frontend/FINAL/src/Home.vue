<script setup>
import { onMounted } from 'vue';

onMounted(() => {
  let openBtn = document.querySelector('#open');
  let closeBtn = document.querySelector('#close');
  let modal = document.querySelector('.modal');

  openBtn.addEventListener("click", () => {
    modal.style.display = "block";
  });

  closeBtn.addEventListener("click", () => {
    modal.style.display = "none";
  });
});
</script>

<template>
  <v-carousel show-arrows="hover" hide-delimiters class="carousel" >
    <v-carousel-item src="../images/Oppenheimer.png" cover>
      <div class="text">
      <h1>Oppenheimer<span class="inline-block">(2023)</span>
</h1>
      <p>
      <span class="inline-block" style="background-color: orange; border-style: none; border-radius: 15px; padding: 5px; color: black; font-weight: bold;" >HD</span>
      <span class="inline-block">6.6</span>
      <span class="inline-block">118 min</span>
      <span class="inline-block">Bioagraphy</span>
      <span class="inline-block">Drama</span>
      <span class="inline-block">History</span>
      </p>      
      <p>The story of J. Robert Oppenheimer’s role in the development of the atomic bomb during World War II.</p>
    </div>
    </v-carousel-item>
    <v-carousel-item src="../images/Avengers.png" cover>
      <div class="text">
      <h1>Avengers Endgame<span class="inline-block">(2019)</span>
</h1>
      <p>
      <span class="inline-block" style="background-color: orange; border-style: none; border-radius: 15px; padding: 5px; color: black; font-weight: bold;" >HD</span>
      <span class="inline-block">9.0</span>
      <span class="inline-block">182 min</span>
      <span class="inline-block">Sci-Fi</span>
      <span class="inline-block">Action</span>
      </p>
      <p>Adrift in space with diminishing oxygen, Tony Stark sends a desperate message to Pepper Potts, while the remaining Avengers, alongside Thor, Black Widow, Captain America, and Bruce Banner, plot to resurrect their fallen allies for an ultimate showdown against Thanos, the malevolent demigod who devastated both Earth and the universe.
</p>
    </div>
    </v-carousel-item>
    <v-carousel-item src="../images/batman.png" cover>
      <div class="text">
      <h1>Star Wars the Force Awakens<span class="inline-block">(2015)</span>
</h1>
      <p>
      <span class="inline-block" style="background-color: orange; border-style: none; border-radius: 15px; padding: 5px; color: black; font-weight: bold;" >HD</span>
      <span class="inline-block">9.3</span>
      <span class="inline-block">126 min</span>
      <span class="inline-block">Sci-Fi</span>
      <span class="inline-block">Action</span>
      </p>
      <p>Star Wars: The Force Awakens is a 2015 American epic space opera film co-produced, co-written, and directed by J. J. Abrams. The sequel to Return of the Jedi, it is the seventh film in the "Skywalker Saga". 
</p>
    </div>
    </v-carousel-item>
  </v-carousel>
  
        <div>
          <button id="open">Click ME</button>
          <div class="modal">
            <h5 class="modal-title">Welcome to FilmFinder.com</h5>
            <p><br>Where Every Movie is Just a Click Away!<br>Unleash the Magic of Cinema, Right at Your Fingertips!</p>
            <button id="close">x</button>
          </div>
        </div>

  <div class="grid">
    <aside>
      <h1 class="title">Movie Reviews</h1>
      <div class="card">
      <iframe  width="353" height="180" src="https://www.youtube.com/embed/UE2KQsOT6RY" allowfullscreen></iframe>
      <br>
      <iframe width="353" height="180" src="https://www.youtube.com/embed/dbsGMZh3pRQ" allowfullscreen></iframe>
      <br>
      <iframe width="353" height="180" src="https://www.youtube.com/embed/kyuNay3-X8k" allowfullscreen></iframe>
      </div>
    </aside>
    <aside>
      <h1 class="title">Upcoming Movies</h1>
      <div class="card">
        <iframe width="353" height="180" src="https://www.youtube.com/embed/73_1biulkYk" allowfullscreen></iframe>      
        <br>
        <iframe width="353" height="180" src="https://www.youtube.com/embed/xy8aJw1vYHo" allowfullscreen></iframe>
        <br>
        <iframe width="353" height="180" src="https://www.youtube.com/embed/YPY7J-flzE8" allowfullscreen></iframe>
      </div>
    </aside>
    <aside>
      <h1 class="title">Top 3 Lifetime Grosses</h1>
      <div class="card">
        <iframe width="353" height="180" src="https://www.youtube.com/embed/5PSNL1qE6VY" allowfullscreen></iframe>    
        <br>
        <iframe width="353" height="180" src="https://www.youtube.com/embed/TcMBFSGVi1c" allowfullscreen></iframe>
        <br>
        <iframe width="353" height="180" src="https://www.youtube.com/embed/d9MyW72ELq0" allowfullscreen></iframe>
      </div>
    </aside>
  </div>
</template>
  
  
  <style scoped>
  .modal{
    display: none;
    position: absolute;
    top: 48%;
    left: 50%;
    transform: translate(-50%, -50%);
    padding: 20px;
    text-align: center;
    border-radius: 10px;
    background-color: #000;
    color: #fff;
    animation-name: modalOpening;
    animation-duration: 1s;
    
  }

  @keyframes modalOpening {
    from {
      opacity: 0;
    }

    to {
      opacity: 100%;
    }
  }

  #open {
    margin-left: 50%;
    margin-top:30px;
    top: 50%;
    left: 50%;
    transform:  translate(-50%, -50%);
    padding: 15px, 50px;
    background-color: rgba(0, 0, 0, 0.157);
    border:  3px solid #000;
    font-weight: 800;
    font-size: 20px;
    text-transform: uppercase;
    letter-spacing: 2px;
    transition: 0.5s;
    cursor: pointer;
    border-radius: 5px;
  }

  #open:hover{
    transform: translate(-50%, -50%) scale(120%);
    background-color: #000;
    color: #fff;
  }

  #close {
    position: absolute;
    top: 15px;
    right: 15px;
    padding: 5px;
    background: transparent;
    border: 1px solid #fff;
    cursor: pointer;
    transition: 0.3s;
    color: #fff;
    border-radius: 4px;
  }

  #close:hover {
    background:#fff;
    color: #000;
    transform: scale(130%);
  }

  .modal-title{
    font-size: 25px;
    text-align: center;
    text-transform: uppercase;
    letter-spacing: 3px;
  }

  .modal p {
    margin-left: 185px;
    font-size:20px ;
  }

  .movieName{
    text-align: center;
  }

p {
  margin-top: 5px;
  max-width: 50%;
}

.inline-block {
  display: inline-block;
  margin-right: 10px; 
}

.text{
  color: white;
  text-align: left;
  padding: 15px;
  margin-top: 260px;
}

.section1 {
  text-align: center;
  margin-top: 10px;
  padding: 20px; 
}

.section1 h1:hover{
  box-shadow: 0 0 5px 2px rgba(0, 0, 0, 0.5);
}

.section1 h1{
    display: inline-block;
    border: 4.5px solid black;  
    padding: 10px;
    border-radius: 10px;
  }
.app {
    min-height: 100vh;
    justify-content: center;
    align-items: center;
  }
  
aside{
    margin-left: 15px;
  }

  iframe:hover {
  box-shadow: 0 0 5px 2px rgba(0, 0, 0, 0.5);
}

iframe{
    border: 4px solid black;
    border-radius: 20px;
  }

.title{
    font-weight: bolder;
    text-align: center;
  }
  
  .title2{
    font-weight: bolder;
    text-align: center;
  }
  
  .center {
    display: flex;
    justify-content: center;
  }

  .grid{
    display: flex;
    gap: 500px;
  }
</style>