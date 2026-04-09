package websocket

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Tiempo permitido para escribir el mensaje al peer.
	writeWait = 10 * time.Second

	// Tiempo permitido para leer el próximo pong desde el peer.
	pongWait = 60 * time.Second

	// Enviar pings al peer con este intervalo. Debe ser menor a pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Tamaño máximo permitido del mensaje recibido por el peer.
	maxMessageSize = 4096 // Aumentable si es necesario (video/audio, etc)
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client es un interlocutor de la conexión WebSocket (Arduino / Python / App).
type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte

	// Metadatos
	ID             string // Puede ser UserID, DeviceID o MacAddress
	Role           string // 'mobile_app', 'arduino_node'
	IntersectionID string // ID al que está emparejado para recibir u originar actualizaciones
}

// readPump fluye mensajes desde la conexión websocket al hub.
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, messageBytes, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error de lectura websocket inesperado: %v", err)
			}
			break
		}

		// Decodificar mensaje
		var wsMsg WSMessage
		if err := json.Unmarshal(messageBytes, &wsMsg); err != nil {
			log.Printf("error decodificando json en WS Client: %v", err)
			continue
		}

		// (En el futuro aquí llamar a servicios / casos de uso dependiendo del mensaje recibido)
		// - Si es sensor_data de Arduino, retransformar y llamar al Hub.Broadcast.
		// - Si es message desde Mobile, pasar a Arduino respectivo.

		// Por ahora: broadcast puro en crudo para debugging
		c.Hub.Broadcast <- messageBytes
	}
}

// writePump fluye mensajes desde el hub al cliente websocket.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// El hub cerró el channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Añadir los mensajes pendientes en cola al current websocket
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
