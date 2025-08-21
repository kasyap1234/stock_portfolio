package store

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kasyap1234/portfolio/server/internal/db"
	"github.com/kasyap1234/portfolio/server/internal/db/portfolio"
	"github.com/kasyap1234/portfolio/server/internal/models"
)

type PortfolioStore interface {
CreatePortfolio(user_id uuid.UUID,name string,invested_value string , current_value string)(*models.Portfolio,error)
}

type portfolioStore  struct {
	q *db.Queries
}

func(p*portfolioStore)CreatePortfolio(user_id uuid.UUID,name,invested_value,current_value string)(*models.Portfolio,error){
	args := portfolio.CreatePortfolioParams{
		UserID: user_id,
		Name: name,
		InvestedValue: pgtype.Text{String: invested_value},
		CurrentValue: pgtype.Text{String: current_value},

	}
	folio,err :=portfolio.Queries.CreatePortfolio(
		
	)
}