package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kien-vu-uet/debt-helper-go/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db        *gorm.DB // Changed to *gorm.DB
	tableName string   // This might be less relevant if GORM uses struct names or tags
}

// NewUserRepository creates a new user repository with GORM
func NewUserRepository(db *gorm.DB, tableName string) domain.UserRepository { // Changed *sql.DB to *gorm.DB
	return &userRepository{
		db:        db,
		tableName: tableName, // Keep for now, but GORM typically infers table from struct
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// GORM handles table name via struct or can be specified with .Table()
	result := ur.db.WithContext(c).Create(user) // GORM Create
	if result.Error != nil {
		// Check for specific SQL errors, like duplicate email
		// GORM might return a different error type, adjust as needed
		// For now, using a general check. You might need to inspect result.Error more closely.
		// Example: if errors.Is(result.Error, gorm.ErrDuplicatedKey) or similar
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) { // This is a common GORM error for duplicates
			return domain.ErrDuplicateEntry
		}
		return fmt.Errorf("creating user with GORM: %w", result.Error)
	}
	// ID is automatically populated by GORM after Create
	return nil
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	var users []domain.User
	// GORM Find. Password should be excluded by struct tags `json:"-"` or select specific fields.
	// To explicitly exclude password: result := ur.db.WithContext(c).Select("id", "name", "email", "created_at", "updated_at").Find(&users)
	result := ur.db.WithContext(c).Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("fetching users with GORM: %w", result.Error)
	}
	return users, nil
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	result := ur.db.WithContext(c).Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.User{}, domain.ErrNotFound
		}
		return domain.User{}, fmt.Errorf("getting user by email %s with GORM: %w", email, result.Error)
	}
	return user, nil
}

func (ur *userRepository) GetByUsername(c context.Context, username string) (domain.User, error) {
	var user domain.User
	// Corrected query to use "name" column as per db tag in domain.User
	result := ur.db.WithContext(c).Where("name = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.User{}, domain.ErrNotFound
		}
		return domain.User{}, fmt.Errorf("getting user by username %s with GORM: %w", username, result.Error)
	}
	return user, nil
}

func (ur *userRepository) GetByID(c context.Context, id int64) (domain.User, error) {
	var user domain.User
	result := ur.db.WithContext(c).First(&user, id) // GORM can find by primary key directly
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.User{}, domain.ErrNotFound
		}
		return domain.User{}, fmt.Errorf("getting user by ID %d with GORM: %w", id, result.Error)
	}
	return user, nil
}

func (ur *userRepository) Update(c context.Context, user *domain.User) error {
	user.UpdatedAt = time.Now()
	result := ur.db.WithContext(c).Model(&domain.User{}).Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		return fmt.Errorf("updating user with GORM: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound // Or a more specific error like ErrNothingChanged
	}
	return nil
}
