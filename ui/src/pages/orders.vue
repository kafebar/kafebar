<template>
    <div>
        <h1 class="text-2xl mb-4">Orders</h1>

        <div class="grid grid-cols-2 mb-4 justify-around text-lg">
            <button 
                @click="showingArchived = false" 
                class="p-2 rounded-md"
                :class="showingArchived ? '': 'bg-gray-800'">
                Current
            </button>
            <button 
                @click="showingArchived = true" 
                class="p-2 rounded-md"
                :class="showingArchived ? 'bg-gray-800': ''">
                Archived
            </button>
        </div>
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
            <orderTile v-for="order of shownOrders" :order="order" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import { useClient } from '../client';
import orderTile from '../components/order-tile.vue';

const showingArchived = ref(false)

const {orders} = useClient()

const shownOrders = computed(() => {
    return orders.value.filter(o => o.isArchived == showingArchived.value)
})
</script>