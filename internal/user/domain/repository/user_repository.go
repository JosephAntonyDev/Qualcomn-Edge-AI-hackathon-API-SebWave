package repository

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User) error
	UpdateUser(ctx context.Context, user *entities.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	GetMyProfile(ctx context.Context, userID uuid.UUID) (*entities.User, error)
}
