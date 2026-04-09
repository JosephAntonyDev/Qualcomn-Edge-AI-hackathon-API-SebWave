package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/domain/repository"
	"github.com/google/uuid"
)

type GetIntersectionUseCase struct {
	repo repository.IntersectionRepository
}

func NewGetIntersectionUseCase(r repository.IntersectionRepository) *GetIntersectionUseCase {
	return &GetIntersectionUseCase{repo: r}
}

func (uc *GetIntersectionUseCase) Execute(ctx context.Context, id string) (*entities.Intersection, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("id de intersección inválido")
	}

	return uc.repo.GetByID(ctx, uid)
}
