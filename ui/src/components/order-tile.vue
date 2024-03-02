<template>
    <div class="p-4 bg-gray-700 rounded-md cursor-pointer" :class="order.isArchived ? 'opacity-50': ''">
        <div class="flex justify-between gap-2 items-start">
            <div>
                <h3 class="text-xl font-bold">#{{ order.id }}</h3>
                <h3 class="text-lg mb-2">{{ order.name }}</h3>
            </div>
            <button @click="updateArchiveStatus2" class="text-lg p-2 bg-gray-800 rounded-md">
                {{order.isArchived ? 'Unarchive': 'Archive'}}
            </button>
        </div>
        <div class="grid grid-cols-2 gap-2">
            <orderItemTile v-for="item of order.items" :active="!order.isArchived" :order-item="item" />
        </div>
    </div>
</template>

<script setup lang="ts">
import {PropType, defineProps} from 'vue'
import { Order } from '../domain';
import orderItemTile from './order-item-tile.vue';
import { useClient } from '../client';

const {updateOrderArchiveStatus} = useClient()

const props = defineProps({
    order: {type: Object as PropType<Order>, required: true}
})

async function updateArchiveStatus2() {
    await updateOrderArchiveStatus(props.order.id, !props.order.isArchived)
}

</script>