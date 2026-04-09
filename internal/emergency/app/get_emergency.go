package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/domain/repository"
)

type GetEmergencyUseCase struct {
	repo repository.EmergencyRepository
}

func NewGetEmergencyUseCase(repo repository.EmergencyRepository) *GetEmergencyUseCase {
	return &GetEmergencyUseCase{repo: repo}
}

func (uc *GetEmergencyUseCase) Execute(ctx context.Context, id int64) (*entities.EmergencyEvent, error) {
	return uc.repo.GetByID(ctx, id)
}
