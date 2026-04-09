package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/domain/repository"
	"github.com/google/uuid"
)

type ListReadingsUseCase struct {
	repo repository.SensorRepository
}

func NewListReadingsUseCase(r repository.SensorRepository) *ListReadingsUseCase {
	return &ListReadingsUseCase{repo: r}
}

func (uc *ListReadingsUseCase) Execute(ctx context.Context, sensorID string, limit int) ([]*entities.SensorReading, error) {
	uid, err := uuid.Parse(sensorID)
	if err != nil {
		return nil, errors.New("id de sensor inválido")
	}

	if limit <= 0 || limit > 1000 {
		limit = 50
	}

	return uc.repo.GetReadings(ctx, uid, limit)
}
