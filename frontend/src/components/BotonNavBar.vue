
<script setup>
// Declarar las props que el componente hijo espera recibir
defineProps({
  message: String, // Una prop llamada "message" de tipo String
  ruta: String
});
</script>

<template>
        <li>
          <router-link :to="ruta">
            <a href="#" 
            class="block py-2 px-3 text-white bg-blue-700 rounded-sm md:bg-transparent md:p-0 dark:text-white" 
            >{{ message }}</a>
          </router-link>
        </li>
</template>


