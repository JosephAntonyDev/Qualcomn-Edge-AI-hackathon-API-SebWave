package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/repository"
)

type PublicUserResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name,omitempty"`
	IsActive bool   `json:"is_active"`
}

type GetUserByUsernameUseCase struct {
	repo repository.UserRepository
}

func NewGetUserByUsernameUseCase(r repository.UserRepository) *GetUserByUsernameUseCase {
	return &GetUserByUsernameUseCase{repo: r}
}

func (uc *GetUserByUsernameUseCase) Execute(ctx context.Context, username string) (*PublicUserResponse, error) {
	user, err := uc.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("usuario no encontrado")
	}

	return &PublicUserResponse{
		Username: user.Username,
		FullName: user.FullName,
		IsActive: user.IsActive,
	}, nil
}
