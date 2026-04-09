package routes

import (
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/middleware"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/infra/controllers"
	"github.com/gin-gonic/gin"
)

func SetupSensorRoutes(
	r *gin.Engine,
	registerCtrl *controllers.RegisterSensorController,
	listCtrl *controllers.ListSensorsController,
	updateCtrl *controllers.UpdateSensorController,
	deleteCtrl *controllers.DeleteSensorController,
	recordReadingCtrl *controllers.RecordReadingController,
	listReadingsCtrl *controllers.ListReadingsController,
	jwtSecret string,
) {
	api := r.Group("/sensors")

	// IoT Endpoints de alta frecuencia
	api.POST("/:id/readings", recordReadingCtrl.Handle)
	api.GET("", listCtrl.Handle)
	api.GET("/:id/readings", listReadingsCtrl.Handle)

	// Admin / Operator endpoints
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(jwtSecret))
	{
		protected.POST("", registerCtrl.Handle)
		protected.PUT("/:id", updateCtrl.Handle)
		protected.DELETE("/:id", deleteCtrl.Handle)
	}
}
