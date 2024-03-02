package kafebar

import (
	"context"
)

type Order struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	IsArchived bool        `json:"isArchived"`
	Items      []OrderItem `json:"items"`
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
	GetOrder(context.Context, int) (Order, error)
	GetOrders(context.Context) ([]Order, error)
	CreateOrder(context.Context, Order) (Order, error)
	UpdateOrderArchiveStatus(ctx context.Context, orderId int, isArchived bool) error

	GetOrderItem(context.Context, int) (OrderItem, error)
	AddOrderItem(context.Context, OrderItem) (OrderItem, error)
	// RemoveOrderItem(context.Context, int, int) error
	UpdateOrderItemStatus(context.Context, int, Status) error
}
