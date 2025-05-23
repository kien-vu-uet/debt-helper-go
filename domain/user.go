package domain

import (
	"context"
)

const (
	CollectionUser = "users" // This can be renamed to TableUser or used as table name
)

type User struct {
	BaseModel          // Embed BaseModel
	Username    string `json:"username" bson:"username" db:"username"`    // Updated json tag to "username"
	FullName    string `json:"full_name" bson:"full_name" db:"full_name"` // Added db tag
	SlackID     string `json:"slack_id" bson:"slack_id" db:"slack_id"`    // Added db tag
	Email       string `json:"email" bson:"email" db:"email"`             // Added db tag
	Password    string `json:"-" bson:"password" db:"password"`           // json:"-" to hide from responses, added db tag
	Avatar      string `json:"avatar" bson:"avatar" db:"avatar"`          // Added db tag
	Role        string `json:"role" bson:"role" db:"role"`                // Added db tag
	AccessToken string `json:"token,omitempty" bson:"-" db:"-"`           // bson:"-" and db:"-" as it's not stored
}

// TableName returns the table name for the User model.
// It uses the 'user' naming convention, avoiding 'profile'.
func (User) TableName() string {
	return CollectionUser
}

// ResetPasswordRequest defines the structure for a password reset request.
type ResetPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// UpdateFullnameRequest defines the structure for updating a user's fullname.
type UpdateFullnameRequest struct {
	Fullname string `json:"fullname" binding:"required"`
}

// UpdateAvatarRequest defines the structure for updating a user's avatar.
type UpdateAvatarRequest struct {
	AvatarURL string `json:"avatar_url" binding:"required,url"`
}

// CreateUserRequest defines the structure for creating a new user.
// This is typically used by an admin.
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	FullName string `json:"full_name"`
	Role     string `json:"role"` // e.g., "user", "admin"
}

type UpdateUserRequest struct {
	Username *string `json:"username"` // Changed from Name to Username
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByUsername(c context.Context, username string) (User, error) // Added GetByUsername
	GetByID(c context.Context, id int64) (User, error)
	Update(c context.Context, user *User) error // Added Update method
}

// UserUsecase represents the usecase for user related operations.
// It follows the 'user' naming convention, avoiding 'profile'.
type UserUsecase interface {
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByUsername(c context.Context, username string) (User, error)
	GetByID(c context.Context, id string) (User, error) // Changed id type from int64 to string
	Update(c context.Context, user *User) error
	InitiatePasswordReset(c context.Context, email string) error
	UpdateFullname(c context.Context, userID string, fullname string) (*User, error)
	UpdateAvatar(c context.Context, userID string, avatarURL string) (*User, error)
	List(c context.Context, page int, limit int) ([]User, error)
	Create(c context.Context, request *CreateUserRequest) (*User, error)
	Delete(c context.Context, userID string) error
}

const (
	// UserIDKey is the key for user ID in context
	UserIDKey = "userID"
)
