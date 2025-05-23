package usecase

import (
	"context"
	"strconv" // Added for string to int64 conversion
	"time"

	"github.com/kien-vu-uet/debt-helper-go/domain"
	"github.com/kien-vu-uet/debt-helper-go/internal/tokenutil"
)

type refreshTokenUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewRefreshTokenUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (rtu *refreshTokenUsecase) GetUserByID(c context.Context, id string) (domain.User, error) { // Changed parameter name from email to id
	ctx, cancel := context.WithTimeout(c, rtu.contextTimeout)
	defer cancel()

	userID, err := strconv.ParseInt(id, 10, 64) // Convert string id to int64
	if err != nil {
		return domain.User{}, err // Handle conversion error
	}
	return rtu.userRepository.GetByID(ctx, userID) // Use converted userID
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	id, err := tokenutil.ExtractIDFromToken(requestToken, secret)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(id, 10), nil
}
