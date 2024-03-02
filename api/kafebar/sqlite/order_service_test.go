package sqlite

import (
	"context"
	"testing"

	kafebar "github.com/kafebar/kafebar/api/kafebar"
	"github.com/stretchr/testify/assert"
)

func TestOrderSErvice(t *testing.T) {
	os := NewOrderService(getTestDb())

	order := kafebar.Order{
		Name: "test",
		Items: []kafebar.OrderItem{
			{
				ProductId: 123,
				Status:    kafebar.StatusInProgress,
				Options:   []string{"abc"},
			},
		},
	}

	createdOrder, err := os.CreateOrder(context.Background(), order)

	assert.NoError(t, err)
	assert.NotEmpty(t, createdOrder.Id)

	fetchedOrder, err := os.GetOrder(context.Background(), createdOrder.Id)
	assert.NoError(t, err)

	assert.Equal(t, createdOrder, fetchedOrder)
}
