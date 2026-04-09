package repository

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/entities"
	"github.com/google/uuid"
)

type AlertFilter struct {
	IntersectionID *uuid.UUID
	Type           *entities.AlertType
	Severity       *entities.AlertSeverity
	Status         *entities.AlertStatus
	Limit          int
	Offset         int
}

type AlertRepository interface {
	Create(ctx context.Context, alert *entities.Alert) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Alert, error)
	List(ctx context.Context, filter AlertFilter) ([]*entities.Alert, error)
	Update(ctx context.Context, alert *entities.Alert) error
	Delete(ctx context.Context, id uuid.UUID) error
	ChangeStatus(ctx context.Context, id uuid.UUID, status entities.AlertStatus, resolvedBy *uuid.UUID, note *string) error
}
