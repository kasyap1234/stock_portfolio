package auth

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/kasyap1234/portfolio/server/internal/db"
	"github.com/kasyap1234/portfolio/server/internal/models"
	"github.com/kasyap1234/portfolio/server/internal/store"
	"github.com/kasyap1234/portfolio/server/pkg/email"
	jwtkeys "github.com/kasyap1234/portfolio/server/pkg/jwt"
	"github.com/kasyap1234/portfolio/server/pkg/security"
)

// exposes auth service methods
type AuthService interface {
	RegisterUser(ctx context.Context, user *models.User) (*models.UserResponse, error)
	VerifyEmail(ctx context.Context, email string) (bool, error)
	LoginUser(ctx context.Context, user *models.User) (*LoginResponse, error)
}

type authService struct {
	store store.AuthStore
}

func NewAuthService(store store.AuthStore) AuthService {
	return &authService{store: store}
}

func (a *authService) RegisterUser(ctx context.Context, user *models.User) (*models.UserResponse, error) {
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("email or password cannot be empty")
	}
	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return &models.UserResponse{}, err
	}
	// Convert domain → db
	dbUser := db.User{
		ID:       uuid.New(),
		Email:    user.Email,
		Password: hashedPassword,
		Name:     user.Name,
	}

	createdUser, err := a.store.RegisterUser(ctx, dbUser)
	if err != nil {
		return nil, err
	}
	err = email.SendTokenEmail(email.GenerateEmailToken(), user.Email)
	if err != nil {
		return &models.UserResponse{}, err
	}
	check, err := a.VerifyEmail(ctx, user.Email)
	if err != nil {
		check = false
	}
	// Convert db → domain
	return &models.UserResponse{
		ID:            createdUser.ID,
		Email:         createdUser.Email,
		Name:          createdUser.Name,
		EmailVerified: check,
	}, nil
}

func (a *authService) VerifyEmail(ctx context.Context, email string) (bool, error) {
	// verify email service function
	return true, nil
}

// login response struct for loginuser service response message .
type LoginResponse struct {
	User  *models.UserResponse `json:"user"`
	Token string               `json:"token"`
}

// login user service function
func (a *authService) LoginUser(ctx context.Context, user *models.User) (*LoginResponse, error) {
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("email or password empty")
	}
	dbUser, err := a.store.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, errors.New("invalid email address")
	}
	if !security.CheckPassword(user.Password, dbUser.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := jwtkeys.GenerateJWT(dbUser.ID, dbUser.Email)
	if err != nil {
		return nil, errors.New("failed to generate authentication token")
	}
	userResponse := &models.UserResponse{
		ID:        dbUser.ID,
		Email:     dbUser.Email,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt.Time.Format(time.RFC3339),
	}
	return &LoginResponse{userResponse, token}, nil
}
