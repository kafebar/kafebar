package sqlite

import (
	"context"
	"testing"

	kafebar "github.com/kafebar/kafebar/api/kafebar"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	ps := NewProductService(getTestDb())

	product := kafebar.Product{
		Name: "test",
		AvailableOptions: []string{
			"ExtraSugar",
			"AlmondMilk",
			"ExtraMilk",
		},
	}

	createdProduct, err := ps.CreateProduct(context.Background(), product)
	assert.NoError(t, err)
	assert.Equal(t, product.Name, createdProduct.Name)
	assert.NotEmpty(t, createdProduct.Id)

	fetchedProducts, err := ps.GetProducts(context.Background())
	assert.NoError(t, err)
	assert.Len(t, fetchedProducts, 1)
	assert.Equal(t, createdProduct, fetchedProducts[0])

}
