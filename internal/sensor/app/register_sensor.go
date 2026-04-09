package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/domain/repository"
	userEntities "github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
	"github.com/google/uuid"
)

type RegisterSensorRequest struct {
	IntersectionID string  `json:"intersection_id" binding:"required"`
	SensorType     string  `json:"sensor_type" binding:"required"`
	LaneDirection  *string `json:"lane_direction,omitempty"`
	ConnectionType *string `json:"connection_type,omitempty"`
	PinAssignment  *string `json:"pin_assignment,omitempty"`
}

type RegisterSensorUseCase struct {
	repo repository.SensorRepository
}

func NewRegisterSensorUseCase(r repository.SensorRepository) *RegisterSensorUseCase {
	return &RegisterSensorUseCase{repo: r}
}

func (uc *RegisterSensorUseCase) Execute(ctx context.Context, requesterRole string, req RegisterSensorRequest) (*entities.Sensor, error) {
	if requesterRole != string(userEntities.RoleAdmin) && requesterRole != string(userEntities.RoleOperator) {
		return nil, errors.New("operación denegada")
	}

	intID, err := uuid.Parse(req.IntersectionID)
	if err != nil {
		return nil, errors.New("id de intersección inválido")
	}

	sensor := &entities.Sensor{
		IntersectionID: intID,
		SensorType:     entities.SensorType(req.SensorType),
		LaneDirection:  req.LaneDirection,
		ConnectionType: req.ConnectionType,
		PinAssignment:  req.PinAssignment,
		IsActive:       true,
		FailureCount:   0,
	}

	if err := uc.repo.RegisterSensor(ctx, sensor); err != nil {
		return nil, err
	}

	return sensor, nil
}
