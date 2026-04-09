package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/app"
	"github.com/gin-gonic/gin"
)

type DeleteSensorController struct {
	useCase *app.DeleteSensorUseCase
}

func NewDeleteSensorController(uc *app.DeleteSensorUseCase) *DeleteSensorController {
	return &DeleteSensorController{
		useCase: uc,
	}
}

func (ctrl *DeleteSensorController) Handle(c *gin.Context) {
	role, _ := c.Get("userRole")
	id := c.Param("id")

	err := ctrl.useCase.Execute(c.Request.Context(), id, role.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sensor eliminado"})
}
