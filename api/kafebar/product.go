package kafebar

import "context"

type Product struct {
	Id               int      `json:"id"`
	Name             string   `json:"name"`
	Price            float64  `json:"price"`
	AvailableOptions []string `json:"availableOptions"`
}

type ProductService interface {
	CreateProduct(context.Context, Product) (Product, error)
	UpdateProduct(context.Context, Product) (Product, error)
	DeleteProduct(context.Context, int) error
	GetProducts(context.Context) ([]Product, error)
}
