package postgres

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

const (
	tableOrders           = "orders"
	tableOrderItems       = "order_items"
	tableOrderItemOptions = "order_item_options"
	tableProducts         = "products"
	tableProductOptions   = "product_options"

	columnId         = "id"
	columnName       = "name"
	columnStatus     = "status"
	columnIsArchived = "is_archived"

	columnOrderId   = "order_id"
	columnProductId = "product_id"
	columnPrice     = "price"

	columnOrderItemId = "order_item_id"
	columnOption      = "option"
)
