package models

import (
	"time"

	"github.com/google/uuid"
)

type Portfolio struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	Name          string    `json:"name"`
	InvestedValue string    `json:"invested_value"`
	CurrentValue  string    `json:"current_value"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
