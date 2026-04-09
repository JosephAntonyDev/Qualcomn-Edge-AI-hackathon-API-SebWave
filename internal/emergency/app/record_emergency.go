package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/domain/repository"
)

type RecordEmergencyUseCase struct {
	repo repository.EmergencyRepository
}

func NewRecordEmergencyUseCase(repo repository.EmergencyRepository) *RecordEmergencyUseCase {
	return &RecordEmergencyUseCase{repo: repo}
}

func (uc *RecordEmergencyUseCase) Execute(ctx context.Context, em *entities.EmergencyEvent) error {
	return uc.repo.Record(ctx, em)
}
