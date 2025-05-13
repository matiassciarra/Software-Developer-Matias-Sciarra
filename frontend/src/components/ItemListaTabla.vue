<script setup lang="ts"">
import { Item } from "../services/axios"
import { PropType } from "vue"; // Importa PropType de Vue

defineProps({
  item: {
    type: Object as PropType<Item>, // Especifica que "item" es del tipo "Item"
    required: true, // Marca la prop como obligatoria
  },
});
</script>
    


<template>
    <tr class="bg-white dark:bg-gray-800 hover:bg-gray-50 dark:hover:bg-gray-600">
        <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">{{ item.Ticker }}</th>
            <td class="px-6 py-4">{{ item.Brokerage }}</td>
            <td class="px-6 py-4">{{ item.Company }}</td>
            <td class="px-6 py-4">{{ item.RatingFrom }}</td>
            <td class="px-6 py-4">{{ item.RatingTo }}</td>
            <td class="px-6 py-4">${{ item.TargetFrom }}</td>
            <td class="px-6 py-4">${{ item.TargetTo }}</td>
            <td class="px-6 py-4">{{ item.Action }}</td>
    </tr>
</template>