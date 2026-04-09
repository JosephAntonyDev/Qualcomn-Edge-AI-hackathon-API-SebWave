package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/app"
	"github.com/gin-gonic/gin"
)

type RegisterSensorController struct {
	useCase *app.RegisterSensorUseCase
}

func NewRegisterSensorController(uc *app.RegisterSensorUseCase) *RegisterSensorController {
	return &RegisterSensorController{
		useCase: uc,
	}
}

func (ctrl *RegisterSensorController) Handle(c *gin.Context) {
	role, _ := c.Get("userRole")

	var req app.RegisterSensorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	res, err := ctrl.useCase.Execute(c.Request.Context(), role.(string), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": res})
}
