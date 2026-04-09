package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/app"
	"github.com/gin-gonic/gin"
)

type GetUserByUsernameController struct {
	useCase *app.GetUserByUsernameUseCase
}

func NewGetUserByUsernameController(uc *app.GetUserByUsernameUseCase) *GetUserByUsernameController {
	return &GetUserByUsernameController{
		useCase: uc,
	}
}

func (ctrl *GetUserByUsernameController) Handle(c *gin.Context) {
	username := c.Param("username")

	user, err := ctrl.useCase.Execute(c.Request.Context(), username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
