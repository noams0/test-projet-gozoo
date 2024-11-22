<script setup>
import {ref, onMounted, reactive, onUnmounted} from "vue";
import api from "../service/api";

// Déclaration des données avec `ref`
const animals = ref([]);
const visitors = ref([]);
let intervalId = null;

// Appel de l'API lors du montage du composant
const fetchSimulationState = async () => {
  try {
    const response = await api.get("/get_animals");
    animals.value = response.data;
  } catch (error) {
    console.error("Erreur lors de la récupération des animaux :", error);
  }
  try {
    const response = await api.get("/get_visitors");
    visitors.value = response.data;
  } catch (error) {
    console.error("Erreur lors de la récupération des visitors :", error);
  }

  // Placer les animaux en haut de la grille
  animals.value.forEach((animal, index) => {
    console.log(animal.escaped);
    cells[index].type = "animal";
    cells[index].id = animal.id;
    cells[index].typeID = `animal-${animal.id}`; // Génère un identifiant combiné
  });

  console.log(cells);
  console.log(visitors.value);
  console.log(animals.value);
// Placer les visiteurs en bas de la grille
  visitors.value.forEach((visitor, index) => {
    cells[cells.length - 1 - index].type = "visitor";
    cells[cells.length - 1 - index].id = visitor.id;
    cells[cells.length - 1 - index].typeID = `visitor-${visitor.id}`;
  });
};

onMounted(() => {
  fetchSimulationState(); // Récupère l'état initial
  intervalId = setInterval(fetchSimulationState, 500); // Met à jour toutes les 3 secondes
});

onUnmounted(() => {
  clearInterval(intervalId); // Nettoie l'intervalle lorsque le composant est démonté
});



const gridSize = 10; // 10x10 grid
const cells = reactive(
    Array.from({ length: gridSize * gridSize }, () => ({
      type: null,
      id: null,
      typeID: null,
    }))
);


const handleCellClick = async (index) => {
  const cell = cells[index];
  if (cell.type === "animal") {
    try {
      // Envoyer la mise à jour au back-end
      await api.post("/update_animal_escape", { id: cell.id, escaped: true });

      // Mise à jour locale
      const animal = animals.value.find((a) => a.id === cell.id);
      if (animal) animal.escaped = true;
      console.log(`Animal ${cell.id} marqué comme échappé`);
    } catch (error) {
      console.error("Erreur lors de la mise à jour de l'animal :", error);
    }
  }
};

const isEscaped = (id) => {
  const animal = animals.value.find((a) => a.id === id);
  return animal ? animal.escaped : false;
};

const getVisitorStateClass = (id) => {
  const visitor = visitors.value.find((v) => v.id === id);
  if (!visitor) return '';
  return visitor.is_panicked ? 'panicked' : 'calm';
};


</script>

<template>
  <div>
    <h2>Liste des animaux</h2>
    <ul v-if="animals.length">
      <li v-for="animal in animals" :key="animal.id">
        {{ animal.name }} - Comportement : {{ animal.behavior }} - Échappé : {{ animal.escaped ? "Oui" : "Non" }}
      </li>
    </ul>
    <p v-else>Chargement des données...</p>
  </div>

  <div>
    <h2>Liste des visiteurs</h2>
    <ul v-if="visitors.length">
      <li v-for="visitor in visitors" :key="visitor.id">
        {{ visitor.id }} - En panique : {{ visitor.is_panicked ? "Oui" : "Non" }}
      </li>
    </ul>
    <p v-else>Chargement des données...</p>
  </div>
  <div class="grid">
    <div
        v-for="(cell, index) in cells"
        :key="index"
        :class="[
    'cell',
    cell.type === 'animal'
      ? (isEscaped(cell.id) ? 'escaped' : 'not-escaped')
      : cell.type === 'visitor'
      ? getVisitorStateClass(cell.id)
      : 'empty',
  ]"
        @click="handleCellClick(index)"
    >
      <span v-if="cell.type === 'animal'">🐾</span>
      <span v-else-if="cell.type === 'visitor'">🧍</span>
      <span v-else></span>
    </div>

  </div>
</template>

<style>
.grid {
  display: grid;
  grid-template-columns: repeat(10, 30px); /* 10 cells wide */
  grid-template-rows: repeat(10, 30px); /* 10 cells tall */
  gap: 2px;
  border: 1px solid #000;
}

.cell {
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid #ddd;
  background-color: #f0f0f0;
}


.cell.empty {
  background-color: #f0f0f0;
}

.cell.not-escaped {
  background-color: red;
}

.cell.escaped {
  background-color: green;
}

.cell.panicked {
  background-color: #ff0000;
}

.cell.calm {
  background-color: lightblue;
}


</style>