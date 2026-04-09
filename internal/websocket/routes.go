package websocket

import (
	"github.com/gin-gonic/gin"
)

// RegisterWebSocketRoutes conecta el Endpoint `/ws` global a las rutinas de gorilla/websockets publicadas.
// NOTA IMPORTANTE: Para este endpoint *no* usamos auth middleware HTTP normal
// en la primera parte porque los Headers HTTP raramente viajan fáciles en clientes como python websockets origin.
// Se recomienda pasar tokens por la query URL `?token=...` en futuros pasos si lo necesitas.
func RegisterWebSocketRoutes(r *gin.Engine, hub *Hub) {
	r.GET("/ws", func(c *gin.Context) {
		ServeWS(hub, c)
	})
}
