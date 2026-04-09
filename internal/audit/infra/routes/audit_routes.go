package routes

import (
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/audit/infra/controllers"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuditRoutes(
	router *gin.Engine,
	jwtSecret string,
	getLogsController *controllers.GetLogsController,
	createLogController *controllers.CreateLogController,
) {
	auditGroup := router.Group("/audit", middleware.AuthMiddleware(jwtSecret))
	{
		auditGroup.GET("/logs", getLogsController.Handle)
		auditGroup.POST("/logs", createLogController.Handle)
	}
}
