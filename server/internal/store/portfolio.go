package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kasyap1234/portfolio/server/internal/db"
	"github.com/kasyap1234/portfolio/server/internal/models"
)

type PortfolioStore interface {
	CreatePortfolio(ctx context.Context, user_id uuid.UUID, name string, invested_value string, current_value string) (*models.Portfolio, error)
	ListPortfolios(ctx context.Context, user_id uuid.UUID) ([]*models.Portfolio, error)
}

type portfolioStore struct {
	q *db.Queries
}

func (p *portfolioStore) CreatePortfolio(ctx context.Context, user_id uuid.UUID, name, invested_value, current_value string) (*models.Portfolio, error) {
	args := db.CreatePortfolioParams{
		UserID:        user_id,
		Name:          name,
		InvestedValue: pgtype.Text{String: invested_value},
		CurrentValue:  pgtype.Text{String: current_value},
	}
	folio, err := p.q.CreatePortfolio(ctx, args)
	portfolio := &models.Portfolio{
		ID:            folio.ID,
		UserID:        folio.UserID,
		Name:          folio.Name,
		InvestedValue: folio.InvestedValue.String,
		CurrentValue:  folio.CurrentValue.String,
		CreatedAt:     folio.CreatedAt.Time,
		UpdatedAt:     folio.UpdatedAt.Time,
	}
	return portfolio, err

}

func (p *portfolioStore) ListPortfolio(ctx context.Context, user_id uuid.UUID) ([]*models.Portfolio, error) {
	dbPortfolio, err := p.q.ListPortfolio(ctx, user_id)
	if err != nil {
		return nil, err
	}
	portfolios := make([]*models.Portfolio, len(dbPortfolio))
	for _, pt := range dbPortfolio {
		portfolios = append(portfolios, &models.Portfolio{
			ID:            pt.ID,
			UserID:        pt.UserID,
			Name:          pt.Name,
			InvestedValue: pt.InvestedValue.String,
			CurrentValue:  pt.CurrentValue.String,
		})
	}
	return portfolios, err

}
