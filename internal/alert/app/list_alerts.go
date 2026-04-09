package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/repository"
)

type ListAlertsUseCase struct {
	repo repository.AlertRepository
}

func NewListAlertsUseCase(repo repository.AlertRepository) *ListAlertsUseCase {
	return &ListAlertsUseCase{repo: repo}
}

func (uc *ListAlertsUseCase) Execute(ctx context.Context, filter repository.AlertFilter) ([]*entities.Alert, error) {
	if filter.Limit == 0 {
		filter.Limit = 50 // default limit
	}
	return uc.repo.List(ctx, filter)
}
