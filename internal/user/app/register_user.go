package app

import (
	"context"
	"errors"
	"time"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/ports"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/repository"
	"github.com/google/uuid"
)

type RegisterUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name,omitempty"`
}

type RegisterUserUseCase struct {
	repo   repository.UserRepository
	hasher ports.PasswordHasher
}

func NewRegisterUserUseCase(r repository.UserRepository, h ports.PasswordHasher) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		repo:   r,
		hasher: h,
	}
}

func (uc *RegisterUserUseCase) Execute(ctx context.Context, req RegisterUserRequest) (*entities.User, error) {
	existingUser, _ := uc.repo.GetUserByUsername(ctx, req.Username)
	if existingUser != nil {
		return nil, errors.New("el usuario ya está registrado en el sistema")
	}

	hashedPassword, err := uc.hasher.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("error al procesar la contraseña")
	}

	newUser := &entities.User{
		ID:           uuid.New(),
		Username:     req.Username,
		Email:        req.Email,
		FullName:     req.FullName,
		PasswordHash: hashedPassword,
		Role:         entities.RoleViewer,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := uc.repo.CreateUser(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}
