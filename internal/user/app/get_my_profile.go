package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/repository"
	"github.com/google/uuid"
)

type GetMyProfileUseCase struct {
	repo repository.UserRepository
}

func NewGetMyProfileUseCase(r repository.UserRepository) *GetMyProfileUseCase {
	return &GetMyProfileUseCase{repo: r}
}

func (uc *GetMyProfileUseCase) Execute(ctx context.Context, requesterID string) (*entities.User, error) {
	parsedUUID, err := uuid.Parse(requesterID)
	if err != nil {
		return nil, errors.New("id de usuario inválido en token")
	}

	user, err := uc.repo.GetMyProfile(ctx, parsedUUID)
	if err != nil {
		return nil, errors.New("perfil no encontrado")
	}

	return user, nil
}
