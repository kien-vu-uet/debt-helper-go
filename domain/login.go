package domain

import (
	"context"
)

type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"` // Changed from Email, added json tag
	Password string `form:"password" json:"password" binding:"required"` // Added json tag
}

type LoginResponse struct {
	Token string `json:"token"` // Changed from AccessToken and RefreshToken
}

// Added for /access-token/extend endpoint
type ExtendAccessTokenRequest struct {
	Token string `json:"token" binding:"required"`
}

// Added for /refresh-token/revoke endpoint
type RevokeRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// Added for /verify-username endpoint
type VerifyUsernameRequest struct {
	Username string `json:"username" binding:"required"`
}

type VerifyUsernameResponse struct {
	Available bool `json:"available"`
}

type LoginUsecase interface {
	GetUserByUsername(c context.Context, username string) (User, error) // Changed from GetUserByEmail
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error) // Kept for internal use if needed
	ExtractIDFromToken(tokenString string, secret string) (int64, error)                       // Added for ExtendAccessToken
	GetUserByID(c context.Context, id int64) (User, error)                                     // Added for ExtendAccessToken
	RevokeRefreshToken(c context.Context, tokenString string) error                            // Added for RevokeRefreshToken
}
