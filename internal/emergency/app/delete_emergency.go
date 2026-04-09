package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/domain/repository"
)

type DeleteEmergencyUseCase struct {
	repo repository.EmergencyRepository
}

func NewDeleteEmergencyUseCase(repo repository.EmergencyRepository) *DeleteEmergencyUseCase {
	return &DeleteEmergencyUseCase{repo: repo}
}

func (uc *DeleteEmergencyUseCase) Execute(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
