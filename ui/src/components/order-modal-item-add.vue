<template>
    <Modal :show="orderItem !== undefined" @close="$emit('close')" name="Select Options">
        <template v-if="Boolean(orderItem)">

            <orderItemForm v-model="orderItem"/>

            <div class="flex justify-center gap-4 py-2"> 
                <button @click="$emit('add')" class="px-4 py-2 rounded-full bg-green-500">Save</button>
            </div>

        </template>
   </Modal>
</template>

<script setup lang="ts">
import { PropType } from 'vue';
import Modal from './modal.vue'
import { OrderItem } from '../domain';
import orderItemForm from './order-item-form.vue';

const orderItem = defineModel({
    type: Object as PropType<OrderItem>,
})


const emit = defineEmits(['close', 'add'])

</script>