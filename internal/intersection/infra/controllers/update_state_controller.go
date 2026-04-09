package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/app"
	"github.com/gin-gonic/gin"
)

type UpdateStateController struct {
	useCase *app.UpdateStateUseCase
}

func NewUpdateStateController(uc *app.UpdateStateUseCase) *UpdateStateController {
	return &UpdateStateController{
		useCase: uc,
	}
}

func (ctrl *UpdateStateController) Handle(c *gin.Context) {
	role, _ := c.Get("userRole")
	id := c.Param("id")

	var req app.UpdateStateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err := ctrl.useCase.Execute(c.Request.Context(), id, role.(string), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estado actualizado"})
}
