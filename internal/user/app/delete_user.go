package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/repository"
	"github.com/google/uuid"
)

type DeleteUserUseCase struct {
	repo repository.UserRepository
}

func NewDeleteUserUseCase(r repository.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{repo: r}
}

func (uc *DeleteUserUseCase) Execute(ctx context.Context, targetUserID, requesterID string, requesterRole string) error {
	if requesterRole != string(entities.RoleAdmin) && requesterRole != string(entities.RoleOperator) && targetUserID != requesterID {
		return errors.New("operación denegada: no tienes permiso para eliminar este usuario")
	}

	parsedUUID, err := uuid.Parse(targetUserID)
	if err != nil {
		return errors.New("id de usuario inválido")
	}

	err = uc.repo.DeleteUser(ctx, parsedUUID)
	if err != nil {
		return errors.New("usuario no encontrado o ya eliminado")
	}

	return nil
}
