package ports

import (
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
	"github.com/google/uuid"
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePasswords(hashedPassword string, providedPassword string) error
}

type TokenManager interface {
	GenerateToken(userID uuid.UUID, role entities.UserRole) (string, error)
	ValidateToken(token string) (bool, map[string]interface{}, error)
}
