package sqlite

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"

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

func RunMigrations(db *sql.DB, migrationsDir string) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("cannot create database driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsDir),
		"sqlite3", driver)

	if err != nil {
		return fmt.Errorf("cannot create migration runner: %w", err)
	}
	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("cannot run migrations instance: %w", err)
	}

	return nil
}
