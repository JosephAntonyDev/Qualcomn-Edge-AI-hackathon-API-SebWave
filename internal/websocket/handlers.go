package websocket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ServeWS maneja las peticiones WebSocket desde clientes (móviles) o (Hardware IoT Python/Arduino)
// Se invoca un Upgrade en el handshake de conectividad initial.
func ServeWS(hub *Hub, c *gin.Context) {
	// 1. Extraer queries esenciales como identidad y rol (Mobile vs Hardware)
	role := c.Query("role")                      // 'arduino' o 'mobile'
	intersectionID := c.Query("intersection_id") // La intersección que observará o servirá
	clientID := c.Query("client_id")             // Hardware ID o JWT Sub UUID

	if role == "" || clientID == "" {
		log.Println("Roles o Client_IDs faltantes al conectar WS.")
		c.JSON(http.StatusBadRequest, gin.H{"error": "role y client_id mandatorios"})
		return
	}

	// 2. Upgradear Check Origin: Idealmente en prod restringir dominio,
	// para hackathons lo marcamos true de todos.
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Fallo al hacer upgrade WebSocket: %v", err)
		return
	}

	// 3. Crear wrapper de Cliente
	client := &Client{
		Hub:            hub,
		Conn:           conn,
		Send:           make(chan []byte, 256),
		ID:             clientID,
		Role:           role,
		IntersectionID: intersectionID,
	}

	// 4. Registrar al hub en otra goroutine para no bloquear
	client.Hub.Register <- client

	// 5. Iniciar loops infinitos de I/O en Goroutines.
	// Allow collection of memory referenced by the caller
	// (run on separate tasks via Go rutines)
	go client.WritePump()
	go client.ReadPump()
}
