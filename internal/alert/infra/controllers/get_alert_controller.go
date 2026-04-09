package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/app"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetAlertController struct {
	uc *app.GetAlertUseCase
}

func NewGetAlertController(uc *app.GetAlertUseCase) *GetAlertController {
	return &GetAlertController{uc: uc}
}

func (ctrl *GetAlertController) Handle(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	alert, err := ctrl.uc.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alerta no encontrada", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": alert})
}
