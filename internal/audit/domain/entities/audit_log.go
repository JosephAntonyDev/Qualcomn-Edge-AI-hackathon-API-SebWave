package entities

import "time"

type AuditLog struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	Action         string    `json:"action"`
	TargetResource string    `json:"target_resource"`
	TargetID       string    `json:"target_id"`
	Details        string    `json:"details"`
	Timestamp      time.Time `json:"timestamp"`
}
