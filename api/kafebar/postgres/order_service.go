package postgres

import (
	"context"
	"fmt"
	"slices"

	sq "github.com/Masterminds/squirrel"
	"github.com/kafebar/kafebar/api/kafebar"
)

type OrderService struct {
	builder sq.StatementBuilderType
}

var _ kafebar.OrderService = (*OrderService)(nil)

func NewOrderService(db sq.StdSqlCtx) *OrderService {
	builder := sq.StatementBuilder.RunWith(sq.WrapStdSqlCtx(db)).PlaceholderFormat(sq.Dollar)

	return &OrderService{
		builder: builder,
	}
}

func (o *OrderService) CreateOrder(ctx context.Context, order kafebar.Order) (kafebar.Order, error) {
	err := o.builder.Insert(tableOrders).
		Columns(columnName, columnIsArchived).
		Values(order.Name, order.IsArchived).
		Suffix("RETURNING id").
		QueryRowContext(ctx).Scan(&order.Id)

	if err != nil {
		return order, fmt.Errorf("cannot insert order record: %w", err)
	}

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
	orders, err := o.getOrdersPredicate(ctx, sq.Eq{columnId: orderId})
	if err != nil {
		return kafebar.Order{}, err
	}
	if len(orders) < 1 {
		return kafebar.Order{}, fmt.Errorf("not found")
	}

	return orders[0], nil
}

func (o *OrderService) GetOrderItem(ctx context.Context, orderItemId int) (kafebar.OrderItem, error) {
	orderItems, err := o.getOrderItemsPredicate(ctx, sq.Eq{columnId: orderItemId})
	if err != nil {
		return kafebar.OrderItem{}, err
	}
	if len(orderItems) < 1 {
		return kafebar.OrderItem{}, fmt.Errorf("not found")
	}

	return orderItems[0], nil
}

func (o *OrderService) EditOrder(ctx context.Context, order kafebar.Order) error {
	return nil
}

func (o *OrderService) AddOrderItem(ctx context.Context, item kafebar.OrderItem) (kafebar.OrderItem, error) {
	err := o.builder.Insert(tableOrderItems).
		Columns(columnOrderId, columnProductId, columnStatus).
		Values(item.OrderId, item.ProductId, item.Status).
		Suffix("RETURNING id").
		QueryRowContext(ctx).Scan(&item.Id)

	if err != nil {
		return item, fmt.Errorf("cannot insert order record: %w", err)
	}

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

func (o *OrderService) UpdateOrderItemStatus(ctx context.Context, itemId int, status kafebar.Status) error {
	_, err := o.builder.Update(tableOrderItems).
		Set(columnStatus, status).
		Where(sq.Eq{columnId: itemId}).
		ExecContext(ctx)

	return err
}

func (o *OrderService) UpdateOrderArchiveStatus(ctx context.Context, orderId int, isArchived bool) error {
	_, err := o.builder.Update(tableOrders).
		Set(columnIsArchived, isArchived).
		Where(sq.Eq{columnId: orderId}).
		ExecContext(ctx)

	return err
}

func (o *OrderService) GetOrders(ctx context.Context) ([]kafebar.Order, error) {
	return o.getOrdersPredicate(ctx, nil)
}

func (o *OrderService) getOrdersPredicate(ctx context.Context, predicate any) ([]kafebar.Order, error) {
	orderRows, err := o.builder.
		Select(columnId, columnName, columnIsArchived).
		From(tableOrders).
		Where(predicate).
		OrderBy(columnId).
		QueryContext(ctx)

	if err != nil {
		return nil, fmt.Errorf("cannot fetch orders: %w", err)
	}

	orders := []kafebar.Order{}
	orderIds := []int{}

	for orderRows.Next() {
		var order kafebar.Order
		err := orderRows.Scan(&order.Id, &order.Name, &order.IsArchived)
		if err != nil {
			return nil, fmt.Errorf("cannot scan order item option: %w", err)
		}

		orders = append(orders, order)
		orderIds = append(orderIds, order.Id)
	}

	orderItems, err := o.getOrderItemsPredicate(ctx, sq.Eq{columnOrderId: orderIds})
	if err != nil {
		return nil, fmt.Errorf("cannot get order items: %w", err)
	}

	for _, orderItem := range orderItems {
		orderIdx := slices.IndexFunc(orders, func(o kafebar.Order) bool { return o.Id == orderItem.OrderId })
		if orderIdx == -1 {
			return orders, fmt.Errorf("found option for non existing item")
		}

		orders[orderIdx].Items = append(orders[orderIdx].Items, orderItem)
	}

	return orders, nil
}

func (o *OrderService) getOrderItemsPredicate(ctx context.Context, predicate any) ([]kafebar.OrderItem, error) {

	itemRows, err := o.builder.
		Select(columnId, columnOrderId, columnProductId, columnStatus).
		From(tableOrderItems).
		Where(predicate).
		OrderBy(columnId).
		QueryContext(ctx)

	if err != nil {
		return nil, fmt.Errorf("cannot get order_items: %w", err)
	}

	orderItems := []kafebar.OrderItem{}
	orderItemIds := []int{}

	for itemRows.Next() {
		var item kafebar.OrderItem
		err := itemRows.Scan(&item.Id, &item.OrderId, &item.ProductId, &item.Status)
		if err != nil {
			return orderItems, fmt.Errorf("cannot scan order item: %w", err)
		}
		orderItems = append(orderItems, item)
		orderItemIds = append(orderItemIds, item.Id)
	}

	itemOptions, err := o.getOrderItemOptionsPredicate(ctx, sq.Eq{
		columnOrderItemId: orderItemIds,
	})

	if err != nil {
		return nil, fmt.Errorf("cannot get order_item_options: %w", err)
	}

	for _, itemOption := range itemOptions {
		itemIdx := slices.IndexFunc(orderItems, func(i kafebar.OrderItem) bool { return i.Id == itemOption.orderItemId })
		if itemIdx == -1 {
			return orderItems, fmt.Errorf("found option for non existing item")
		}
		orderItems[itemIdx].Options = append(orderItems[itemIdx].Options, itemOption.option)
	}

	return orderItems, nil
}

type orderItemOption struct {
	id          int
	orderItemId int
	orderId     int
	option      string
}

func (o *OrderService) getOrderItemOptionsPredicate(ctx context.Context, predicate any) ([]orderItemOption, error) {
	itemOptionRows, err := o.builder.
		Select(columnId, columnOrderId, columnOrderItemId, columnOption).
		From(tableOrderItemOptions).
		Where(predicate).
		OrderBy(columnId).
		QueryContext(ctx)

	if err != nil {
		return nil, fmt.Errorf("cannot get order_items: %w", err)
	}

	orderItemOptions := []orderItemOption{}

	for itemOptionRows.Next() {
		var oiOpt orderItemOption
		err := itemOptionRows.Scan(&oiOpt.id, &oiOpt.orderId, &oiOpt.orderItemId, &oiOpt.option)
		if err != nil {
			return orderItemOptions, fmt.Errorf("cannot scan order item option: %w", err)
		}
		orderItemOptions = append(orderItemOptions, oiOpt)
	}

	return orderItemOptions, nil
}
