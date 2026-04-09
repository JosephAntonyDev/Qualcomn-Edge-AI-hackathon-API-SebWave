package infra

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/audit/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/audit/infra/controllers"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/audit/infra/repository"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/audit/infra/routes"
)

func SetupDependencies(router *gin.Engine, db *sql.DB, jwtSecret string) {
	// Repositories
	repo := repository.NewPostgresAuditRepository(db)

	// UseCases
	createLogUseCase := app.NewCreateLogUseCase(repo)
	getLogsUseCase := app.NewGetLogsUseCase(repo)

	// Controllers
	createLogController := controllers.NewCreateLogController(createLogUseCase)
	getLogsController := controllers.NewGetLogsController(getLogsUseCase)

	// Routes
	routes.RegisterAuditRoutes(
		router,
		jwtSecret,
		getLogsController,
		createLogController,
	)
}
