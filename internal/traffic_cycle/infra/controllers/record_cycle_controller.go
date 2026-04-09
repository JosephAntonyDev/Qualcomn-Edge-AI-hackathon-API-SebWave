package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/app"
	"github.com/gin-gonic/gin"
)

type RecordCycleController struct {
	useCase *app.RecordCycleUseCase
}

func NewRecordCycleController(uc *app.RecordCycleUseCase) *RecordCycleController {
	return &RecordCycleController{useCase: uc}
}

func (ctrl *RecordCycleController) Handle(c *gin.Context) {
	// Role might be empty if called from an IoT specific route that uses an API Key instead of JWT
	role, exists := c.Get("userRole")
	roleStr := ""
	if exists {
		roleStr = role.(string)
	}

	var req app.RecordCycleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "details": err.Error()})
		return
	}

	res, err := ctrl.useCase.Execute(c.Request.Context(), roleStr, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": res})
}
