package kafebar

import (
	"context"
)

type Order struct {
	Id     int         `json:"id"`
	Name   string      `json:"name"`
	Status Status      `json:"status"`
	Items  []OrderItem `json:"items"`
}

type Status string

const (
	StatusTodo       Status = "Todo"
	StatusInProgress Status = "InProgress"
	StatusDone       Status = "Done"
	StatusArchived   Status = "Archived"
)

type OrderItem struct {
	Id        int      `json:"id"`
	OrderId   int      `json:"orderId"`
	ProductId int      `json:"productId"`
	Status    Status   `json:"status"`
	Options   []string `json:"options"`
}

type OrderService interface {
	CreateOrder(context.Context, Order) (Order, error)
	// EditOrder(context.Context, Order) error

	AddOrderItem(context.Context, OrderItem) (OrderItem, error)
	// RemoveOrderItem(context.Context, int, int) error
	// EditOrderItem(context.Context, int, OrderItem) error

	GetOrders(context.Context) ([]Order, error)
}
