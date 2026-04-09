package infra

import (
	"database/sql"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/infra/controllers"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/infra/repository"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/infra/routes"
	"github.com/gin-gonic/gin"
)

func SetupDependencies(r *gin.Engine, db *sql.DB, jwtSecret string) {
	repo := repository.NewPostgresIntersectionRepository(db)

	createUC := app.NewCreateIntersectionUseCase(repo)
	getUC := app.NewGetIntersectionUseCase(repo)
	listUC := app.NewListIntersectionsUseCase(repo)
	updateUC := app.NewUpdateIntersectionUseCase(repo)
	deleteUC := app.NewDeleteIntersectionUseCase(repo)
	stateUC := app.NewUpdateStateUseCase(repo)
	hbUC := app.NewHeartbeatUseCase(repo)

	createCtrl := controllers.NewCreateIntersectionController(createUC)
	getCtrl := controllers.NewGetIntersectionController(getUC)
	listCtrl := controllers.NewListIntersectionsController(listUC)
	updateCtrl := controllers.NewUpdateIntersectionController(updateUC)
	deleteCtrl := controllers.NewDeleteIntersectionController(deleteUC)
	stateCtrl := controllers.NewUpdateStateController(stateUC)
	hbCtrl := controllers.NewHeartbeatController(hbUC)

	routes.SetupIntersectionRoutes(
		r,
		createCtrl,
		getCtrl,
		listCtrl,
		updateCtrl,
		deleteCtrl,
		stateCtrl,
		hbCtrl,
		jwtSecret,
	)
}
