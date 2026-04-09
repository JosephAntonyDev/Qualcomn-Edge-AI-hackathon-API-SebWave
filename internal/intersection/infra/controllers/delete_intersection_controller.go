package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/app"
	"github.com/gin-gonic/gin"
)

type DeleteIntersectionController struct {
	useCase *app.DeleteIntersectionUseCase
}

func NewDeleteIntersectionController(uc *app.DeleteIntersectionUseCase) *DeleteIntersectionController {
	return &DeleteIntersectionController{
		useCase: uc,
	}
}

func (ctrl *DeleteIntersectionController) Handle(c *gin.Context) {
	role, _ := c.Get("userRole")
	id := c.Param("id")

	err := ctrl.useCase.Execute(c.Request.Context(), id, role.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Intersección eliminada"})
}
