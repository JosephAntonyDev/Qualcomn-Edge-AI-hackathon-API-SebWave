package routes

import (
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/infra/controllers"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterMetricsRoutes(
	r *gin.Engine,
	listDailyCtrl *controllers.ListDailyMetricsController,
	listSystemCtrl *controllers.ListSystemMetricsController,
	saveDailyCtrl *controllers.SaveDailyMetricController,
	saveSystemCtrl *controllers.SaveSystemMetricController,
	jwtSecret string,
) {
	metricsGroup := r.Group("/metrics")
	metricsGroup.Use(middleware.AuthMiddleware(jwtSecret))
	{
		metricsGroup.GET("/daily", listDailyCtrl.Handle)
		metricsGroup.POST("/daily", saveDailyCtrl.Handle)

		metricsGroup.GET("/system", listSystemCtrl.Handle)
		metricsGroup.POST("/system", saveSystemCtrl.Handle)
	}
}
