package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/domain/repository"
)

type ListEmergenciesUseCase struct {
	repo repository.EmergencyRepository
}

func NewListEmergenciesUseCase(repo repository.EmergencyRepository) *ListEmergenciesUseCase {
	return &ListEmergenciesUseCase{repo: repo}
}

func (uc *ListEmergenciesUseCase) Execute(ctx context.Context, filter repository.EmergencyFilter) ([]*entities.EmergencyEvent, error) {
	if filter.Limit == 0 {
		filter.Limit = 50 // default
	}
	return uc.repo.List(ctx, filter)
}
