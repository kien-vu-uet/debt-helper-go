package controller

import (
	"net/http"
	"strconv" // Added for parsing query parameters

	"github.com/gin-gonic/gin"
	"github.com/kien-vu-uet/debt-helper-go/bootstrap"
	"github.com/kien-vu-uet/debt-helper-go/domain"
)

// UserController is the controller for user related operations
// It uses the 'user' naming convention, avoiding 'profile'.
type UserController struct {
	UserUsecase domain.UserUsecase
	Env         *bootstrap.Env
}

// NewUserController will create a new UserController
func NewUserController(userUsecase domain.UserUsecase, env *bootstrap.Env) *UserController {
	return &UserController{
		UserUsecase: userUsecase,
		Env:         env,
	}
}

// FetchUser godoc
// @Summary Fetch authenticated user's information
// @Description Get user details for the currently authenticated user
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} domain.User
// @Failure 401 {object} domain.ErrorResponse "User ID not found in token"
// @Failure 404 {object} domain.ErrorResponse "User not found"
// @Failure 500 {object} domain.ErrorResponse "User ID in token is not in expected format"
// @Router /me [get]
func (uc *UserController) FetchUser(c *gin.Context) {
	userID, ok := c.Get(domain.UserIDKey)
	if !ok {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User ID not found in token"})
		return
	}

	// Ensure userID is a string
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "User ID in token is not in expected format"})
		return
	}

	user, err := uc.UserUsecase.GetByID(c.Request.Context(), userIDStr)
	if err != nil {
		// Differentiate between not found and other errors if the usecase provides that distinction
		// For now, assuming any error from GetByID means user not found or other server error.
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// ResetPassword godoc
// @Summary Initiate password reset
// @Description Initiates the password reset process for a user by their email.
// @Tags users
// @Accept  json
// @Produce  json
// @Param credentials body domain.ResetPasswordRequest true "Email for password reset"
// @Success 200 {object} domain.SuccessResponse "Password reset initiated successfully"
// @Failure 400 {object} domain.ErrorResponse "Invalid request payload"
// @Failure 500 {object} domain.ErrorResponse "Failed to initiate password reset"
// @Router /reset-password [post]
func (uc *UserController) ResetPassword(c *gin.Context) {
	var request domain.ResetPasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid request: " + err.Error()})
		return
	}

	// Assuming UserUsecase has a method to handle password reset initiation
	err := uc.UserUsecase.InitiatePasswordReset(c.Request.Context(), request.Email)
	if err != nil {
		// Specific error handling can be added here if Usecase returns different error types
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Failed to initiate password reset: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Password reset initiated successfully. Please check your email."})
}

// UpdateFullname godoc
// @Summary Update authenticated user's fullname
// @Description Updates the fullname for the currently authenticated user.
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param fullnameInfo body domain.UpdateFullnameRequest true "Fullname information"
// @Success 200 {object} domain.User "Successfully updated fullname"
// @Failure 400 {object} domain.ErrorResponse "Invalid request payload"
// @Failure 401 {object} domain.ErrorResponse "User ID not found in token or invalid token"
// @Failure 500 {object} domain.ErrorResponse "Failed to update fullname or user ID in token is not in expected format"
// @Router /fullname [patch]
func (uc *UserController) UpdateFullname(c *gin.Context) {
	userID, ok := c.Get(domain.UserIDKey)
	if !ok {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User ID not found in token"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "User ID in token is not in expected format"})
		return
	}

	var request domain.UpdateFullnameRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid request: " + err.Error()})
		return
	}

	// Assuming UserUsecase has a method like UpdateFullname(ctx, userID, newFullname)
	updatedUser, err := uc.UserUsecase.UpdateFullname(c.Request.Context(), userIDStr, request.Fullname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Failed to update fullname: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// UpdateAvatar godoc
// @Summary Update authenticated user's avatar
// @Description Updates the avatar URL for the currently authenticated user.
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param avatarInfo body domain.UpdateAvatarRequest true "Avatar information"
// @Success 200 {object} domain.User "Successfully updated avatar"
// @Failure 400 {object} domain.ErrorResponse "Invalid request payload"
// @Failure 401 {object} domain.ErrorResponse "User ID not found in token or invalid token"
// @Failure 500 {object} domain.ErrorResponse "Failed to update avatar or user ID in token is not in expected format"
// @Router /avatar [patch]
func (uc *UserController) UpdateAvatar(c *gin.Context) {
	userID, ok := c.Get(domain.UserIDKey)
	if !ok {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User ID not found in token"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "User ID in token is not in expected format"})
		return
	}

	var request domain.UpdateAvatarRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid request: " + err.Error()})
		return
	}

	// Assuming UserUsecase has a method like UpdateAvatar(ctx, userID, newAvatarURL)
	updatedUser, err := uc.UserUsecase.UpdateAvatar(c.Request.Context(), userIDStr, request.AvatarURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Failed to update avatar: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// FetchUsers godoc
// @Summary Fetch a list of users
// @Description Retrieves a paginated list of users.
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param page query int false "Page number for pagination" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} domain.User "List of users"
// @Failure 500 {object} domain.ErrorResponse "Failed to fetch users"
// @Router /users [get]
func (uc *UserController) FetchUsers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	// Assuming UserUsecase has a method like List(ctx, page, limit)
	users, err := uc.UserUsecase.List(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Failed to fetch users: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Creates a new user. This is typically an admin-only action.
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param userInfo body domain.CreateUserRequest true "User creation information"
// @Success 201 {object} domain.User "Successfully created user"
// @Failure 400 {object} domain.ErrorResponse "Invalid request payload"
// @Failure 500 {object} domain.ErrorResponse "Failed to create user"
// @Router /users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var request domain.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid request: " + err.Error()})
		return
	}

	// Assuming UserUsecase has a method like Create(ctx, &request)
	createdUser, err := uc.UserUsecase.Create(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Failed to create user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Retrieves a specific user by their ID.
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_id path string true "User ID"
// @Success 200 {object} domain.User "User details"
// @Failure 400 {object} domain.ErrorResponse "Invalid user ID format"
// @Failure 404 {object} domain.ErrorResponse "User not found"
// @Failure 500 {object} domain.ErrorResponse "Failed to fetch user"
// @Router /users/{user_id} [get]
func (uc *UserController) GetUserByID(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "User ID is required"})
		return
	}

	user, err := uc.UserUsecase.GetByID(c.Request.Context(), userID)
	if err != nil {
		// This could be refined if GetByID returns specific error types (e.g., ErrNotFound)
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found or error fetching user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete user by ID
// @Description Deletes a specific user by their ID. This is typically an admin-only action.
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_id path string true "User ID"
// @Success 200 {object} domain.SuccessResponse "User deleted successfully"
// @Failure 400 {object} domain.ErrorResponse "Invalid user ID format"
// @Failure 404 {object} domain.ErrorResponse "User not found"
// @Failure 500 {object} domain.ErrorResponse "Failed to delete user"
// @Router /users/{user_id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "User ID is required"})
		return
	}

	// Assuming UserUsecase has a method like Delete(ctx, userID)
	err := uc.UserUsecase.Delete(c.Request.Context(), userID)
	if err != nil {
		// This could be refined if Delete returns specific error types
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Failed to delete user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "User deleted successfully"})
}
