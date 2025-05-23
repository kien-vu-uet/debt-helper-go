package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/kien-vu-uet/debt-helper-go/domain"
	"github.com/kien-vu-uet/debt-helper-go/internal/tokenutil"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByUsername(c context.Context, username string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByUsername(ctx, username)
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

// ExtractIDFromToken extracts the user ID from the token.
func (lu *loginUsecase) ExtractIDFromToken(tokenString string, secret string) (int64, error) {
	return tokenutil.ExtractIDFromToken(tokenString, secret)
}

// GetUserByID retrieves a user by their ID.
func (lu *loginUsecase) GetUserByID(c context.Context, id int64) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByID(ctx, id)
}

// RevokeRefreshToken is a placeholder for the actual revoke logic.
// In a real application, this would involve a denylist or other mechanism.
func (lu *loginUsecase) RevokeRefreshToken(c context.Context, tokenString string) error {
	// Placeholder logic: In a real scenario, you might add the token to a denylist (e.g., in Redis)
	// until its original expiry. For this example, we\'ll just print a message.
	// Actual implementation depends on how you manage token validity.
	fmt.Printf("Token %s would be revoked here.\n", tokenString)
	// Returning nil to simulate success as per current simplified model.
	// If you had a persistent denylist, you would interact with it here.
	return nil
}
