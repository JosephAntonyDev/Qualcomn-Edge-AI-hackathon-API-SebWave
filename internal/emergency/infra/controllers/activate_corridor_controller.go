package controllers

import (
	"net/http"
	"strconv"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/app"
	"github.com/gin-gonic/gin"
)

type ActivateCorridorRequest struct {
	ResponseTimeMs *int `json:"response_time_ms"`
}

type ActivateCorridorController struct {
	uc *app.ActivateCorridorUseCase
}

func NewActivateCorridorController(uc *app.ActivateCorridorUseCase) *ActivateCorridorController {
	return &ActivateCorridorController{uc: uc}
}

func (ctrl *ActivateCorridorController) Handle(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido."})
		return
	}

	var req ActivateCorridorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// es opcional, si falla el JSON ignóralo o responde error
		_ = err
	}

	if err := ctrl.uc.Execute(c.Request.Context(), id, req.ResponseTimeMs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al activar el corredor verde", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Corredor verde activado exitosamente"})
}
