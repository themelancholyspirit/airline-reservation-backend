package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key") // Replace with a secure secret key
type Claims struct {
	UserID uint    `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, email, name string) (string, error) {
	// Define token expiration time (24 hours)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the claims, which includes the user's ID, email, and standard claims
	claims := &Claims{
		UserID: userID,
		Email:  email,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()), // Token is valid immediately
		},
	}

	// Create the token using the claims and HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key (jwtKey)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Parse the token with claims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil // Your secret key
	})

	// Check for parsing errors
	if err != nil {
		// You can make the error more informative if needed
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("signature validation failed")
		}
		return nil, fmt.Errorf("could not parse token: %w", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Return the claims (token is valid)
	return claims, nil
}

func RefreshToken(claims *Claims) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.NotBefore = jwt.NewNumericDate(time.Now())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
