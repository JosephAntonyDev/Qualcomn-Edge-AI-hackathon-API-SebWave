package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/app"
	"github.com/gin-gonic/gin"
)

type GetUserByIDController struct {
	useCase *app.GetUserByIDUseCase
}

func NewGetUserByIDController(uc *app.GetUserByIDUseCase) *GetUserByIDController {
	return &GetUserByIDController{
		useCase: uc,
	}
}

func (ctrl *GetUserByIDController) Handle(c *gin.Context) {
	id := c.Param("id")

	user, err := ctrl.useCase.Execute(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "id de usuario inválido" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
