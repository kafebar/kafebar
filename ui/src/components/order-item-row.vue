<template>
    <div v-if="product" class="p-2 bg-gray-800 w-full rounded-md cursor-pointer flex mb-2 justify-between">
        <div>
            <h5 class="text-lg">{{ product.name }}</h5>
            <div v-if="orderItem.options" class="flex flex-wrap gap-2">
                <div v-for="option of orderItem.options" class="rounded-md p-1 bg-gray-700">
                    {{ option }}
                </div>
            </div>
        </div>
        <h3 class="text-xl font-bold">{{ product.price.toFixed(2) }}</h3>
    </div>

</template>

<script setup lang="ts">
import {PropType, computed, defineProps} from 'vue'
import { OrderItem } from '../domain';
import { useClient } from '../client';

const {products} = useClient()

const props = defineProps({
    orderItem: {type: Object as PropType<OrderItem>, required: true}
})

const product = computed(() => {
    return products.value.find(p => p.id == props.orderItem.productId)
})

</script>