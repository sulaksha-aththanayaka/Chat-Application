# Go WebSocket Chat Application

A simple real-time chat server built with Go and the Gorilla WebSocket library. It allows multiple clients to connect and exchange messages through a central Hub.

---

## Project Structure

```
.
├── main.go       # HTTP server, WebSocket upgrader, connection handler
└── hub.go        # Hub struct with thread-safe client management
```

---

## How It Works

### WebSocket Upgrade
Incoming HTTP connections are upgraded to persistent WebSocket connections using the Gorilla upgrader. This allows full-duplex, real-time communication between the server and each client.

### Hub
A central `Hub` holds a map of all active client connections. It is responsible for:
- **Registering** new clients on connect
- **Unregistering** clients on disconnect
- **Broadcasting** messages from one client to all others

### Handler
Each client gets its own `handler()` running in a separate goroutine. It:
1. Upgrades the HTTP connection to WebSocket
2. Registers the connection with the Hub
3. Listens for incoming messages in an infinite loop
4. Broadcasts received messages to all other connected clients
5. Unregisters the client on disconnect or error

### Concurrency Safety
The Hub's client map is protected by a `sync.Mutex` to prevent race conditions from concurrent goroutine access. All map reads and writes happen through Hub methods (`Register`, `Unregister`, `Broadcast`) that handle locking internally.

---

## Getting Started

### Prerequisites
- Go 1.18+
- Gorilla WebSocket

```bash
go get github.com/gorilla/websocket
```

### Run the Server

```bash
go run .
```

The server starts on port `8080`.

---

## Testing

You can test the chat using **Postman**:

1. Open two Postman windows
2. In both, create a new **WebSocket** request
3. Connect both to `ws://localhost:8080/ws`
4. Send a message from one — it appears in the other

Alternatively, use `wscat` from the terminal:

```bash
npm install -g wscat
wscat -c ws://localhost:8080/ws
```

---

## Limitations & Next Steps

- No authentication — any client can connect
- No private messaging — all messages are broadcast to every connected client
- No message history — messages are not persisted
- No room/channel support
- Read limit is not set — clients can send arbitrarily large messages

---

## Dependencies

| Package | Purpose |
|--------|---------|
| [gorilla/websocket](https://github.com/gorilla/websocket) | WebSocket protocol implementation |
| `net/http` | HTTP server (standard library) |
| `sync` | Mutex for concurrent map access (standard library) |
