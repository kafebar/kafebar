package postgres

import (
	"context"
	"database/sql"
	"os"
	"testing"

	kafebar "github.com/kafebar/kafebar/api/kafebar"
	"github.com/stretchr/testify/assert"

	_ "github.com/jackc/pgx/stdlib"
)

var postgresConnString = os.Getenv("POSTGRES_CONNSTRING")

func TestCreateProduct(t *testing.T) {
	ps := NewProductService(getTestDb())

	product := kafebar.Product{
		Name:  "test",
		Price: 7,
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
	assert.Contains(t, fetchedProducts, createdProduct)
}

func getTestDb() *sql.DB {
	db, err := sql.Open("pgx", postgresConnString)
	if err != nil {
		panic(err)
	}

	return db
}
