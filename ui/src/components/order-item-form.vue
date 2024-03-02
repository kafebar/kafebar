<template>
    <div>
        <label v-if="product" v-for="option of product.availableOptions" class="flex w-full py-2 gap-2 text-lg" :for="option">
            <input :id="option" type="checkbox" v-model="orderItem.options" :value="option" >
            <p>{{ option }}</p>
        </label>
    </div>
</template>

<script setup lang="ts">
import { PropType, computed, defineModel } from 'vue';
import { OrderItem } from '../domain'
import {useClient} from '../client'

const {products} = useClient()

const orderItem = defineModel({
    type: Object as PropType<OrderItem>,
    required: true
})

const product = computed(() => {
    return products.value.find(p => p.id == orderItem.value.productId)
})

</script>