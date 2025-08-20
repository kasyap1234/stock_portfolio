package jwtkeys

import (
	"crypto/rsa"
	"fmt"
	"io"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
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
	UserID   int
	Username string
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
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
