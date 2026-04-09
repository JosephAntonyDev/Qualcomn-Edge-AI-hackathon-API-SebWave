package routes

import (
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/infra/controllers"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupIntersectionRoutes(
	r *gin.Engine,
	createCtrl *controllers.CreateIntersectionController,
	getCtrl *controllers.GetIntersectionController,
	listCtrl *controllers.ListIntersectionsController,
	updateCtrl *controllers.UpdateIntersectionController,
	deleteCtrl *controllers.DeleteIntersectionController,
	stateCtrl *controllers.UpdateStateController,
	hbCtrl *controllers.HeartbeatController,
	jwtSecret string,
) {
	api := r.Group("/intersections")

	api.POST("/heartbeat", hbCtrl.Handle)
    api.GET("", listCtrl.Handle)
	api.GET("/:id", getCtrl.Handle)
	// Admin & Operators Endpoints
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(jwtSecret))
	{
		
		protected.POST("", createCtrl.Handle)
		protected.PUT("/:id", updateCtrl.Handle)
		protected.DELETE("/:id", deleteCtrl.Handle)

		protected.PATCH("/:id/state", stateCtrl.Handle)
	}
}
