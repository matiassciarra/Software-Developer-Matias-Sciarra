
<script setup>
    import ItemListaTabla from "../components/ItemListaTabla.vue"
    import { ref, onMounted } from 'vue';
    import { fetchItems } from '../services/axios'; // Importa la función desde donde la guardaste

const data = ref(null); // Variable reactiva para almacenar los datos
const error = ref(null); // Variable reactiva para manejar errores
const isLoading = ref(false); // Variable reactiva para el estado de carga

async function loadData() {
  isLoading.value = true;
  try {
    data.value = await fetchItems(); // Llama a la función externa
  } catch (err) {
    error.value = err.message; // Maneja el error
  } finally {
    isLoading.value = false; // Finaliza el estado de carga
  }
}

onMounted(() => {
  loadData(); // Ejecuta la función al montar el componente
});
</script>


<template>
  <div class="flex h-screen">
    <!-- Nav-Bar -->
    <aside id="default-sidebar" class="w-64 h-full overflow-y-auto bg-gray-50 dark:bg-gray-800">
      <div class="h-full px-3 py-4">
        <ul class="space-y-2 font-medium">
          <li>
            <a href="#" class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group">
              <svg class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 22 21">
                <path d="M16.975 11H10V4.025a1 1 0 0 0-1.066-.998 8.5 8.5 0 1 0 9.039 9.039.999.999 0 0 0-1-1.066h.002Z" />
                <path d="M12.5 0c-.157 0-.311.01-.565.027A1 1 0 0 0 11 1.02V10h8.975a1 1 0 0 0 1-.935c.013-.188.028-.374.028-.565A8.51 8.51 0 0 0 12.5 0Z" />
              </svg>
              <span class="ml-3">Filtros</span>
            </a>
          </li>
        </ul>
      </div>
    </aside>

    <!-- Tabla -->
    <div class="flex-1">
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
            <ItemListaTabla v-for="item in data" :item="item"/>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

