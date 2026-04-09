package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/repository"
	"github.com/google/uuid"
)

type GetUserByIDUseCase struct {
	repo repository.UserRepository
}

func NewGetUserByIDUseCase(r repository.UserRepository) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{repo: r}
}

func (uc *GetUserByIDUseCase) Execute(ctx context.Context, id string) (*PublicUserResponse, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("id de usuario inválido")
	}

	user, err := uc.repo.GetUserByID(ctx, parsedUUID)
	if err != nil {
		return nil, errors.New("usuario no encontrado")
	}

	return &PublicUserResponse{
		Username: user.Username,
		FullName: user.FullName,
		IsActive: user.IsActive,
	}, nil
}
