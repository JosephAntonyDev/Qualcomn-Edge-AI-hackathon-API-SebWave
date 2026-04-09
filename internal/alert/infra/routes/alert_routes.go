package routes

import (
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/infra/controllers"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAlertRoutes(
	r *gin.Engine,
	createCtrl *controllers.CreateAlertController,
	getCtrl *controllers.GetAlertController,
	listCtrl *controllers.ListAlertsController,
	updateCtrl *controllers.UpdateAlertController,
	deleteCtrl *controllers.DeleteAlertController,
	changeStatusCtrl *controllers.ChangeStatusAlertController,
	jwtSecret string,
) {
	alertGroup := r.Group("/alerts")
	alertGroup.Use(middleware.AuthMiddleware(jwtSecret))
	{
		alertGroup.POST("/", createCtrl.Handle)
		alertGroup.GET("/:id", getCtrl.Handle)
		alertGroup.GET("/", listCtrl.Handle)
		alertGroup.PUT("/:id", updateCtrl.Handle)
		alertGroup.DELETE("/:id", deleteCtrl.Handle)
		alertGroup.PATCH("/:id/status", changeStatusCtrl.Handle)
	}
}
