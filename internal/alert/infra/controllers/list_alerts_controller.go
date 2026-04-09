package controllers

import (
	"net/http"
	"strconv"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListAlertsController struct {
	uc *app.ListAlertsUseCase
}

func NewListAlertsController(uc *app.ListAlertsUseCase) *ListAlertsController {
	return &ListAlertsController{uc: uc}
}

func (ctrl *ListAlertsController) Handle(c *gin.Context) {
	var filter repository.AlertFilter

	if intersectionID := c.Query("intersection_id"); intersectionID != "" {
		if id, err := uuid.Parse(intersectionID); err == nil {
			filter.IntersectionID = &id
		}
	}

	if alertType := c.Query("type"); alertType != "" {
		t := entities.AlertType(alertType)
		filter.Type = &t
	}

	if severity := c.Query("severity"); severity != "" {
		s := entities.AlertSeverity(severity)
		filter.Severity = &s
	}

	if status := c.Query("status"); status != "" {
		st := entities.AlertStatus(status)
		filter.Status = &st
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

	alerts, err := ctrl.uc.Execute(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener alertas", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": alerts})
}
