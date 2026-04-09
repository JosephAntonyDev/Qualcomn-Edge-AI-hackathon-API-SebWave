package controllers

import (
	"net/http"
	"strconv"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/app"
	"github.com/gin-gonic/gin"
)

type GetEmergencyController struct {
	uc *app.GetEmergencyUseCase
}

func NewGetEmergencyController(uc *app.GetEmergencyUseCase) *GetEmergencyController {
	return &GetEmergencyController{uc: uc}
}

func (ctrl *GetEmergencyController) Handle(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido. Debe ser entero."})
		return
	}

	em, err := ctrl.uc.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Emergencia no encontrada", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": em})
}
