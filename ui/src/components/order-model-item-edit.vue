<template>
    <Modal :show="orderItem !== undefined" @close="$emit('close')" name="Select Options">
        <template v-if="Boolean(itemEdit)">
            <orderItemForm v-model="itemEdit"/>

            <div class="flex justify-center gap-4 py-2"> 
                <button @click="$emit('remove')" class="px-4 py-2 rounded-full bg-red-500">Remove</button>
                <button @click="$emit('update', itemEdit)" class="px-4 py-2 rounded-full bg-green-500">Save</button>
            </div>
        </template>
   </Modal>
</template>

<script setup lang="ts">
import { PropType, ref, watchEffect } from 'vue';
import Modal from './modal.vue'
import { OrderItem } from '../domain';
import orderItemForm from './order-item-form.vue';

const props = defineProps({
    orderItem: {type: Object as PropType<OrderItem>}
})

const itemEdit = ref<OrderItem>()

watchEffect(() => {
    if(props.orderItem) {
        itemEdit.value = JSON.parse(JSON.stringify(props.orderItem))
    }
})

const emit = defineEmits(['close', 'remove', 'update'])

</script>