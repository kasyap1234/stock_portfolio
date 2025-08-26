package store

import (
	"context"

	"github.com/kasyap1234/portfolio/server/internal/db"
)

type AuthStore interface {
	RegisterUser(ctx context.Context, user db.User) (db.User, error)
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
	LoginUser(ctx context.Context, user db.User) (db.User, error)
}

type authStore struct {
	q *db.Queries
}

func NewAuthStore(q *db.Queries) AuthStore {
	return &authStore{q: q}
}

func (a *authStore) CreateUser(ctx context.Context, user db.User) (db.User, error) {
	return a.q.CreateUser(ctx, db.CreateUserParams{
		Email:    user.Email,
		Password: user.Password, // already hashed
		Name:     user.Name,
	})
}

// RegisterUser creates a new user and returns the created user.
func (a *authStore) RegisterUser(ctx context.Context, user db.User) (db.User, error) {
	return a.CreateUser(ctx, user)
}

func (a *authStore) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	return a.q.GetUserByEmail(ctx, email)
}

func(a*authStore)LoginUser(ctx context.Context,user db.User)(db.User,error){}
