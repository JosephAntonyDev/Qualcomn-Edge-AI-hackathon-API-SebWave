package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/app"
	"github.com/gin-gonic/gin"
)

type ListIntersectionsController struct {
	useCase *app.ListIntersectionsUseCase
}

func NewListIntersectionsController(uc *app.ListIntersectionsUseCase) *ListIntersectionsController {
	return &ListIntersectionsController{
		useCase: uc,
	}
}

func (ctrl *ListIntersectionsController) Handle(c *gin.Context) {
	res, err := ctrl.useCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
