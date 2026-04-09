package infra

import (
	"database/sql"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/infra/controllers"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/infra/repository"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/infra/routes"
	"github.com/gin-gonic/gin"
)

func SetupDependencies(r *gin.Engine, db *sql.DB, jwtSecret string) {
	// Repositories
	repo := repository.NewPostgresEmergencyRepository(db)

	// Use Cases
	recordUC := app.NewRecordEmergencyUseCase(repo)
	getUC := app.NewGetEmergencyUseCase(repo)
	listUC := app.NewListEmergenciesUseCase(repo)
	activateUC := app.NewActivateCorridorUseCase(repo)
	deactivateUC := app.NewDeactivateCorridorUseCase(repo)
	deleteUC := app.NewDeleteEmergencyUseCase(repo)

	// Controllers
	recordCtrl := controllers.NewRecordEmergencyController(recordUC)
	getCtrl := controllers.NewGetEmergencyController(getUC)
	listCtrl := controllers.NewListEmergenciesController(listUC)
	activateCtrl := controllers.NewActivateCorridorController(activateUC)
	deactivateCtrl := controllers.NewDeactivateCorridorController(deactivateUC)
	deleteCtrl := controllers.NewDeleteEmergencyController(deleteUC)

	// Routes
	routes.RegisterEmergencyRoutes(
		r,
		recordCtrl,
		getCtrl,
		listCtrl,
		activateCtrl,
		deactivateCtrl,
		deleteCtrl,
		jwtSecret,
	)
}
