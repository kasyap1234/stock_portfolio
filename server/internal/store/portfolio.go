package store

import (
	"github.com/google/uuid"
	"github.com/kasyap1234/portfolio/server/internal/db"
	"github.com/kasyap1234/portfolio/server/internal/models"
)

type PortfolioStore interface {
CreatePortfolio(user_id uuid.UUID,name string,invested_value string , current_value string)(*models.Portfolio,error)
}

type portfolioStore  struct {
	q *db.Queries
}

