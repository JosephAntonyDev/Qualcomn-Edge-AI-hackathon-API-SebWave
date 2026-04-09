package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/domain/repository"
	userEntities "github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
	"github.com/google/uuid"
)

type CreateIntersectionRequest struct {
	SerialNumber     string  `json:"serial_number" binding:"required"`
	Name             string  `json:"name" binding:"required"`
	Description      *string `json:"description,omitempty"`
	Latitude         float64 `json:"latitude" binding:"required"`
	Longitude        float64 `json:"longitude" binding:"required"`
	NodeID           *string `json:"node_id,omitempty"`
	MaxCongestionPct float64 `json:"max_congestion_pct"`
	MinGreenTimeS    int     `json:"min_green_time_s"`
	MaxGreenTimeS    int     `json:"max_green_time_s"`
	DefaultGreenS    int     `json:"default_green_s"`
	DefaultRedS      int     `json:"default_red_s"`
	YellowTimeS      int     `json:"yellow_time_s"`
}

type CreateIntersectionUseCase struct {
	repo repository.IntersectionRepository
}

func NewCreateIntersectionUseCase(r repository.IntersectionRepository) *CreateIntersectionUseCase {
	return &CreateIntersectionUseCase{repo: r}
}

func (uc *CreateIntersectionUseCase) Execute(ctx context.Context, requesterRole string, requesterID string, req CreateIntersectionRequest) (*entities.Intersection, error) {
	if requesterRole != string(userEntities.RoleAdmin) && requesterRole != string(userEntities.RoleOperator) {
		return nil, errors.New("operación denegada")
	}

	uid, err := uuid.Parse(requesterID)
	if err != nil {
		return nil, errors.New("id de creador inválido")
	}

	val, _ := uc.repo.GetBySerialNumber(ctx, req.SerialNumber)
	if val != nil {
		return nil, errors.New("una intersección con este número de serie ya existe")
	}

	if req.MaxCongestionPct == 0 {
		req.MaxCongestionPct = 80.0
	}
	if req.MinGreenTimeS == 0 {
		req.MinGreenTimeS = 8
	}
	if req.MaxGreenTimeS == 0 {
		req.MaxGreenTimeS = 90
	}
	if req.DefaultGreenS == 0 {
		req.DefaultGreenS = 15
	}
	if req.DefaultRedS == 0 {
		req.DefaultRedS = 15
	}
	if req.YellowTimeS == 0 {
		req.YellowTimeS = 3
	}

	is := &entities.Intersection{
		SerialNumber:     req.SerialNumber,
		Name:             req.Name,
		Description:      req.Description,
		Latitude:         req.Latitude,
		Longitude:        req.Longitude,
		NodeID:           req.NodeID,
		MaxCongestionPct: req.MaxCongestionPct,
		MinGreenTimeS:    req.MinGreenTimeS,
		MaxGreenTimeS:    req.MaxGreenTimeS,
		DefaultGreenS:    req.DefaultGreenS,
		DefaultRedS:      req.DefaultRedS,
		YellowTimeS:      req.YellowTimeS,
		Status:           entities.StatusDisconnected,
		OperationMode:    entities.ModeFixed,
		CreatedBy:        &uid,
	}

	if err := uc.repo.Create(ctx, is); err != nil {
		return nil, err
	}

	return is, nil
}
