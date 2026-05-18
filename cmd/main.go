package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var hub = Hub{
	clients: make(map[*websocket.Conn]bool),
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // Upgrade HTTP connection to WebSocket connection

	if err != nil {
		log.Println(err)
		return
	}

	hub.Register(conn)

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			hub.Unregister(conn)
			return
		}

		hub.Broadcast(conn, messageType, p)

	}
}

func main() {
	http.HandleFunc("/ws", handler)
	http.ListenAndServe(":8080", nil)
}
