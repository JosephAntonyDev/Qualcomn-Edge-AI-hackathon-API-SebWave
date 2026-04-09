package infra

import (
	"database/sql"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/infra/controllers"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/infra/repository"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/infra/routes"
	"github.com/gin-gonic/gin"
)

func SetupDependencies(r *gin.Engine, db *sql.DB, jwtSecret string) {
	repo := repository.NewPostgresSensorRepository(db)

	registerUC := app.NewRegisterSensorUseCase(repo)
	listUC := app.NewListSensorsUseCase(repo)
	updateUC := app.NewUpdateSensorUseCase(repo)
	deleteUC := app.NewDeleteSensorUseCase(repo)
	recordReadingUC := app.NewRecordReadingUseCase(repo)
	listReadingsUC := app.NewListReadingsUseCase(repo)

	registerCtrl := controllers.NewRegisterSensorController(registerUC)
	listCtrl := controllers.NewListSensorsController(listUC)
	updateCtrl := controllers.NewUpdateSensorController(updateUC)
	deleteCtrl := controllers.NewDeleteSensorController(deleteUC)
	recordReadingCtrl := controllers.NewRecordReadingController(recordReadingUC)
	listReadingsCtrl := controllers.NewListReadingsController(listReadingsUC)

	routes.SetupSensorRoutes(
		r,
		registerCtrl,
		listCtrl,
		updateCtrl,
		deleteCtrl,
		recordReadingCtrl,
		listReadingsCtrl,
		jwtSecret,
	)
}
