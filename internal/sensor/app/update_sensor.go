package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/domain/repository"
	userEntities "github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
	"github.com/google/uuid"
)

type UpdateSensorRequest struct {
	LaneDirection  *string `json:"lane_direction,omitempty"`
	ConnectionType *string `json:"connection_type,omitempty"`
	PinAssignment  *string `json:"pin_assignment,omitempty"`
	IsActive       *bool   `json:"is_active,omitempty"`
}

type UpdateSensorUseCase struct {
	repo repository.SensorRepository
}

func NewUpdateSensorUseCase(r repository.SensorRepository) *UpdateSensorUseCase {
	return &UpdateSensorUseCase{repo: r}
}

func (uc *UpdateSensorUseCase) Execute(ctx context.Context, id string, requesterRole string, req UpdateSensorRequest) (*entities.Sensor, error) {
	if requesterRole != string(userEntities.RoleAdmin) && requesterRole != string(userEntities.RoleOperator) {
		return nil, errors.New("operación denegada")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("id de sensor inválido")
	}

	sensor, err := uc.repo.GetSensor(ctx, uid)
	if err != nil {
		return nil, err
	}

	if req.LaneDirection != nil {
		sensor.LaneDirection = req.LaneDirection
	}
	if req.ConnectionType != nil {
		sensor.ConnectionType = req.ConnectionType
	}
	if req.PinAssignment != nil {
		sensor.PinAssignment = req.PinAssignment
	}
	if req.IsActive != nil {
		sensor.IsActive = *req.IsActive
	}

	if err := uc.repo.UpdateSensor(ctx, sensor); err != nil {
		return nil, err
	}

	return sensor, nil
}
