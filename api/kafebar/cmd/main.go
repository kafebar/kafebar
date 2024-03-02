package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kafebar/kafebar/api/kafebar/frontend"
	"github.com/kafebar/kafebar/api/kafebar/orderhandler"
	"github.com/kafebar/kafebar/api/kafebar/producthandler"
	"github.com/kafebar/kafebar/api/kafebar/sqlite"
	"github.com/kafebar/kafebar/api/kafebar/sse"
	"github.com/rs/cors"

	_ "github.com/mattn/go-sqlite3"
)

var (
	port   = os.Getenv("PORT")
	uiPath = os.Getenv("UI_PATH")

	tlsKeyFile  = os.Getenv("TLS_KEY_FILE")
	tlsCertFile = os.Getenv("TLS_CERT_FILE")

	migrationsDirectory = os.Getenv("MIGRATIONS_DIRECTORY")
	sqliteLocation      = os.Getenv("SQLITE_LOCATION")
)

func main() {
	mux := http.NewServeMux()

	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s", sqliteLocation))
	if err != nil {
		panic(fmt.Errorf("cannot open db: %w", err))
	}

	err = sqlite.RunMigrations(db, migrationsDirectory)
	if err != nil {
		panic(fmt.Errorf("cannot run migrations: %w", err))
	}

	serverEvents := sse.NewService()
	products := sqlite.NewProductService(db)
	orders := sqlite.NewOrderService(db)

	productHandler := producthandler.New(products, serverEvents)
	orderHandler := orderhandler.New(orders, serverEvents)

	mux.Handle("/api/events", serverEvents)

	mux.Handle("/api/orders", orderHandler)
	mux.Handle("/api/orders/", orderHandler)

	mux.Handle("/api/products", productHandler)
	mux.Handle("/api/products/", productHandler)

	if uiPath != "" {
		mux.Handle("/", frontend.NewHandler(uiPath))
	}

	addr := fmt.Sprintf(":%s", port)

	fmt.Printf("starting server on %s\n", addr)

	handler := cors.AllowAll().Handler(mux)

	if tlsCertFile != "" {
		err = http.ListenAndServeTLS(addr, tlsCertFile, tlsKeyFile, handler)
	} else {
		err = http.ListenAndServe(addr, handler)
	}

	if err != nil {
		log.Fatalf("cannot start server: %s", err.Error())
	}
}
