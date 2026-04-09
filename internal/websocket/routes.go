package websocket

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, hub *Hub, onSensorData func(SensorDataMsg), onEmergency func(EmergencyMsg)) {
	wsGroup := router.Group("/ws")
	{
		wsGroup.GET("/arduino", HandleArduinoWS(hub, onSensorData, onEmergency))
		wsGroup.GET("/mobile", HandleMobileWS(hub))
	}
}
