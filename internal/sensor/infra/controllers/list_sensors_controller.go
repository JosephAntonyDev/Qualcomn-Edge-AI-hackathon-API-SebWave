package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/app"
	"github.com/gin-gonic/gin"
)

type ListSensorsController struct {
	useCase *app.ListSensorsUseCase
}

func NewListSensorsController(uc *app.ListSensorsUseCase) *ListSensorsController {
	return &ListSensorsController{
		useCase: uc,
	}
}

func (ctrl *ListSensorsController) Handle(c *gin.Context) {
	intersectionID := c.Query("intersection_id")
	if intersectionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "intersection_id requerido"})
		return
	}

	res, err := ctrl.useCase.Execute(c.Request.Context(), intersectionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
