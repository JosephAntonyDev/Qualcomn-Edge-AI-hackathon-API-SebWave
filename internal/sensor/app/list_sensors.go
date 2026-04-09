package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/domain/repository"
	"github.com/google/uuid"
)

type ListSensorsUseCase struct {
	repo repository.SensorRepository
}

func NewListSensorsUseCase(r repository.SensorRepository) *ListSensorsUseCase {
	return &ListSensorsUseCase{repo: r}
}

func (uc *ListSensorsUseCase) Execute(ctx context.Context, intersectionID string) ([]*entities.Sensor, error) {
	uid, err := uuid.Parse(intersectionID)
	if err != nil {
		return nil, errors.New("id de intersección inválido")
	}

	return uc.repo.ListSensorsByIntersection(ctx, uid)
}
