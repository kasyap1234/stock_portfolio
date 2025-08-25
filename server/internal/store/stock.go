package store

import (
	"time"

	"github.com/google/uuid"
)

type StockStore interface {
	AddStock(PortfolioID uuid.UUID, UserID uuid.UUID, Symbol string, Name string, Quantity int, PurchasePrice float64, PurchaseDate time.Time)
	DeleteStock(id uuid.UUID, UserID uuid.UUID, PortfolioId uuid.UUID)
	UpdateStock
}
