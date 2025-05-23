package tokenutil

import (
	"fmt"
	"strconv" // Added import for strconv
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/kien-vu-uet/debt-helper-go/domain"
)

func CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := &domain.JwtCustomClaims{
		Name: user.Username,
		ID:   fmt.Sprintf("%d", user.ID), // Convert int64 to string
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	claimsRefresh := &domain.JwtCustomRefreshClaims{
		ID: fmt.Sprintf("%d", user.ID), // Convert int64 to string
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return rt, err
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"]) // Changed to lowercase
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (int64, error) { // Changed return type from string to int64
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"]) // Changed to lowercase
		}
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err // Return 0 for int64 in case of error
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid { // Combined conditions
		return 0, fmt.Errorf("invalid token") // Changed to lowercase
	}

	idStr, ok := claims["id"].(string)
	if !ok {
		return 0, fmt.Errorf("id claim is not a string") // Changed to lowercase
	}

	// Use strconv.ParseInt instead of jwt.ParseInt
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse id from token: %w", err) // Changed to lowercase
	}
	return id, nil // Return parsed int64 ID
}
