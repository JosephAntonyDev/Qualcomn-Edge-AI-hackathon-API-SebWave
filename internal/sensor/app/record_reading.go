package app

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/domain/repository"
	"github.com/google/uuid"
)

type RecordReadingRequest struct {
	SensorID       string          `json:"sensor_id" binding:"required"`
	IntersectionID string          `json:"intersection_id" binding:"required"`
	IsOccupied     *bool           `json:"is_occupied,omitempty"`
	DistanceMM     *int            `json:"distance_mm,omitempty"`
	EnergyPct      *int            `json:"energy_pct,omitempty"`
	Confidence     *float64        `json:"confidence,omitempty"`
	RawData        json.RawMessage `json:"raw_data,omitempty"`
}

type RecordReadingUseCase struct {
	repo repository.SensorRepository
}

func NewRecordReadingUseCase(r repository.SensorRepository) *RecordReadingUseCase {
	return &RecordReadingUseCase{repo: r}
}

func (uc *RecordReadingUseCase) Execute(ctx context.Context, req RecordReadingRequest) error {
	senID, err := uuid.Parse(req.SensorID)
	if err != nil {
		return errors.New("id de sensor inválido")
	}

	intID, err := uuid.Parse(req.IntersectionID)
	if err != nil {
		return errors.New("id de intersección inválido")
	}

	reading := &entities.SensorReading{
		SensorID:       senID,
		IntersectionID: intID,
		IsOccupied:     req.IsOccupied,
		DistanceMM:     req.DistanceMM,
		EnergyPct:      req.EnergyPct,
		Confidence:     req.Confidence,
		RawData:        req.RawData,
	}

	return uc.repo.RecordReading(ctx, reading)
}
