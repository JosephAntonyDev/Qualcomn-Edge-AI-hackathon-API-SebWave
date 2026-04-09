package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/domain/repository"
	"github.com/gin-gonic/gin"
)

type ListSystemMetricsController struct {
	uc *app.ListSystemMetricsUseCase
}

func NewListSystemMetricsController(uc *app.ListSystemMetricsUseCase) *ListSystemMetricsController {
	return &ListSystemMetricsController{uc: uc}
}

func (ctrl *ListSystemMetricsController) Handle(c *gin.Context) {
	var filter repository.MetricsFilter

	if startStr := c.Query("start_date"); startStr != "" {
		if d, err := time.Parse("2006-01-02", startStr); err == nil {
			filter.StartDate = &d
		}
	}

	if endStr := c.Query("end_date"); endStr != "" {
		if d, err := time.Parse("2006-01-02", endStr); err == nil {
			filter.EndDate = &d
		}
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			filter.Limit = limit
		}
	}

	if offsetStr := c.Query("offset"); offsetStr != "" {
		if offset, err := strconv.Atoi(offsetStr); err == nil {
			filter.Offset = offset
		}
	}

	metrics, err := ctrl.uc.Execute(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener métricas del sistema", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": metrics})
}
