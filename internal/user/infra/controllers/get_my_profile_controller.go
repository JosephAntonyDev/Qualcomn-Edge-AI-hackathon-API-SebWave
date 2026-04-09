package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/app"
	"github.com/gin-gonic/gin"
)

type GetMyProfileController struct {
	useCase *app.GetMyProfileUseCase
}

func NewGetMyProfileController(uc *app.GetMyProfileUseCase) *GetMyProfileController {
	return &GetMyProfileController{
		useCase: uc,
	}
}

func (ctrl *GetMyProfileController) Handle(c *gin.Context) {
	requesterIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se pudo identificar tu usuario"})
		return
	}
	requesterID := requesterIDVal.(string)

	user, err := ctrl.useCase.Execute(c.Request.Context(), requesterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
