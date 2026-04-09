package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/domain/repository"
)

type HeartbeatRequest struct {
	SerialNumber    string  `json:"serial_number" binding:"required"`
	FirmwareVersion *string `json:"firmware_version,omitempty"`
}

type HeartbeatUseCase struct {
	repo repository.IntersectionRepository
}

func NewHeartbeatUseCase(r repository.IntersectionRepository) *HeartbeatUseCase {
	return &HeartbeatUseCase{repo: r}
}

func (uc *HeartbeatUseCase) Execute(ctx context.Context, req HeartbeatRequest) error {
	// The edge node itself calls this using an IoT API key or basic auth, we assume authorization is handled middleware or IoT Gateway
	return uc.repo.Heartbeat(ctx, req.SerialNumber, req.FirmwareVersion)
}
