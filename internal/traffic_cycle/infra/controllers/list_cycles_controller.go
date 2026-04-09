package controllers

import (
	"net/http"
	"strconv"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/app"
	"github.com/gin-gonic/gin"
)

type ListCyclesController struct {
	useCase *app.ListCyclesUseCase
}

func NewListCyclesController(uc *app.ListCyclesUseCase) *ListCyclesController {
	return &ListCyclesController{useCase: uc}
}

func (ctrl *ListCyclesController) Handle(c *gin.Context) {
	intersectionID := c.Query("intersection_id")
	if intersectionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El parámetro intersection_id es requerido"})
		return
	}

	limit := 20
	if l, err := strconv.Atoi(c.Query("limit")); err == nil {
		limit = l
	}

	offset := 0
	if o, err := strconv.Atoi(c.Query("offset")); err == nil {
		offset = o
	}

	res, err := ctrl.useCase.Execute(c.Request.Context(), intersectionID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
