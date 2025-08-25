package models

import (
	"time"

	"github.com/google/uuid"
)

type Stock struct {
	ID             uuid.UUID
	PortfolioID    uuid.UUID
	UserID         uuid.UUID
	Symbol         string
	Name           string
	Quantity       int
	PurcahasePrice float64
	PurchaseDate   time.Time
}
