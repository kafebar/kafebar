package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kafebar/kafebar/api/kafebar/frontend"
	"github.com/kafebar/kafebar/api/kafebar/orderhandler"
	"github.com/kafebar/kafebar/api/kafebar/postgres"
	"github.com/kafebar/kafebar/api/kafebar/producthandler"
	"github.com/kafebar/kafebar/api/kafebar/sse"
	"github.com/rs/cors"

	_ "github.com/jackc/pgx/stdlib"
)

var (
	port   = os.Getenv("PORT")
	uiPath = os.Getenv("UI_PATH")

	postgresConnString = os.Getenv("POSTGRES_CONNSTRING")
)

func main() {
	mux := http.NewServeMux()

	db, err := sql.Open("pgx", postgresConnString)
	if err != nil {
		log.Fatalf("cannot open db: %s", err.Error())
	}

	serverEvents := sse.NewService()
	products := postgres.NewProductService(db)
	orders := postgres.NewOrderService(db)

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

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("cannot start server: %s", err.Error())
	}
}
