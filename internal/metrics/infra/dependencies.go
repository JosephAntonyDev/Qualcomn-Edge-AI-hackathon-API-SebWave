package infra

import (
	"database/sql"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/infra/controllers"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/infra/repository"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/infra/routes"
	"github.com/gin-gonic/gin"
)

func SetupDependencies(r *gin.Engine, db *sql.DB, jwtSecret string) {
	// Repositories
	repo := repository.NewPostgresMetricsRepository(db)

	// Use Cases
	listDailyUC := app.NewListDailyMetricsUseCase(repo)
	listSystemUC := app.NewListSystemMetricsUseCase(repo)
	saveDailyUC := app.NewSaveDailyMetricUseCase(repo)
	saveSystemUC := app.NewSaveSystemMetricUseCase(repo)

	// Controllers
	listDailyCtrl := controllers.NewListDailyMetricsController(listDailyUC)
	listSystemCtrl := controllers.NewListSystemMetricsController(listSystemUC)
	saveDailyCtrl := controllers.NewSaveDailyMetricController(saveDailyUC)
	saveSystemCtrl := controllers.NewSaveSystemMetricController(saveSystemUC)

	// Routes
	routes.RegisterMetricsRoutes(
		r,
		listDailyCtrl,
		listSystemCtrl,
		saveDailyCtrl,
		saveSystemCtrl,
		jwtSecret,
	)
}
