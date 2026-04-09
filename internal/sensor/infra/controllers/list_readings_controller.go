package controllers

import (
	"net/http"
	"strconv"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/app"
	"github.com/gin-gonic/gin"
)

type ListReadingsController struct {
	useCase *app.ListReadingsUseCase
}

func NewListReadingsController(uc *app.ListReadingsUseCase) *ListReadingsController {
	return &ListReadingsController{useCase: uc}
}

func (ctrl *ListReadingsController) Handle(c *gin.Context) {
	id := c.Param("id")
	limitStr := c.Query("limit")
	limit := 50
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	res, err := ctrl.useCase.Execute(c.Request.Context(), id, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
