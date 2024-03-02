<template>
    <Modal :show="product !== undefined" @close="$emit('close')" name="New Product">
        <productForm v-model="product"/>

        <div class="flex justify-center gap-4 py-2"> 

            <button @click="save" class="px-4 py-2 rounded-full bg-green-500">Save</button>
        </div>

   </Modal>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import Modal from './modal.vue'
import productForm from './product-form.vue';
import { Product } from '../domain';
import { useClient } from '../client';

const {createProduct} = useClient()

const product = ref<Product>(emptyProduct())


function emptyProduct(): Product {
    return {
        id: 0,
        name: "",
        price: 0,
        availableOptions: []
    }
}

function save() {
    createProduct(product.value); 
    product.value = emptyProduct();
    emit('close')
}

const emit = defineEmits(['close'])

</script>