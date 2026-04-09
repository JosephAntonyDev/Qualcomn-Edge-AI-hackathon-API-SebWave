package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/domain/repository"
	userEntities "github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
	"github.com/google/uuid"
)

type UpdateIntersectionRequest struct {
	Name             *string  `json:"name,omitempty"`
	Description      *string  `json:"description,omitempty"`
	Latitude         *float64 `json:"latitude,omitempty"`
	Longitude        *float64 `json:"longitude,omitempty"`
	NodeID           *string  `json:"node_id,omitempty"`
	MaxCongestionPct *float64 `json:"max_congestion_pct,omitempty"`
	MinGreenTimeS    *int     `json:"min_green_time_s,omitempty"`
	MaxGreenTimeS    *int     `json:"max_green_time_s,omitempty"`
	DefaultGreenS    *int     `json:"default_green_s,omitempty"`
	DefaultRedS      *int     `json:"default_red_s,omitempty"`
	YellowTimeS      *int     `json:"yellow_time_s,omitempty"`
}

type UpdateIntersectionUseCase struct {
	repo repository.IntersectionRepository
}

func NewUpdateIntersectionUseCase(r repository.IntersectionRepository) *UpdateIntersectionUseCase {
	return &UpdateIntersectionUseCase{repo: r}
}

func (uc *UpdateIntersectionUseCase) Execute(ctx context.Context, id string, requesterRole string, req UpdateIntersectionRequest) (*entities.Intersection, error) {
	if requesterRole != string(userEntities.RoleAdmin) && requesterRole != string(userEntities.RoleOperator) {
		return nil, errors.New("operación denegada")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("id de intersección inválido")
	}

	is, err := uc.repo.GetByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		is.Name = *req.Name
	}
	if req.Description != nil {
		is.Description = req.Description
	}
	if req.Latitude != nil {
		is.Latitude = *req.Latitude
	}
	if req.Longitude != nil {
		is.Longitude = *req.Longitude
	}
	if req.NodeID != nil {
		is.NodeID = req.NodeID
	}
	if req.MaxCongestionPct != nil {
		is.MaxCongestionPct = *req.MaxCongestionPct
	}
	if req.MinGreenTimeS != nil {
		is.MinGreenTimeS = *req.MinGreenTimeS
	}
	if req.MaxGreenTimeS != nil {
		is.MaxGreenTimeS = *req.MaxGreenTimeS
	}
	if req.DefaultGreenS != nil {
		is.DefaultGreenS = *req.DefaultGreenS
	}
	if req.DefaultRedS != nil {
		is.DefaultRedS = *req.DefaultRedS
	}
	if req.YellowTimeS != nil {
		is.YellowTimeS = *req.YellowTimeS
	}

	if err := uc.repo.Update(ctx, is); err != nil {
		return nil, err
	}

	return is, nil
}
