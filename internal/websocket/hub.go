package websocket

import "sync"

type Hub struct {
	MobileClients map[*Client]bool
	ArduinoClient *Client

	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan WSMessage
	mu         sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		MobileClients: make(map[*Client]bool),
		Register:      make(chan *Client),
		Unregister:    make(chan *Client),
		Broadcast:     make(chan WSMessage, 256),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			if client.IsArduino {
				h.ArduinoClient = client
			} else {
				h.MobileClients[client] = true
			}
			h.mu.Unlock()

		case client := <-h.Unregister:
			h.mu.Lock()
			if client.IsArduino {
				if h.ArduinoClient == client {
					h.ArduinoClient = nil
				}
			} else {
				delete(h.MobileClients, client)
			}
			h.mu.Unlock()

		case msg := <-h.Broadcast:
			h.mu.RLock()
			for client := range h.MobileClients {
				select {
				case client.Send <- msg:
				default:
					close(client.Send)
					delete(h.MobileClients, client)
				}
			}
			h.mu.RUnlock()
		}
	}
}

func (h *Hub) BroadcastToMobile(msg WSMessage) {
	h.Broadcast <- msg
}
