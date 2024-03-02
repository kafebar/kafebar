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

	EventTypeOrderCreated EventType = "ProductCreated"
	EventTypeOrderUpdated EventType = "ProductUpdated"
	EventTypeOrderDeleted EventType = "ProductDeleted"
)

type EventsService interface {
	http.Handler
	Broadcast(context.Context, Event) error
}
