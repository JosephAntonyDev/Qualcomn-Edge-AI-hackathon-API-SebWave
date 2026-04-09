package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/domain/entities"
	"github.com/gin-gonic/gin"
)

type SaveSystemMetricController struct {
	uc *app.SaveSystemMetricUseCase
}

func NewSaveSystemMetricController(uc *app.SaveSystemMetricUseCase) *SaveSystemMetricController {
	return &SaveSystemMetricController{uc: uc}
}

func (ctrl *SaveSystemMetricController) Handle(c *gin.Context) {
	var metric entities.SystemDailyMetric
	if err := c.ShouldBindJSON(&metric); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "details": err.Error()})
		return
	}

	if err := ctrl.uc.Execute(c.Request.Context(), &metric); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar métrica de sistema", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Métrica de sistema guardada exitosamente", "data": metric})
}
