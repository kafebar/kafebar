package kafebar

import (
	"context"
)

type Order struct {
	Id     string      `json:"id"`
	Name   string      `json:"name"`
	Status Status      `json:"status"`
	Items  []OrderItem `json:"items" firestore:"-"`
}

type Status string

const (
	StatusTodo       Status = "Todo"
	StatusInProgress Status = "InProgress"
	StatusDone       Status = "Done"
	StatusArchived   Status = "Archived"
)

type OrderItem struct {
	Id      string   `json:"id"`
	OrderId string   `json:"orderId"`
	Product Product  `json:"productId"`
	Status  Status   `json:"status"`
	Options []string `json:"options"`
}

type Product struct {
	Id               string   `json:"id"`
	Name             string   `json:"name"`
	AvailableOptions []string `json:"availableOptions"`
}

type ProductService interface {
	GetProducts() ([]Product, error)
}

type OrderService interface {
	CreateOrder(context.Context, Order) error
}
