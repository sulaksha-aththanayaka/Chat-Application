package main

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

func (h *Hub) Register(conn *websocket.Conn) {
	h.mu.Lock()
	h.clients[conn] = true
	h.mu.Unlock()
}

func (h *Hub) Unregister(conn *websocket.Conn) {
	h.mu.Lock()
	delete(h.clients, conn)
	h.mu.Unlock()
}

func (h *Hub) Broadcast(conn *websocket.Conn, messageType int, p []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for key := range h.clients {
		if key != conn {
			if err := key.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
		}
	}
}
