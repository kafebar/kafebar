<template>
    <div class="p-4 bg-gray-500 rounded-md cursor-pointer">
        <div class="flex mb-2 justify-between">
            <h3 class="text-xl mb-2">{{ product.name }}</h3>
            <h3 class="text-xl mb-2 font-bold">{{ product.price?.toFixed(2) }}</h3>
        </div>
        <div class="flex flex-wrap gap-1">
            <div v-for="option of product.availableOptions" class="rounded-md px-1 bg-gray-700">
                {{ option }}
            </div>
        </div>
    </div>

</template>

<script setup lang="ts">
import {PropType, defineProps} from 'vue'
import { Product } from '../domain';

defineProps({
    product: {type: Object as PropType<Product>, required: true}
})

</script>