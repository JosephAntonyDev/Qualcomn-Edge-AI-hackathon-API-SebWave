package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/app"
	"github.com/gin-gonic/gin"
)

type CreateIntersectionController struct {
	useCase *app.CreateIntersectionUseCase
}

func NewCreateIntersectionController(uc *app.CreateIntersectionUseCase) *CreateIntersectionController {
	return &CreateIntersectionController{
		useCase: uc,
	}
}

func (ctrl *CreateIntersectionController) Handle(c *gin.Context) {
	role, _ := c.Get("userRole")
	id, _ := c.Get("userID")

	var req app.CreateIntersectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	res, err := ctrl.useCase.Execute(c.Request.Context(), role.(string), id.(string), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": res})
}
