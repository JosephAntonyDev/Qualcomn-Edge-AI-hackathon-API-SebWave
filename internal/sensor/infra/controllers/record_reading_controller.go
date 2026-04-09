package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/app"
	"github.com/gin-gonic/gin"
)

type RecordReadingController struct {
	useCase *app.RecordReadingUseCase
}

func NewRecordReadingController(uc *app.RecordReadingUseCase) *RecordReadingController {
	return &RecordReadingController{useCase: uc}
}

func (ctrl *RecordReadingController) Handle(c *gin.Context) {
	var req app.RecordReadingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if req.SensorID == "" {
		req.SensorID = c.Param("id")
	}

	err := ctrl.useCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Lectura registrada"})
}
