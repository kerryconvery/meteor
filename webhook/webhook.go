package webhook

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// Webhook represents a webhook
type Webhook struct {
	clients  map[*websocket.Conn]bool
	messages chan interface{}
	upgrader websocket.Upgrader
}

// New returns a new instance of Webhook
func New() Webhook {
	return Webhook{
		clients:  make(map[*websocket.Conn]bool),
		messages: make(chan interface{}, 100),
		upgrader: websocket.Upgrader{},
	}
}

// AddClient adds a client into the list of clients listening for messages
func (wh *Webhook) AddClient(w http.ResponseWriter, r *http.Request) error {
	ws, err := wh.upgrader.Upgrade(w, r, nil)

	if err != nil {
		return err
	}

	//defer ws.Close()

	wh.clients[ws] = true

	return nil
}

// Broadcast adds a message into the queue to be sent to all connected clients
func (wh *Webhook) Broadcast(payload interface{}) {
	wh.messages <- payload
}

// Start processing messages
func (wh *Webhook) Start() {
	go processQueue(wh.clients, wh.messages)
}

// Stop the processing new messages
func (wh *Webhook) Stop() {
	close(wh.messages)
}

func processQueue(clients map[*websocket.Conn]bool, messages chan interface{}) {
	for message := range messages {
		for client := range clients {
			err := client.WriteJSON(message)

			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
