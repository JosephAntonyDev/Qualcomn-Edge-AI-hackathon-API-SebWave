package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/domain/entities"
	"github.com/gin-gonic/gin"
)

type RecordEmergencyController struct {
	uc *app.RecordEmergencyUseCase
}

func NewRecordEmergencyController(uc *app.RecordEmergencyUseCase) *RecordEmergencyController {
	return &RecordEmergencyController{uc: uc}
}

func (ctrl *RecordEmergencyController) Handle(c *gin.Context) {
	var em entities.EmergencyEvent
	if err := c.ShouldBindJSON(&em); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "details": err.Error()})
		return
	}

	if err := ctrl.uc.Execute(c.Request.Context(), &em); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar emergencia", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Emergencia registrada exitosamente", "data": em})
}
