package repository

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/audit/domain/entities"
)

type AuditRepository interface {
	CreateLog(ctx context.Context, log *entities.AuditLog) error
	GetLogs(ctx context.Context, limit int, offset int) ([]entities.AuditLog, error)
}
