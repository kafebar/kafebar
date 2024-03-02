<template>
    <Modal :show="productEdit !== undefined" @close="$emit('close')" name="Edit product">
        <template v-if="productEdit !== undefined">
            <productForm v-model="productEdit"/>

            <div class="flex justify-center gap-4 py-2"> 
                <button @click="deleteProduct(productEdit.id); $emit('close')" class="px-4 py-2 rounded-full bg-red-500">Delete</button>
                <button @click="updateProduct(productEdit); $emit('close')" class="px-4 py-2 rounded-full bg-green-500">Save</button>
            </div>
        </template>
    </Modal>
</template>

<script setup lang="ts">
import { PropType, ref, watchEffect } from 'vue';
import Modal from './modal.vue'
import productForm from './product-form.vue';
import { Product } from '../domain';
import { useClient } from '../client';

const {deleteProduct, updateProduct} = useClient()

const props = defineProps({
    product: {type: Object as PropType<Product>}
})

const productEdit = ref<Product>()

watchEffect(() => {
    if(props.product) {
        productEdit.value = JSON.parse(JSON.stringify(props.product))
    }
})

defineEmits(['close'])

</script>