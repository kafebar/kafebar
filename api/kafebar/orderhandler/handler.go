package orderhandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	kafebar "github.com/kafebar/kafebar/api/kafebar"
)

type handler struct {
	os kafebar.OrderService
	es kafebar.EventsService
}

func New(os kafebar.OrderService, es kafebar.EventsService) http.Handler {
	mux := http.NewServeMux()

	h := &handler{os, es}

	mux.HandleFunc("GET /orders", h.getOrders)
	mux.HandleFunc("POST /orders", h.createOrder)

	return mux
}

func (h *handler) getOrders(w http.ResponseWriter, req *http.Request) {
	orders, err := h.os.GetOrders(req.Context())
	if err != nil {
		http.Error(w, "cannot get orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func (h *handler) createOrder(w http.ResponseWriter, req *http.Request) {
	var order kafebar.Order

	err := json.NewDecoder(req.Body).Decode(&order)
	if err != nil {
		http.Error(w, "invalid order body: "+err.Error(), http.StatusBadRequest)
		return
	}

	createdOrder, err := h.os.CreateOrder(req.Context(), order)
	if err != nil {
		fmt.Println("cannot create order: ", err)
		http.Error(w, "cannot create order", http.StatusInternalServerError)
		return
	}

	h.es.Broadcast(req.Context(), kafebar.Event{
		Type: kafebar.EventTypeOrderCreated,
		Data: createdOrder,
	})

	json.NewEncoder(w).Encode(createdOrder)
}
