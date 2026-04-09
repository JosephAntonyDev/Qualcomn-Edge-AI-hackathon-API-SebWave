package infra

import (
	"database/sql"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/infra/controllers"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/infra/repository"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/infra/routes"
	"github.com/gin-gonic/gin"
)

func SetupDependencies(r *gin.Engine, db *sql.DB, jwtSecret string) {
	// Repositories
	repo := repository.NewPostgresAlertRepository(db)

	// Use Cases
	createUC := app.NewCreateAlertUseCase(repo)
	getUC := app.NewGetAlertUseCase(repo)
	listUC := app.NewListAlertsUseCase(repo)
	updateUC := app.NewUpdateAlertUseCase(repo)
	deleteUC := app.NewDeleteAlertUseCase(repo)
	changeStatusUC := app.NewChangeStatusAlertUseCase(repo)

	// Controllers
	createCtrl := controllers.NewCreateAlertController(createUC)
	getCtrl := controllers.NewGetAlertController(getUC)
	listCtrl := controllers.NewListAlertsController(listUC)
	updateCtrl := controllers.NewUpdateAlertController(updateUC)
	deleteCtrl := controllers.NewDeleteAlertController(deleteUC)
	changeStatusCtrl := controllers.NewChangeStatusAlertController(changeStatusUC)

	// Routes
	routes.RegisterAlertRoutes(
		r,
		createCtrl,
		getCtrl,
		listCtrl,
		updateCtrl,
		deleteCtrl,
		changeStatusCtrl,
		jwtSecret,
	)
}
