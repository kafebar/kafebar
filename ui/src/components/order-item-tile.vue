<template>
    <div class="p-4 rounded-md cursor-pointer" :class="bgColorClass" @click="updateStatus">
        <div class="flex mb-2 justify-between">
            <h3 class="text-xl mb-2">{{ product?.name }}</h3>
        </div>
        <div class="flex flex-wrap gap-1">
            <div v-for="option of orderItem.options" class="rounded-md px-1 bg-gray-700">
                {{ option }}
            </div>
        </div>
    </div>

</template>

<script setup lang="ts">
import {PropType, computed, defineProps} from 'vue'
import { OrderItem, Status } from '../domain';
import { useClient } from '../client';

const {products, updateOrderItemStatus} = useClient()


const props = defineProps({
    orderItem: {type: Object as PropType<OrderItem>, required: true},
    active: {type: Boolean, default: true}
})

const product = computed(() => {
    return products.value.find(p => p.id == props.orderItem.productId)
})

const bgColorClass = computed(() => {
    if(!props.active) {
        return 'bg-gray-500'
    }
    switch (props.orderItem.status) {
        case 'Todo':
            return 'bg-gray-500'
        case 'InProgress':
            return 'bg-yellow-500'
        case 'Done':
            return 'bg-green-500'
    }
})



async function updateStatus() {
    if(!props.active) {
        return
    }
    await updateOrderItemStatus(props.orderItem.id, getNextStatus(props.orderItem.status))
}

function getNextStatus(status: Status): Status {
    switch(status) {
        case 'Todo':
            return 'InProgress'
        case 'InProgress':
            return "Done"
        case "Done":
            return "Todo"
    }
}

</script>