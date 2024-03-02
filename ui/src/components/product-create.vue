<template>
    <div>
        <h3>Create Product</h3>
        {{ p }}
        <div>
            <p>Name:</p>
            <input v-model="p.name">
        </div>
        <div>
            <p>Options:</p>
            <button @click="p.availableOptions.push('')">add</button>
            <div v-for="i of p.availableOptions.length">
                product {{ i }}
                <input v-model="p.availableOptions[i-1]">
                <button @click="p.availableOptions.splice(i-1, 1)">Remove</button>
            </div>
        </div>
        <button @click="createProduct(p)">Create</button>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Product } from '../domain'
import { useClient } from '../client';

const {createProduct} = useClient()

const p = ref<Product>({
    id: 0,
    name: "",
    availableOptions: []
})

</script>