package websocket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleArduinoWS(hub *Hub, onSensorData func(SensorDataMsg), onEmergency func(EmergencyMsg)) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("WS upgrade error: %v", err)
			return
		}

		client := &Client{
			Hub:       hub,
			Conn:      conn,
			Send:      make(chan WSMessage, 256),
			IsArduino: true,
		}

		hub.Register <- client
		go client.WritePump()

		client.ReadPump(func(msg WSMessage) {
			log.Printf("[WS Arduino] Recibido: type=%s", msg.Type)

			switch msg.Type {
			case "sensor_data":
				var data SensorDataMsg
				if err := json.Unmarshal(msg.Data, &data); err == nil {
					log.Printf("[WS Arduino] Sensor data: intersection=%s fase=%s", data.IntersectionID, data.Fase)
					if onSensorData != nil {
						onSensorData(data)
					}

					// Convertimos el struct de vuelta a raw JSON o enviamos directo a móviles
					// ya que WSMessage.Data ahora es json.RawMessage
					raw, _ := json.Marshal(data)
					hub.BroadcastToMobile(WSMessage{
						Type: "intersection_update",
						Data: raw,
					})
				} else {
					log.Printf("[WS Arduino] Error al parsear sensor_data: %v", err)
				}

			case "emergency":
				var data EmergencyMsg
				if err := json.Unmarshal(msg.Data, &data); err == nil {
					log.Printf("[WS Arduino] Emergency data: intersection=%s active=%v", data.IntersectionID, data.Active)
					if onEmergency != nil {
						onEmergency(data)
					}

					raw, _ := json.Marshal(data)
					hub.BroadcastToMobile(WSMessage{
						Type: "emergency_alert",
						Data: raw,
					})
				} else {
					log.Printf("[WS Arduino] Error al parsear emergency: %v", err)
				}
			}
		})
	}
}

func HandleMobileWS(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		// // TODO: verificar JWT desde query auth token si es necesario
		token := c.Query("token")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized, token required"})
			return
		}
		// (Aquí iría la validación del JWT usando jwtSecret)

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("Mobile WS upgrade error: %v", err)
			return
		}

		client := &Client{
			Hub:       hub,
			Conn:      conn,
			Send:      make(chan WSMessage, 256),
			IsArduino: false,
		}

		hub.Register <- client
		go client.WritePump()

		client.ReadPump(func(msg WSMessage) {
			switch msg.Type {
			case "trigger_emergency", "adjust_cycle":
				// Reenviar app -> arduino
				hub.mu.RLock()
				if hub.ArduinoClient != nil {
					// mandarlo para el Arduino
					hub.ArduinoClient.Send <- msg
				}
				hub.mu.RUnlock()
			}
		})
	}
}
