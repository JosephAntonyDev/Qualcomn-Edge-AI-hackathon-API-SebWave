package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/repository"
	"github.com/google/uuid"
)

type UpdateUserRequest struct {
	Email    *string `json:"email,omitempty"`
	FullName *string `json:"full_name,omitempty"`
	IsActive *bool   `json:"is_active,omitempty"`
	Role     *string `json:"role,omitempty"`
}

type UpdateUserUseCase struct {
	repo repository.UserRepository
}

func NewUpdateUserUseCase(r repository.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{repo: r}
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, targetUserID, requesterID string, requesterRole string, req UpdateUserRequest) (*entities.User, error) {
	if requesterRole != string(entities.RoleAdmin) && requesterRole != string(entities.RoleOperator) && targetUserID != requesterID {
		return nil, errors.New("operación denegada: no tienes permiso para actualizar este usuario")
	}

	parsedUUID, err := uuid.Parse(targetUserID)
	if err != nil {
		return nil, errors.New("id de usuario inválido")
	}

	user, err := uc.repo.GetUserByID(ctx, parsedUUID)
	if err != nil {
		return nil, errors.New("usuario no encontrado")
	}

	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.FullName != nil {
		user.FullName = *req.FullName
	}
	if req.IsActive != nil {
		if requesterRole == string(entities.RoleAdmin) || requesterRole == string(entities.RoleOperator) {
			user.IsActive = *req.IsActive
		} else if *req.IsActive != user.IsActive {
			return nil, errors.New("operación denegada: no tienes permiso para cambiar el estado activo")
		}
	}
	if req.Role != nil {
		if requesterRole == string(entities.RoleAdmin) {
			user.Role = entities.UserRole(*req.Role)
		} else if string(user.Role) != *req.Role {
			return nil, errors.New("operación denegada: solo un administrador puede cambiar el rol")
		}
	}

	if err := uc.repo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
