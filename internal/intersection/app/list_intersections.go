package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/domain/repository"
)

type ListIntersectionsUseCase struct {
	repo repository.IntersectionRepository
}

func NewListIntersectionsUseCase(r repository.IntersectionRepository) *ListIntersectionsUseCase {
	return &ListIntersectionsUseCase{repo: r}
}

func (uc *ListIntersectionsUseCase) Execute(ctx context.Context) ([]*entities.Intersection, error) {
	return uc.repo.GetAll(ctx)
}
