package orderhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	kafebar "github.com/kafebar/kafebar/api/kafebar"
)

type handler struct {
	os kafebar.OrderService
	es kafebar.EventsService
}

func New(os kafebar.OrderService, es kafebar.EventsService) http.Handler {
	mux := http.NewServeMux()

	h := &handler{os, es}

	mux.HandleFunc("GET /api/orders", h.getOrders)
	mux.HandleFunc("POST /api/orders", h.createOrder)
	mux.HandleFunc("PUT /api/orders/{orderId}/isArchived/{isArchived}", h.UpdateOrderArchiveStatus)
	mux.HandleFunc("PUT /api/orders/items/{orderItemId}/status/{status}", h.updateOrderItemStatus)

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

func (h *handler) updateOrderItemStatus(w http.ResponseWriter, req *http.Request) {
	orderItemId, err := strconv.Atoi(req.PathValue("orderItemId"))
	if err != nil {
		fmt.Println("invalid path parameter: ", err)
		http.Error(w, "invalid path parameter", http.StatusBadRequest)
		return
	}

	status := kafebar.Status(req.PathValue("status"))

	orderItem, err := h.os.GetOrderItem(req.Context(), orderItemId)
	if err != nil {
		fmt.Println("cannot get order item: ", err)
		http.Error(w, "cannot get order item", http.StatusInternalServerError)
		return
	}

	err = h.os.UpdateOrderItemStatus(req.Context(), orderItemId, status)
	if err != nil {
		fmt.Println("cannot update order item status: ", err)
		http.Error(w, "cannot update order item status", http.StatusInternalServerError)
		return
	}

	order, err := h.os.GetOrder(req.Context(), orderItem.OrderId)
	if err != nil {
		fmt.Println("cannot get order: ", err)
		http.Error(w, "cannot get order", http.StatusInternalServerError)
		return
	}

	h.es.Broadcast(req.Context(), kafebar.Event{
		Type: kafebar.EventTypeOrderUpdated,
		Data: order,
	})
}

func (h *handler) UpdateOrderArchiveStatus(w http.ResponseWriter, req *http.Request) {
	orderId, err := strconv.Atoi(req.PathValue("orderId"))
	if err != nil {
		fmt.Println("invalid path parameter: ", err)
		http.Error(w, "invalid path parameter", http.StatusBadRequest)
		return
	}

	isArchived, _ := strconv.ParseBool(req.PathValue("isArchived"))

	err = h.os.UpdateOrderArchiveStatus(req.Context(), orderId, isArchived)
	if err != nil {
		fmt.Println("cannot update order archive status: ", err)
		http.Error(w, "cannot update order archive status", http.StatusInternalServerError)
		return
	}

	order, err := h.os.GetOrder(req.Context(), orderId)
	if err != nil {
		fmt.Println("cannot get order: ", err)
		http.Error(w, "cannot get order", http.StatusInternalServerError)
		return
	}

	h.es.Broadcast(req.Context(), kafebar.Event{
		Type: kafebar.EventTypeOrderUpdated,
		Data: order,
	})
}
