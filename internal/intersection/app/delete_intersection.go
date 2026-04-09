package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/domain/repository"
	userEntities "github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
	"github.com/google/uuid"
)

type DeleteIntersectionUseCase struct {
	repo repository.IntersectionRepository
}

func NewDeleteIntersectionUseCase(r repository.IntersectionRepository) *DeleteIntersectionUseCase {
	return &DeleteIntersectionUseCase{repo: r}
}

func (uc *DeleteIntersectionUseCase) Execute(ctx context.Context, id string, requesterRole string) error {
	if requesterRole != string(userEntities.RoleAdmin) {
		return errors.New("operación denegada")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return errors.New("id de intersección inválido")
	}

	return uc.repo.Delete(ctx, uid)
}
