package app

import (
	"context"
	"time"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/audit/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/audit/domain/repository"
)

type CreateLogUseCase struct {
	Repo repository.AuditRepository
}

func NewCreateLogUseCase(repo repository.AuditRepository) *CreateLogUseCase {
	return &CreateLogUseCase{Repo: repo}
}

func (uc *CreateLogUseCase) Execute(ctx context.Context, userID, action, targetResource, targetID, details string) error {
	log := &entities.AuditLog{
		UserID:         userID,
		Action:         action,
		TargetResource: targetResource,
		TargetID:       targetID,
		Details:        details,
		Timestamp:      time.Now(),
	}
	return uc.Repo.CreateLog(ctx, log)
}
