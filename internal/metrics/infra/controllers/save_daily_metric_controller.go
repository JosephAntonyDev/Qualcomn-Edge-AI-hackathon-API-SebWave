package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/domain/entities"
	"github.com/gin-gonic/gin"
)

type SaveDailyMetricController struct {
	uc *app.SaveDailyMetricUseCase
}

func NewSaveDailyMetricController(uc *app.SaveDailyMetricUseCase) *SaveDailyMetricController {
	return &SaveDailyMetricController{uc: uc}
}

func (ctrl *SaveDailyMetricController) Handle(c *gin.Context) {
	var metric entities.DailyMetric
	if err := c.ShouldBindJSON(&metric); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "details": err.Error()})
		return
	}

	if err := ctrl.uc.Execute(c.Request.Context(), &metric); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar métrica diaria", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Métrica diaria guardada exitosamente", "data": metric})
}
