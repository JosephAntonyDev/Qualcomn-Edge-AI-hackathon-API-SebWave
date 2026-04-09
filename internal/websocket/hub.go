package websocket

import (
	"log"
)

// Hub mantiene el conjunto de clientes (Apps Móviles y Arduinos) activos
// y transmite mensajes a todos o a grupos específicos.
type Hub struct {
	// Clientes registrados.
	Clients map[*Client]bool

	// Mensajes entrantes para broadcast general (puede ser modificado para enviar por IntersectionID).
	Broadcast chan []byte

	// Registrar peticiones de clientes.
	Register chan *Client

	// Desregistrar peticiones de clientes.
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			log.Printf("Nuevo cliente registrado: %s (%s) en intersección: %s", client.ID, client.Role, client.IntersectionID)

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
				log.Printf("Cliente desconectado: %s (%s)", client.ID, client.Role)
			}

		case message := <-h.Broadcast:
			// En un futuro, el broadcast general puede refinarse para enviar solo
			// a los clientes móviles suscritos a un "IntersectionID" particular.
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
