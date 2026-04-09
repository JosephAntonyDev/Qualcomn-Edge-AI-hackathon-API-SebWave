package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/domain/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListEmergenciesController struct {
	uc *app.ListEmergenciesUseCase
}

func NewListEmergenciesController(uc *app.ListEmergenciesUseCase) *ListEmergenciesController {
	return &ListEmergenciesController{uc: uc}
}

func (ctrl *ListEmergenciesController) Handle(c *gin.Context) {
	var filter repository.EmergencyFilter

	if intersectionID := c.Query("intersection_id"); intersectionID != "" {
		if id, err := uuid.Parse(intersectionID); err == nil {
			filter.IntersectionID = &id
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

	if methodStr := c.Query("detection_method"); methodStr != "" {
		m := entities.DetectionMethod(methodStr)
		filter.DetectionMethod = &m
	}

	if corrStr := c.Query("corridor_activated"); corrStr != "" {
		corr := strings.ToLower(corrStr) == "true"
		filter.CorridorActivated = &corr
	}

	emergencies, err := ctrl.uc.Execute(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al listar las emergencias", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": emergencies})
}
