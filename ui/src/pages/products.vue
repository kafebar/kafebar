<template>
    <div class="max-w-screen-lg mx-auto p-4">
        <div class="flex justify-between">
            <h1 class="text-2xl mb-4">Products</h1>
            <button @click="creatingProduct = true">New Product</button>
        </div>
        <div class="grid grid-cols-3 gap-4">
            <ProductTile v-for="product of products" :product="product" @click="editedProduct = product" />
        </div>
    </div>
    <productModalEdit :show="editedProduct !== undefined" :product="editedProduct" @close="editedProduct = undefined"/>
    <productModalCreate :show="creatingProduct" @close="creatingProduct = false"/>
</template>

<script setup lang="ts">
import ProductTile from '../components/product-tile.vue';
import productModalEdit from '../components/product-modal-edit.vue';
import productModalCreate from '../components/product-modal-create.vue';
import { useClient } from '../client';
import { ref } from 'vue';
import { Product } from '../domain';

const {products} = useClient()

const editedProduct = ref<Product>()
const creatingProduct = ref(false)

</script>