package store

import (
	"context"

	"github.com/kasyap1234/portfolio/server/internal/db"
)

type AuthStore interface {
	RegisterUser(ctx context.Context, user *db.User) (bool, error)
	LoginUser(ctx context.Context, user *db.User) (*db.User, error)
}

type authStore struct {
}
