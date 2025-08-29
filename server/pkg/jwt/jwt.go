package jwtkeys

import (
	"crypto/rsa"
	"fmt"
	"io"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

// InitKeys loads RSA private and public keys from PEM files
func InitKeys() error {
	// --- Load private key ---
	privFile, err := os.Open("private.pem")
	if err != nil {
		return err
	}
	defer privFile.Close()

	privData, err := io.ReadAll(privFile)
	if err != nil {
		return err
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privData)
	if err != nil {
		return err
	}

	// --- Load public key ---
	pubFile, err := os.Open("public.pem")
	if err != nil {
		return err
	}
	defer pubFile.Close()

	pubData, err := io.ReadAll(pubFile)
	if err != nil {
		return err
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(pubData)
	if err != nil {
		return err
	}

	return nil
}

type Claims struct {
	UserID   uuid.UUID
	Username string
	jwt.RegisteredClaims
}

func GenerateJWT(userID uuid.UUID, username string, typ TokenType) (string, error) {
	var expires time.Time

	switch typ {
	case jwtkeys.AccessToken:
		expires = time.Now().Add(15 * time.Hour)
	case jwtkeys.RefreshToken:
		expires = time.Now().Add(24 * 15 * time.Hour)
	}
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expires),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}

func ParseJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	// Parse the token with claims
	tok, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		// Ensure signing method is RS256
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !tok.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func GenerateTo