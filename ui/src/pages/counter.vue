<template>
    <div>
        <h1 class="text-2xl mb-4">New Order</h1>
        <div class="mb-4">
            <h3 class="text-xl">Name:</h3>
            <input v-model="order.name">
        </div>

        <div v-if="order.items.length">
            <h3 class="text-xl">Order:</h3>
            <orderItemTable v-model="order.items" class="mb-4"/>
            <div class="flex justify-center gap-4 py-2"> 
                <button @click="order = emptyOrder()" class="px-4 py-2 rounded-full bg-red-500">Discard</button>
                <button @click="save" class="px-4 py-2 rounded-full bg-green-500">Finalize Order</button>
            </div>
        </div>

        <h3 class="text-xl mb-4">Add a product:</h3>
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
            <productTile v-for="product of products" :product="product" @click="startAddingOrderItem(product)" />
        </div>

        <orderModalAddItem 
            :show="Boolean(addingOrderItem)" 
            v-model="addingOrderItem" 
            @close="addingOrderItem = undefined" 
            @add="order.items.push(addingOrderItem!); addingOrderItem = undefined"/>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Order, OrderItem, Product } from '../domain';
import { useClient } from '../client';
import productTile from '../components/product-tile.vue';
import orderModalAddItem from '../components/order-modal-item-add.vue';
import orderItemTable from '../components/order-item-table.vue';

const {products, createOrder} = useClient()

const order = ref<Order>(emptyOrder())
const addingOrderItem = ref<OrderItem>()

function startAddingOrderItem(product: Product) {
    const newOrder = emptyOrderItem()
    newOrder.productId = product.id

    if(product.availableOptions.length == 0) {
        order.value.items.push(newOrder)
    } else {
        addingOrderItem.value = newOrder
    }
}

function emptyOrder(): Order {
    return {
        id: 0,
        name: "",
        status: 'Todo',
        items: []
    }
}

async function save() {
    await createOrder(order.value);
    order.value = emptyOrder()
}

function emptyOrderItem(): OrderItem {
    return {
        id: 0,
        orderId: 0,
        productId: 0,
        status: 'Todo',
        options: []
    }
}
</script>