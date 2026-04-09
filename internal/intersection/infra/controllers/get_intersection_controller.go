package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/app"
	"github.com/gin-gonic/gin"
)

type GetIntersectionController struct {
	useCase *app.GetIntersectionUseCase
}

func NewGetIntersectionController(uc *app.GetIntersectionUseCase) *GetIntersectionController {
	return &GetIntersectionController{
		useCase: uc,
	}
}

func (ctrl *GetIntersectionController) Handle(c *gin.Context) {
	id := c.Param("id")

	res, err := ctrl.useCase.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
