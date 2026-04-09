package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/app"
	"github.com/gin-gonic/gin"
)

type UpdateSensorController struct {
	useCase *app.UpdateSensorUseCase
}

func NewUpdateSensorController(uc *app.UpdateSensorUseCase) *UpdateSensorController {
	return &UpdateSensorController{
		useCase: uc,
	}
}

func (ctrl *UpdateSensorController) Handle(c *gin.Context) {
	role, _ := c.Get("userRole")
	id := c.Param("id")

	var req app.UpdateSensorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	res, err := ctrl.useCase.Execute(c.Request.Context(), id, role.(string), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
