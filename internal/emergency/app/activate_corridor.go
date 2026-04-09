package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/domain/repository"
)

type ActivateCorridorUseCase struct {
	repo repository.EmergencyRepository
}

func NewActivateCorridorUseCase(repo repository.EmergencyRepository) *ActivateCorridorUseCase {
	return &ActivateCorridorUseCase{repo: repo}
}

func (uc *ActivateCorridorUseCase) Execute(ctx context.Context, id int64, responseTimeMs *int) error {
	return uc.repo.UpdateCorridorStatus(ctx, id, true, responseTimeMs)
}
