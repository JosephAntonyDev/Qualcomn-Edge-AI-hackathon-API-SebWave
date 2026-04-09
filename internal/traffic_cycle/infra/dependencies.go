package infra

import (
	"database/sql"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/infra/controllers"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/infra/repository"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/infra/routes"
	"github.com/gin-gonic/gin"
)

func SetupDependencies(r *gin.Engine, db *sql.DB, jwtSecret string) {
	repo := repository.NewPostgresTrafficCycleRepository(db)

	recordUC := app.NewRecordCycleUseCase(repo)
	getUC := app.NewGetCycleUseCase(repo)
	listUC := app.NewListCyclesUseCase(repo)
	deleteUC := app.NewDeleteCycleUseCase(repo)

	recordCtrl := controllers.NewRecordCycleController(recordUC)
	getCtrl := controllers.NewGetCycleController(getUC)
	listCtrl := controllers.NewListCyclesController(listUC)
	deleteCtrl := controllers.NewDeleteCycleController(deleteUC)

	routes.SetupTrafficCycleRoutes(
		r,
		recordCtrl,
		getCtrl,
		listCtrl,
		deleteCtrl,
		jwtSecret,
	)
}
