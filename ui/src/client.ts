import { ref } from "vue";
import { EventData, Order, Product } from "./domain";

const products = ref<Product[]>([])
const orders = ref<Order[]>([])

let startedSyncing = false

export function useClient() {
    if(!startedSyncing) {
        startedSyncing = true
        startSync()
    }

    return {
        products, 
        orders,
        createProduct,
        updateProduct,
        deleteProduct,
        createOrder
    }
}


export async function startSync() {
    const eventSource = new EventSource(`${apiUrl}/events`);
    console.log(eventSource)

    let loadedInitialData = false
    let bufferedEvents: EventData[] = []

    eventSource.onopen = async function() {
        const [fetchedProducts, fetchedOrders] = await Promise.all([ getProducts(),getOrders() ])
        products.value = fetchedProducts
        orders.value = fetchedOrders
        loadedInitialData = true
        for(const bufferedEvent of bufferedEvents) {
            handleEvent(bufferedEvent)
        }
    }

    eventSource.onmessage = function(event) {
        console.log("received event", event)

        const evData = JSON.parse(event.data) as EventData

        if(!loadedInitialData) {
            bufferedEvents.push(evData)
        } else {
            handleEvent(evData)
        }
    };
}

async function handleEvent(ev: EventData) {
    switch(ev.type) {
        case "ProductCreated":
        case "ProductUpdated": {
            console.log("product created/updated", ev.data)
            const idx = products.value.findIndex(p => p.id == ev.data.id)
            if (idx == -1) {
                products.value.push(ev.data)
            } else {
                products.value[idx] = ev.data
            }
            break;
        }
            
        case "ProductDeleted": {
            console.log("product deleted", ev.data)
            const idx = products.value.findIndex(p => p.id == ev.data)
            if (idx == -1) {
                console.log("product not found")
            }
            products.value.splice(idx, 1)
            break
        }

        case "OrderCreated":
        case "OrderUpdated": {
            console.log("order created/updated", ev.data)
            const idx = orders.value.findIndex(p => p.id == ev.data.id)
            if (idx == -1) {
                orders.value.push(ev.data)
            } else {
                orders.value[idx] = ev.data
            }
            break;
        }
            
        case "OrderDeleted": {
            console.log("order deleted", ev.data)
            const idx = orders.value.findIndex(p => p.id == ev.data.id)
            if (idx == -1) {
                console.log("order not found")
            }
            orders.value.splice(idx, 1)
            break
        }
    }
}


const apiUrl = import.meta.env.VITE_API_BASE_URI

async function createProduct(p: Product):Promise<Product> {
    const res = await fetch(`${apiUrl}/products`, {
        method: "POST",
        body: JSON.stringify(p),
    })

    if (!res.ok) {
        throw Error("cannot create product")
    }
    const createdProduct = await res.json()

    return createdProduct
}

async function updateProduct(p: Product):Promise<Product> {
    const res = await fetch(`${apiUrl}/products`, {
        method: "PUT",
        body: JSON.stringify(p),
    })

    if (!res.ok) {
        throw Error("cannot update product")
    }
    const createdProduct = await res.json()

    return createdProduct
}


async function deleteProduct(id: number):Promise<Product> {
    const res = await fetch(`${apiUrl}/products/${id}`, {
        method: "DELETE"
    })

    if (!res.ok) {
        throw Error("cannot delete product")
    }
    const createdProduct = await res.json()

    return createdProduct
}


async function getProducts():Promise<Product[]> {
    const res = await fetch(`${apiUrl}/products`)

    if (!res.ok) {
        throw Error("cannot get products")
    }
    const products = await res.json()
    return products
}

async function createOrder(o: Order):Promise<Order> {
    const res = await fetch(`${apiUrl}/orders`, {
        method: "POST",
        body: JSON.stringify(o),
    })

    if (!res.ok) {
        throw Error("cannot order product")
    }
    const createdOrder = await res.json()
    return createdOrder
}

async function getOrders():Promise<Order[]> {
    const res = await fetch(`${apiUrl}/orders`)

    if (!res.ok) {
        throw Error("cannot get orders")
    }
    const products = await res.json()
    return products
}
