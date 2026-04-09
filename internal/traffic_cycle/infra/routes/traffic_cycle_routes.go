package routes

import (
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/middleware"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/infra/controllers"
	"github.com/gin-gonic/gin"
)

func SetupTrafficCycleRoutes(
	r *gin.Engine,
	recordCtrl *controllers.RecordCycleController,
	getCtrl *controllers.GetCycleController,
	listCtrl *controllers.ListCyclesController,
	deleteCtrl *controllers.DeleteCycleController,
	jwtSecret string,
) {
	api := r.Group("/traffic-cycles")

	// Endpoint IoT de alta frecuencia (Registro al finalizar ciclo) - Abierto o protegido con token IoT, aquí lo dejamos abierto por MVP
	api.POST("", recordCtrl.Handle)

	// Endpoints públicos de GET
	api.GET("", listCtrl.Handle)
	api.GET("/:id", getCtrl.Handle)

	// Admin / Operator endpoints
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(jwtSecret))
	{
		protected.DELETE("/:id", deleteCtrl.Handle)
	}
}
