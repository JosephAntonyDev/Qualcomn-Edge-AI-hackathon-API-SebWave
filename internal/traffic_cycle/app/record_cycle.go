package app

import (
	"context"
	"errors"
	"time"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/domain/repository"
	userEntities "github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
	"github.com/google/uuid"
)

type RecordCycleRequest struct {
	IntersectionID     string     `json:"intersection_id" binding:"required"`
	OperationMode      string     `json:"operation_mode" binding:"required"`
	GreenNSMs          int        `json:"green_ns_ms" binding:"required"`
	GreenEOMs          int        `json:"green_eo_ms" binding:"required"`
	YellowMs           int        `json:"yellow_ms"`
	RadarOccupied      *bool      `json:"radar_occupied,omitempty"`
	RadarDistanceCm    *int       `json:"radar_distance_cm,omitempty"`
	RadarEnergy        *int       `json:"radar_energy,omitempty"`
	ModulinoOccupied   *bool      `json:"modulino_occupied,omitempty"`
	ModulinoDistanceMm *int       `json:"modulino_distance_mm,omitempty"`
	CameraNSOccupied   *bool      `json:"camera_ns_occupied,omitempty"`
	CameraEOOccupied   *bool      `json:"camera_eo_occupied,omitempty"`
	WaitTimeSavedMs    *int       `json:"wait_time_saved_ms,omitempty"`
	CO2SavedG          *float64   `json:"co2_saved_g,omitempty"`
	ModelConfidence    *float64   `json:"model_confidence,omitempty"`
	StartedAt          *time.Time `json:"started_at,omitempty"`
	EndedAt            *time.Time `json:"ended_at,omitempty"`
}

type RecordCycleUseCase struct {
	repo repository.TrafficCycleRepository
}

func NewRecordCycleUseCase(r repository.TrafficCycleRepository) *RecordCycleUseCase {
	return &RecordCycleUseCase{repo: r}
}

func (uc *RecordCycleUseCase) Execute(ctx context.Context, requesterRole string, req RecordCycleRequest) (*entities.TrafficCycle, error) {
	// IoT devices (que no manejan el JWT formal pero comparten credenciales básicas) o Admin/Operator
	if requesterRole != "" && requesterRole != string(userEntities.RoleAdmin) && requesterRole != string(userEntities.RoleOperator) {
		return nil, errors.New("operación denegada")
	}

	intID, err := uuid.Parse(req.IntersectionID)
	if err != nil {
		return nil, errors.New("id de intersección inválido")
	}

	yellow := req.YellowMs
	if yellow <= 0 {
		yellow = 3000 // default 3s
	}

	start := time.Now()
	if req.StartedAt != nil {
		start = *req.StartedAt
	}

	cycle := &entities.TrafficCycle{
		IntersectionID:     intID,
		OperationMode:      req.OperationMode,
		GreenNSMs:          req.GreenNSMs,
		GreenEOMs:          req.GreenEOMs,
		YellowMs:           yellow,
		RadarOccupied:      req.RadarOccupied,
		RadarDistanceCm:    req.RadarDistanceCm,
		RadarEnergy:        req.RadarEnergy,
		ModulinoOccupied:   req.ModulinoOccupied,
		ModulinoDistanceMm: req.ModulinoDistanceMm,
		CameraNSOccupied:   req.CameraNSOccupied,
		CameraEOOccupied:   req.CameraEOOccupied,
		WaitTimeSavedMs:    req.WaitTimeSavedMs,
		CO2SavedG:          req.CO2SavedG,
		ModelConfidence:    req.ModelConfidence,
		StartedAt:          start,
		EndedAt:            req.EndedAt,
	}

	if err := uc.repo.RecordCycle(ctx, cycle); err != nil {
		return nil, err
	}

	return cycle, nil
}
