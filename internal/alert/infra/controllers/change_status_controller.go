package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ChangeStatusRequest struct {
	Status         entities.AlertStatus `json:"status" binding:"required"`
	ResolvedBy     *uuid.UUID           `json:"resolved_by"`
	ResolutionNote *string              `json:"resolution_note"`
}

type ChangeStatusAlertController struct {
	uc *app.ChangeStatusAlertUseCase
}

func NewChangeStatusAlertController(uc *app.ChangeStatusAlertUseCase) *ChangeStatusAlertController {
	return &ChangeStatusAlertController{uc: uc}
}

func (ctrl *ChangeStatusAlertController) Handle(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req ChangeStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "details": err.Error()})
		return
	}

	if err := ctrl.uc.Execute(c.Request.Context(), id, req.Status, req.ResolvedBy, req.ResolutionNote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cambiar estado de la alerta", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estado de alerta serializado exitosamente"})
}
