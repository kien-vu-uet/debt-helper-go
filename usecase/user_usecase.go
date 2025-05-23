package usecase

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/kien-vu-uet/debt-helper-go/domain"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

// Fetch implements domain.UserUsecase
func (uu *userUsecase) Fetch(c context.Context) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.Fetch(ctx)
}

// GetByEmail implements domain.UserUsecase
func (uu *userUsecase) GetByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.GetByEmail(ctx, email)
}

// GetByID fetches a user by their ID.
// It calls the UserRepository to retrieve the user from the database.
func (uu *userUsecase) GetByID(c context.Context, id string) (domain.User, error) {
	// Convert id from string to int64 for repository interaction
	int64ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return domain.User{}, fmt.Errorf("invalid user ID format: %w", err)
	}
	return uu.userRepository.GetByID(c, int64ID)
}

// GetByUsername implements domain.UserUsecase
// This adheres to the 'user' naming convention, avoiding 'profile'.
func (uu *userUsecase) GetByUsername(c context.Context, username string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.GetByUsername(ctx, username)
}

// Update implements domain.UserUsecase
// This adheres to the 'user' naming convention, avoiding 'profile'.
func (uu *userUsecase) Update(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.Update(ctx, user)
}

// Create implements domain.UserUsecase
func (uu *userUsecase) Create(c context.Context, request *domain.CreateUserRequest) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	// You might want to hash the password here before creating the user
	// For example, using bcrypt.GenerateFromPassword
	// This is a placeholder for password hashing logic
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword), // Store the hashed password
		FullName: request.FullName,
		Role:     request.Role,
	}

	if user.Role == "" {
		user.Role = "user" // Default role
	}

	err = uu.userRepository.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	// The Create method in the repository might not return the created user object with ID.
	// If it does, you can return it directly. Otherwise, you might need to fetch it after creation if the ID is needed.
	// For now, returning the input user object (which might not have the ID or other DB-generated fields like CreatedAt)
	// Or, if the repository's Create method updates the user pointer with the ID, this is fine.
	// Consider fetching the user by email or username after creation to get the full object with ID.
	// For simplicity, returning the user object as is.
	// Clear the password from the response
	user.Password = ""
	return user, nil
}

// Helper function to hash password (example)
// You should place this in a more appropriate package like internal/auth or internal/security
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// List implements domain.UserUsecase
func (uu *userUsecase) List(c context.Context, page int, limit int) ([]domain.User, error) {
	// ctx, cancel := context.WithTimeout(c, uu.contextTimeout) // Marked as unused
	// defer cancel()
	// Assuming your repository has a List method.
	// If not, you'll need to implement it in the repository layer.
	// This is a placeholder for the actual implementation.
	// users, err := uu.userRepository.List(ctx, page, limit)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to list users: %w", err)
	// }
	// return users, nil
	return nil, fmt.Errorf("List method not implemented yet") // Placeholder
}

// Delete implements domain.UserUsecase
func (uu *userUsecase) Delete(c context.Context, userID string) error {
	// ctx, cancel := context.WithTimeout(c, uu.contextTimeout) // Marked as unused
	// defer cancel()

	// int64ID, err := strconv.ParseInt(userID, 10, 64) // Marked as unused
	// if err != nil {
	// 	return fmt.Errorf("invalid user ID format: %w", err)
	// }
	// Assuming your repository has a Delete method.
	// If not, you'll need to implement it in the repository layer.
	// return uu.userRepository.Delete(ctx, int64ID)
	return fmt.Errorf("Delete method not implemented in repository yet") // Placeholder
}

// InitiatePasswordReset implements domain.UserUsecase
func (uu *userUsecase) InitiatePasswordReset(c context.Context, email string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	// Placeholder for password reset logic
	// 1. Validate email
	// 2. Generate a reset token
	// 3. Store token with expiry (e.g., in DB or Redis)
	// 4. Send email with reset link
	_, err := uu.userRepository.GetByEmail(ctx, email) // Use ctx here
	if err != nil {
		return fmt.Errorf("user with email %s not found or error fetching: %w", email, err)
	}
	// Actual implementation would involve token generation and email sending
	return fmt.Errorf("password reset initiation not fully implemented yet")
}

// UpdateFullname implements domain.UserUsecase
func (uu *userUsecase) UpdateFullname(c context.Context, userID string, fullname string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	int64ID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format: %w", err)
	}

	user, err := uu.userRepository.GetByID(ctx, int64ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	user.FullName = fullname
	err = uu.userRepository.Update(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user fullname: %w", err)
	}
	user.Password = "" // Clear password from response
	return &user, nil
}

// UpdateAvatar implements domain.UserUsecase
func (uu *userUsecase) UpdateAvatar(c context.Context, userID string, avatarURL string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	int64ID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format: %w", err)
	}

	user, err := uu.userRepository.GetByID(ctx, int64ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	user.Avatar = avatarURL
	err = uu.userRepository.Update(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user avatar: %w", err)
	}
	user.Password = "" // Clear password from response
	return &user, nil
}
