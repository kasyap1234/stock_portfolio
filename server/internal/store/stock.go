package store

import (
	"context"
	"errors"
	"math/big"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kasyap1234/portfolio/server/internal/db"
	"github.com/kasyap1234/portfolio/server/internal/models"
	pgxhelper "github.com/kasyap1234/portfolio/server/pkg/pgx"
)

type StockStore interface {
	AddStock(ctx context.Context, PortfolioID uuid.UUID, UserID uuid.UUID, Symbol string, Name string, Quantity int, PurchasePrice float64) (*models.Stock, error)
	DeleteStock(ctx context.Context, id uuid.UUID, UserID uuid.UUID, PortfolioId uuid.UUID) (bool, error)
	UpdateStock(ctx context.Context, id uuid.UUID, UserID uuid.UUID, PortfolioId uuid.UUID,
		symbol, name *string, quantity *int, purchasePrice *float64) (*models.Stock, error)
}

type stockStore struct {
	q *db.Queries
}

func (s *stockStore) AddStock(ctx context.Context, PortfolioID uuid.UUID, UserID uuid.UUID, Symbol string, Name string, Quantity int, PurchasePrice float64) (*models.Stock, error) {
	purchasePrice, err := pgxhelper.FloatToNumeric2Decimal(PurchasePrice)
	if err != nil {
		return &models.Stock{}, errors.New("failed to convert purchaseprice from float to numeric pgtype")
	}
	args := db.AddStockParams{
		PortfolioID:   PortfolioID,
		UserID:        UserID,
		Symbol:        Symbol,
		Name:          Name,
		Quantity:      pgtype.Numeric{Int: big.NewInt(int64(Quantity))},
		PurchasePrice: purchasePrice,
	}

	stock, err := s.q.AddStock(ctx, args)
	modelStock := &models.Stock{
		ID:          stock.ID,
		PortfolioID: stock.PortfolioID,
		UserID:      stock.UserID,
		Symbol:      stock.Symbol,
		Name:        stock.Name,
		Quantity:    int(pgxhelper.NumericToIntFast(stock.Quantity)),
	}
	return modelStock, err
}

func (s *stockStore) DeleteStock(ctx context.Context, id uuid.UUID, UserID uuid.UUID, PortfolioID uuid.UUID) (bool, error) {
	args := db.DeleteStockParams{
		ID:          id,
		UserID:      UserID,
		PortfolioID: PortfolioID,
	}
	err := s.q.DeleteStock(ctx, args)
	if err != nil {
		return false, err
	}
	return true, nil

}

func (s *stockStore) UpdateStock(
	ctx context.Context,
	id uuid.UUID,
	UserID uuid.UUID,
	PortfolioID uuid.UUID,
	symbol, name *string,
	quantity *int,
	purchasePrice *float64,
) (*models.Stock, error) {

	var quantityNum *pgtype.Numeric
	if quantity != nil {
		quantityNum = &pgtype.Numeric{
			Int:   big.NewInt(int64(*quantity)),
			Valid: true,
		}
	}

	var purchasePriceNum *pgtype.Numeric
	if purchasePrice != nil {
		num, err := pgxhelper.FloatToNumeric2Decimal(*purchasePrice)
		if err != nil {
			return nil, errors.New("failed to convert purchase price")
		}
		purchasePriceNum = &num
	}

	args := db.UpdateStockParams{
		ID:            id,
		PortfolioID:   PortfolioID,
		Symbol:        *symbol,
		Name:          *name,
		Quantity:      *quantityNum,
		PurchasePrice: *purchasePriceNum,
	}

	stock, err := s.q.UpdateStock(ctx, args)
	if err != nil {
		return nil, err
	}
	purchasePriceFloat, err := pgxhelper.NumericToFloat(stock.PurchasePrice)
	if err != nil {
		return nil, err
	}
	return &models.Stock{
		ID:             stock.ID,
		PortfolioID:    stock.PortfolioID,
		UserID:         stock.UserID,
		Symbol:         stock.Symbol,
		Name:           stock.Name,
		Quantity:       int(pgxhelper.NumericToIntFast(stock.Quantity)),
		PurcahasePrice: purchasePriceFloat,
		PurchaseDate:   stock.CreatedAt.Time,
	}, nil
}
