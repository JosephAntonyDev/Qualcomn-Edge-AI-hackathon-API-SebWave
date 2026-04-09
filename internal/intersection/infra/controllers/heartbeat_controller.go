package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/app"
	"github.com/gin-gonic/gin"
)

type HeartbeatController struct {
	useCase *app.HeartbeatUseCase
}

func NewHeartbeatController(uc *app.HeartbeatUseCase) *HeartbeatController {
	return &HeartbeatController{
		useCase: uc,
	}
}

func (ctrl *HeartbeatController) Handle(c *gin.Context) {
	var req app.HeartbeatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err := ctrl.useCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Heartbeat registrado"})
}
