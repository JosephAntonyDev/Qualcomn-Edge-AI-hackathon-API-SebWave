package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/domain/repository"
	userEntities "github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
	"github.com/google/uuid"
)

type DeleteSensorUseCase struct {
	repo repository.SensorRepository
}

func NewDeleteSensorUseCase(r repository.SensorRepository) *DeleteSensorUseCase {
	return &DeleteSensorUseCase{repo: r}
}

func (uc *DeleteSensorUseCase) Execute(ctx context.Context, id string, requesterRole string) error {
	if requesterRole != string(userEntities.RoleAdmin) {
		return errors.New("operación denegada")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return errors.New("id de sensor inválido")
	}

	return uc.repo.DeleteSensor(ctx, uid)
}
