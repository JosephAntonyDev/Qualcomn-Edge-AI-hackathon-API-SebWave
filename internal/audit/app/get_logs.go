package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/audit/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/audit/domain/repository"
)

type GetLogsUseCase struct {
	Repo repository.AuditRepository
}

func NewGetLogsUseCase(repo repository.AuditRepository) *GetLogsUseCase {
	return &GetLogsUseCase{Repo: repo}
}

func (uc *GetLogsUseCase) Execute(ctx context.Context, limit, offset int) ([]entities.AuditLog, error) {
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}
	return uc.Repo.GetLogs(ctx, limit, offset)
}
