package sse

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"sync"

	"github.com/google/uuid"
	"github.com/kafebar/kafebar/api/kafebar"
)

type Service struct {
	connMutex sync.RWMutex
	conns     []Conn
}

type Conn struct {
	id     uuid.UUID
	writer http.ResponseWriter
}

var _ kafebar.EventsService = (*Service)(nil)

func NewService() *Service {
	return &Service{}
}

func (s *Service) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	w.(http.Flusher).Flush()

	connUid := s.addConnection(w)

	<-req.Context().Done()
	s.removeConnection(connUid)
}

func (s *Service) addConnection(writer http.ResponseWriter) uuid.UUID {
	s.connMutex.Lock()
	defer s.connMutex.Unlock()

	fmt.Println("adding connection")

	connUid := uuid.New()
	s.conns = append(s.conns, Conn{connUid, writer})
	return connUid
}

func (s *Service) removeConnection(connUid uuid.UUID) {
	s.connMutex.Lock()
	defer s.connMutex.Unlock()
	fmt.Println("removing connection")

	s.conns = slices.DeleteFunc(s.conns, func(c Conn) bool {
		return c.id == connUid
	})
}

func (s *Service) Broadcast(ctx context.Context, event kafebar.Event) error {
	s.connMutex.RLock()
	defer s.connMutex.RUnlock()

	eventStr, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("cannot marshal data: %w", err)
	}
	sseEvent := fmt.Sprintf("event: message\ndata: %s\n\n", eventStr)

	for _, conn := range s.conns {
		fmt.Fprint(conn.writer, sseEvent)
		conn.writer.(http.Flusher).Flush()
	}

	return nil
}
