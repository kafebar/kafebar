package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"slices"

	sq "github.com/Masterminds/squirrel"
	"github.com/kafebar/kafebar/api/kafebar"
)

type OrderService struct {
	builder sq.StatementBuilderType
}

var _ kafebar.OrderService = (*OrderService)(nil)

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{
		builder: sq.StatementBuilder.RunWith(sq.WrapStdSql(db)),
	}
}

func (o *OrderService) CreateOrder(ctx context.Context, order kafebar.Order) (kafebar.Order, error) {
	res, err := o.builder.Insert(tableOrders).
		Columns(columnName, columnStatus).
		Values(order.Name, order.Status).
		Exec()

	if err != nil {
		return order, fmt.Errorf("cannot insert order record: %w", err)
	}

	orderId, err := res.LastInsertId()
	if err != nil {
		return order, fmt.Errorf("cannot get created order id: %w", err)
	}
	order.Id = int(orderId)

	for i, item := range order.Items {
		item.OrderId = order.Id
		createdItem, err := o.AddOrderItem(ctx, item)
		if err != nil {
			return order, fmt.Errorf("cannot create order item: %w", err)
		}
		order.Items[i] = createdItem
	}

	return order, nil
}

func (o *OrderService) GetOrder(ctx context.Context, orderId int) (kafebar.Order, error) {
	orderRow := o.builder.
		Select(columnId, columnName, columnStatus).
		From(tableOrders).
		Where(sq.Eq{columnId: orderId}).
		QueryRow()

	var order kafebar.Order

	err := orderRow.Scan(&order.Id, &order.Name, &order.Status)
	if err != nil {
		return order, fmt.Errorf("cannot get order: %w", err)
	}

	itemRows, err := o.builder.
		Select(columnId, columnOrderId, columnProductId, columnStatus).
		From(tableOrderItems).
		Where(sq.Eq{columnOrderId: orderId}).
		Query()

	if err != nil {
		return order, fmt.Errorf("cannot get order_items: %w", err)
	}

	for itemRows.Next() {
		var item kafebar.OrderItem
		err := itemRows.Scan(&item.Id, &item.OrderId, &item.ProductId, &item.Status)
		if err != nil {
			return order, fmt.Errorf("cannot scan order item: %w", err)
		}
		order.Items = append(order.Items, item)
	}

	itemOptionRows, err := o.builder.
		Select(columnOrderItemId, columnOption).
		From(tableOrderItemOptions).
		Where(sq.Eq{columnOrderId: orderId}).
		Query()

	if err != nil {
		return order, fmt.Errorf("cannot get order_items: %w", err)
	}

	for itemOptionRows.Next() {
		var itemId int
		var option string
		err := itemOptionRows.Scan(&itemId, &option)
		if err != nil {
			return order, fmt.Errorf("cannot scan order item option: %w", err)
		}
		itemIdx := slices.IndexFunc(order.Items, func(i kafebar.OrderItem) bool { return i.Id == itemId })
		if itemIdx == -1 {
			return order, fmt.Errorf("found option for non existing item")
		}
		order.Items[itemIdx].Options = append(order.Items[itemIdx].Options, option)
	}

	return order, nil
}

func (o *OrderService) EditOrder(ctx context.Context, order kafebar.Order) error {
	return nil
}

func (o *OrderService) AddOrderItem(ctx context.Context, item kafebar.OrderItem) (kafebar.OrderItem, error) {
	res, err := o.builder.Insert(tableOrderItems).
		Columns(columnOrderId, columnProductId, columnStatus).
		Values(item.OrderId, item.ProductId, item.Status).
		Exec()

	if err != nil {
		return item, fmt.Errorf("cannot insert order record: %w", err)
	}

	itemId, err := res.LastInsertId()
	if err != nil {
		return item, fmt.Errorf("cannot get created order id: %w", err)
	}
	item.Id = int(itemId)

	if len(item.Options) == 0 {
		return item, nil
	}

	insertOptionsStmt := o.builder.Insert(tableOrderItemOptions).
		Columns(columnOrderId, columnOrderItemId, columnOption)

	for _, opt := range item.Options {
		insertOptionsStmt = insertOptionsStmt.Values(item.OrderId, item.Id, opt)
	}

	_, err = insertOptionsStmt.Exec()
	if err != nil {
		return item, fmt.Errorf("cannot insert item options: %w", err)
	}

	return item, nil
}

func (o *OrderService) RemoveOrderItem(ctx context.Context, orderItemId int) error {
	return nil
}

func (o *OrderService) EditOrderItem(ctx context.Context, item kafebar.OrderItem) error {
	return nil
}

func (o *OrderService) GetOrders(ctx context.Context) ([]kafebar.Order, error) {
	return nil, nil
}
