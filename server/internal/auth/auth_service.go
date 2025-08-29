package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kasyap1234/portfolio/server/internal/db"
	"github.com/kasyap1234/portfolio/server/internal/models"
	"github.com/kasyap1234/portfolio/server/internal/store"
	"github.com/kasyap1234/portfolio/server/pkg/email"
	jwtkeys "github.com/kasyap1234/portfolio/server/pkg/jwt"
	redis "github.com/kasyap1234/portfolio/server/pkg/redis"
	"github.com/kasyap1234/portfolio/server/pkg/security"
)

// exposes auth service methods
type AuthService interface {
	RegisterUser(ctx context.Context, user *models.User) (*models.UserResponse, error)
	VerifyEmail(ctx context.Context, email string, token string) (bool, error)
	LoginUser(ctx context.Context, user *models.User) (*LoginResponse, error)
	LogoutUser(ctx context.Context, token string) error
}

type authService struct {
	store       store.AuthStore
	redisClient redis.RedisClient
}

func NewAuthService(store store.AuthStore, redisClient redis.RedisClient) AuthService {
	return &authService{store: store, redisClient: redisClient}
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
	token := email.GenerateEmailToken()

	err = email.SendTokenEmail(token, user.Email)
	if err != nil {
		return &models.UserResponse{}, err
	}
	err = a.StoreEmailVerificationToken(ctx, user.Email, token)
	if err != nil {
		log.Printf("failed to store email verification token in redis cache")
	}

	check, err := a.VerifyEmail(ctx, user.Email, token)

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

func (a *authService) VerifyEmail(ctx context.Context, email string, token string) (bool, error) {
	// verify email service function
	dbToken, err := a.GetEmailToken(ctx, email)
	if dbToken == token {
		return true, nil
	}

	return false, err
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

	accessToken, err := jwtkeys.GenerateJWT(dbUser.ID, dbUser.Email, jwtkeys.AccessToken)

	if err != nil {
		return nil, errors.New("failed to generate authentication token")
	}
	userResponse := &models.UserResponse{
		ID:        dbUser.ID,
		Email:     dbUser.Email,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt.Time.Format(time.RFC3339),
	}
	return &LoginResponse{userResponse, accessToken}, nil
}

func (a *authService) StoreEmailVerificationToken(ctx context.Context, email, token string) error {
	key := fmt.Sprintf("email_verification_token:%s", token)
	StatusCmd := a.redisClient.Set(ctx, key, token, time.Hour)
	if err := StatusCmd.Err(); err != nil {
		return err
	}
	return nil
}

func (a *authService) GetEmailToken(ctx context.Context, email string) (string, error) {
	// Use the same key format as in storeEmailVerificationToken
	key := fmt.Sprintf("email_verification_token:%s", email)

	stringCmd := a.redisClient.Get(ctx, key)
	if err := stringCmd.Err(); err != nil {
		return "", err
	}

	// Extract the actual token value
	token, err := stringCmd.Result()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *authService) LogoutUser(ctx context.Context, token string) error {
	claims, err := jwtkeys.ParseJWT(token)
	if err != nil {
		return errors.New("invalid token")
	}
	expiresAt := claims
}
