<template>
    <div>
        <orderItemRow 
            v-for="(item, i) of items"
            @click="editedItemIdx = i"
            :order-item="item"/>

        <p class="text-right text-xl font-bold mr-2">
            Total: {{ total.toFixed(2) }}
        </p>

        <orderModelItemEdit 
            :show="editedItemIdx != -1" 
            :order-item="items[editedItemIdx]" 

            @update="items[editedItemIdx] = $event; editedItemIdx = -1"
            @remove="items.splice(editedItemIdx, 1); editedItemIdx = -1"
            @close="editedItemIdx = -1"
        />
    </div>
</template>

<script lang="ts" setup>
import { PropType, computed, ref } from 'vue';
import { OrderItem } from '../domain';
import orderItemRow from './order-item-row.vue';
import { useClient } from '../client';
import orderModelItemEdit from './order-model-item-edit.vue';

const {products} = useClient()

const editedItemIdx = ref(-1)

const items = defineModel({
    type: Array as PropType<OrderItem[]>,
     default: []
})

const total = computed(() => {
    let total = 0

    for (const item of items.value) {
        const product = products.value.find(p => p.id == item.productId)
        total += product?.price ?? 0
    }
    return total
})
</script>