package routes

import (
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/infra/controllers"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterEmergencyRoutes(
	r *gin.Engine,
	recordCtrl *controllers.RecordEmergencyController,
	getCtrl *controllers.GetEmergencyController,
	listCtrl *controllers.ListEmergenciesController,
	activateCtrl *controllers.ActivateCorridorController,
	deactivateCtrl *controllers.DeactivateCorridorController,
	deleteCtrl *controllers.DeleteEmergencyController,
	jwtSecret string,
) {
	emergencyGroup := r.Group("/emergencies")
	emergencyGroup.Use(middleware.AuthMiddleware(jwtSecret))
	{
		emergencyGroup.POST("/", recordCtrl.Handle)
		emergencyGroup.GET("/:id", getCtrl.Handle)
		emergencyGroup.GET("/", listCtrl.Handle)
		emergencyGroup.DELETE("/:id", deleteCtrl.Handle)

		emergencyGroup.POST("/:id/corridor/activate", activateCtrl.Handle)
		emergencyGroup.POST("/:id/corridor/deactivate", deactivateCtrl.Handle)
	}
}
