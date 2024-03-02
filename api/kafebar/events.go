package kafebar

import (
	"context"
	"net/http"
)

type Event struct {
	Type EventType `json:"type"`
	Data any       `json:"data"`
}

type EventType string

const (
	EventTypeProductCreated EventType = "ProductCreated"
	EventTypeProductUpdated EventType = "ProductUpdated"
	EventTypeProductDeleted EventType = "ProductDeleted"

	EventTypeOrderCreated EventType = "OrderCreated"
	EventTypeOrderUpdated EventType = "OrderUpdated"
	EventTypeOrderDeleted EventType = "OrderDeleted"
)

type EventsService interface {
	http.Handler
	Broadcast(context.Context, Event) error
}
