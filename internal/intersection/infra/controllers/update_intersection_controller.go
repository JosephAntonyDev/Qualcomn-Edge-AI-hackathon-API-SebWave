package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/app"
	"github.com/gin-gonic/gin"
)

type UpdateIntersectionController struct {
	useCase *app.UpdateIntersectionUseCase
}

func NewUpdateIntersectionController(uc *app.UpdateIntersectionUseCase) *UpdateIntersectionController {
	return &UpdateIntersectionController{
		useCase: uc,
	}
}

func (ctrl *UpdateIntersectionController) Handle(c *gin.Context) {
	role, _ := c.Get("userRole")
	id := c.Param("id")

	var req app.UpdateIntersectionRequest
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
