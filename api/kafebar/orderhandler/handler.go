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
		err := fmt.Errorf("cannot get orders: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func (h *handler) createOrder(w http.ResponseWriter, req *http.Request) {
	var order kafebar.Order

	err := json.NewDecoder(req.Body).Decode(&order)
	if err != nil {
		err := fmt.Errorf("invalid order body: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdOrder, err := h.os.CreateOrder(req.Context(), order)
	if err != nil {
		err := fmt.Errorf("cannot create order: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		err := fmt.Errorf("invalid path parameter: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status := kafebar.Status(req.PathValue("status"))

	orderItem, err := h.os.GetOrderItem(req.Context(), orderItemId)
	if err != nil {
		err := fmt.Errorf("cannot get order item: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.os.UpdateOrderItemStatus(req.Context(), orderItemId, status)
	if err != nil {
		err := fmt.Errorf("cannot update order item status: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	order, err := h.os.GetOrder(req.Context(), orderItem.OrderId)
	if err != nil {
		err := fmt.Errorf("cannot get order: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		err := fmt.Errorf("invalid path parameter: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isArchived, _ := strconv.ParseBool(req.PathValue("isArchived"))

	err = h.os.UpdateOrderArchiveStatus(req.Context(), orderId, isArchived)
	if err != nil {

		err := fmt.Errorf("cannot update order archive status: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	order, err := h.os.GetOrder(req.Context(), orderId)
	if err != nil {
		err := fmt.Errorf("cannot get order: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.es.Broadcast(req.Context(), kafebar.Event{
		Type: kafebar.EventTypeOrderUpdated,
		Data: order,
	})
}
