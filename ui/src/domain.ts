
export type EventData = {
    type: string
    data: any
}

export type Product = {
    id: number
    name: string
    price: number
    availableOptions: string[]
}

export type Order = {
    id: number
    name: string
    isArchived: boolean
    items: OrderItem[]
}

export type OrderItem = {
    id: number
    orderId: number
    productId: number
    status: Status
    options: string[]
}

export type Status = "Todo" | "InProgress" | "Done" 