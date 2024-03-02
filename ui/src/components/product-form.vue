<template>
    <div>
        <div>
            <p>Name:</p>
            <input v-model="product.name">
        </div>
        <div>
            <p>Price:</p>
            <input type="number" v-model="product.price">
        </div>
        <div>
            <div class="flex justify-between items-center py-4">
                <p>Options:</p>
                <button @click="product.availableOptions.push('')" class="bg-gray-700 px-2 rounded-full">+</button>
            </div>
            <div v-for="i of product.availableOptions.length" class="flex gap-2 items-center">
                <input v-model="product.availableOptions[i-1]">
                <button @click="product.availableOptions.splice(i-1, 1)" class="bg-red-500 px-2 rounded-full">-</button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { PropType, defineModel } from 'vue';
import { Product } from '../domain'

const product = defineModel({
    type: Object as PropType<Product>,
    required: true
})

</script>