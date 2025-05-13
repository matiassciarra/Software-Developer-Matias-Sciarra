<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { fetchItems, fetchBestItems } from '../services/axios'; // Importa la función para obtener los datos
import { filterByFieldContains } from '../services/filters'; // Importa la función de filtrado

import ItemListaTabla from "../components/ItemListaTabla.vue";
import ItemListaSideBar from "../components/ItemListaSideBar.vue";
import { fetchRatings } from '../services/axios';

// Datos originales y filtrados
const data = ref([]); // Todos los datos
const filteredData = ref([]); // Datos filtrados

// Control de errores y carga
const error = ref(null);
const isLoading = ref(false);

// Control de paginación
const currentPage = ref(1);
const itemsPerPage = 10;

// Filtros
const filters = ref({
  Ticker: '',
  Brokerage: '',
  Company: '',
  RatingFrom: '',
  RatingTo: ''
});
// Variables reactivas para almacenar las opciones de RatingFrom y RatingTo
const ratingFromOptions = ref<string[]>([]);
const ratingToOptions = ref<string[]>([]);
const isLoadingRatings = ref(false);
const ratingError = ref<string | null>(null);

// Función para cargar las opciones de ratings
async function loadRatings() {
  isLoadingRatings.value = true;
  try {
    const ratings = await fetchRatings(); // Llama a la función para obtener los ratings
    ratingFromOptions.value = ratings; // Asigna las opciones a RatingFrom
    ratingToOptions.value = ratings; // Asigna las opciones a RatingTo
  } catch (error) {
    ratingError.value = 'Error al cargar las opciones de rating';
    console.error(error);
  } finally {
    isLoadingRatings.value = false;
  }
}


// Función para cargar los datos
async function loadData() {
  isLoading.value = true;
  try {
    const items = await fetchItems(); // Llama a la función para obtener los datos
    data.value = items;
    filteredData.value = items; // Inicialmente, los datos filtrados son los mismos que los originales
  } catch (err) {
    error.value = err.message;
  } finally {
    isLoading.value = false;
  }
}

async function getBestItems() {
  isLoading.value = true;
  try {
    const result = await fetchBestItems(); // Llama a la función para obtener los datos

    // Transformamos los datos del JSON para extraer solo los "item"
    const transformedItems = result.recommendations.map((entry: { item: any; score: number }) => {
      return {
        ...entry.item, // Incluimos todas las propiedades del "item"
        score: entry.score, // Añadimos la propiedad "difference" para la tabla si es necesario
      };
    });

    data.value = transformedItems; // Asignamos los datos transformados
    filteredData.value = transformedItems; // Inicialmente, los datos filtrados son los mismos que los originales
  } catch (err) {
    error.value = err.message;
  } finally {
    isLoading.value = false;
  }
}

// Función para aplicar los filtros
function applyFilters() {
  let result = data.value;

  // Filtra por cada campo si hay un valor en el filtro
  result = filterByFieldContains(result, 'Ticker', filters.value.Ticker);
  result = filterByFieldContains(result, 'Brokerage', filters.value.Brokerage);
  result = filterByFieldContains(result, 'Company', filters.value.Company);
  result = filterByFieldContains(result, 'RatingFrom', filters.value.RatingFrom);
  result = filterByFieldContains(result, 'RatingTo', filters.value.RatingTo);

  // Actualiza los datos filtrados y reinicia la paginación
  filteredData.value = result;
  currentPage.value = 1;
  filters.value.RatingFrom = ''
  filters.value.RatingTo = ''
}

// Calcular los elementos a mostrar en la página actual
const paginatedData = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage;
  const end = start + itemsPerPage;
  return filteredData.value.slice(start, end);
});

// Función para ir a la página siguiente
function nextPage() {
  if (currentPage.value * itemsPerPage < filteredData.value.length) {
    currentPage.value++;
  }
}

// Función para ir a la página anterior
function prevPage() {
  if (currentPage.value > 1) {
    currentPage.value--;
  }
}

// Cargar los datos al montar el componente
onMounted(() => {
  loadData();
  loadRatings();
});
</script>

<template>
  <div class="flex h-screen">
    <!-- Nav-Bar -->
    <aside id="default-sidebar" class="w-64 h-full overflow-y-auto bg-gray-50 dark:bg-gray-800">
      <div class="h-full px-3 py-4">
        <ul class="space-y-2 font-medium">
          <ItemListaSideBar name="Filtros"></ItemListaSideBar>

          <ItemListaSideBar name="Ticker"></ItemListaSideBar>
          <input
            v-model="filters.Ticker"
            type="text"
            id="ticker"
            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          />

          <ItemListaSideBar name="Brokerage"></ItemListaSideBar>
          <input
            v-model="filters.Brokerage"
            type="text"
            id="brokerage"
            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          />

          <ItemListaSideBar name="Company"></ItemListaSideBar>
          <input
            v-model="filters.Company"
            type="text"
            id="company"
            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          />
          <ItemListaSideBar name="Rating from"></ItemListaSideBar>
          <select
      id="ratingFrom"
      class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
      v-model="filters.RatingFrom"
    >
      <option v-for="(rating, index) in ratingFromOptions" :key="index" :value="rating">
        {{ rating }}
      </option>
    </select>
          
          <ItemListaSideBar name="Rating to"></ItemListaSideBar>
         <select
      id="ratingTo"
      v-model="filters.RatingTo"
      class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
    >
      <option v-for="(rating, index) in ratingToOptions" :key="index" :value="rating">
        {{ rating }}
      </option>
    </select> 
        <ItemListaSideBar name="Best Items" @click="getBestItems"></ItemListaSideBar>
          
        </ul>
  
  <button
            @click="applyFilters"
            class="text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-blue-900 dark:hover:bg-blue-700 dark:focus:ring-blue-700 dark:border-blue-700"
          >
            Filtrar
          </button>

      </div>
    </aside>

    <!-- Tabla -->
    <div class="flex-1 overflow-x-auto">
      <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
        <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
          <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
              <th scope="col" class="px-6 py-3">Ticker</th>
              <th scope="col" class="px-6 py-3">Brokerage</th>
              <th scope="col" class="px-6 py-3">Company</th>
              <th scope="col" class="px-6 py-3">Rating From</th>
              <th scope="col" class="px-6 py-3">Rating To</th>
              <th scope="col" class="px-6 py-3">Target From</th>
              <th scope="col" class="px-6 py-3">Target To</th>
              <th scope="col" class="px-6 py-3">Action</th>
            </tr>
          </thead>
          <tbody>
            <ItemListaTabla v-for="(item, index) in paginatedData" :key="index" :item="item" />
          </tbody>
        </table>
      </div>
      <div class="flex items-center mt-4">
        <button
          @click="prevPage"
          class="text-white bg-gray-800 hover:bg-gray-900 focus:outline-none focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-gray-800 dark:hover:bg-gray-700 dark:focus:ring-gray-700 dark:border-gray-700"
          :disabled="currentPage === 1"
        >
          Atrás
        </button>
        <button
          @click="nextPage"
          class="text-white bg-gray-800 hover:bg-gray-900 focus:outline-none focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-gray-800 dark:hover:bg-gray-700 dark:focus:ring-gray-700 dark:border-gray-700"
          :disabled="currentPage * itemsPerPage >= filteredData.length"
        >
          Siguiente
        </button>
      </div>
    </div>
  </div>
</template>

