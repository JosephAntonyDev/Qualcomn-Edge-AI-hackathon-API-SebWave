package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/domain/repository"
	userEntities "github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
	"github.com/google/uuid"
)

type UpdateStateRequest struct {
	Status           string   `json:"status"`
	OperationMode    string   `json:"operation_mode"`
	CurrentPhase     *string  `json:"current_phase,omitempty"`
	CurrentDensityNS *float64 `json:"current_density_ns,omitempty"`
	CurrentDensityEO *float64 `json:"current_density_eo,omitempty"`
}

type UpdateStateUseCase struct {
	repo repository.IntersectionRepository
}

func NewUpdateStateUseCase(r repository.IntersectionRepository) *UpdateStateUseCase {
	return &UpdateStateUseCase{repo: r}
}

// Para ser llamado manualmente (admin/operator) o internamente
func (uc *UpdateStateUseCase) Execute(ctx context.Context, id string, requesterRole string, req UpdateStateRequest) error {
	if requesterRole != "" && requesterRole != string(userEntities.RoleAdmin) && requesterRole != string(userEntities.RoleOperator) {
		return errors.New("operación denegada")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return errors.New("id de intersección inválido")
	}

	status := entities.NodeStatus(req.Status)
	mode := entities.OperationMode(req.OperationMode)

	return uc.repo.UpdateState(ctx, uid, status, mode, req.CurrentPhase, req.CurrentDensityNS, req.CurrentDensityEO)
}
