package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kien-vu-uet/debt-helper-go/api/controller"
	"github.com/kien-vu-uet/debt-helper-go/api/middleware"
	"github.com/kien-vu-uet/debt-helper-go/bootstrap"
	"github.com/kien-vu-uet/debt-helper-go/domain"
	"github.com/kien-vu-uet/debt-helper-go/repository"
	"github.com/kien-vu-uet/debt-helper-go/usecase"
	"gorm.io/gorm"
)

// NewUserRouter creates a new user route
func NewUserRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	// Adhering to the 'user' naming convention, not 'profile'
	userUsecase := usecase.NewUserUsecase(ur, timeout)                     // Assumes usecase.NewUserUsecase exists
	signupUsecase := usecase.NewSignupUsecase(ur, timeout)                 // Create SignupUsecase
	userController := controller.NewUserController(userUsecase, env)       // Assumes controller.NewUserController exists
	signupController := controller.NewSignupController(signupUsecase, env) // Use signupUsecase

	// Routes for authenticated users
	group.GET("/me", middleware.JwtAuthMiddleware(env.AccessTokenSecret), userController.FetchUser)
	group.PATCH("/me/fullname", middleware.JwtAuthMiddleware(env.AccessTokenSecret), userController.UpdateFullname) // Changed route
	group.PATCH("/me/avatar", middleware.JwtAuthMiddleware(env.AccessTokenSecret), userController.UpdateAvatar)     // Changed route

	// Public routes or routes that might not require authentication initially
	// Signup is typically handled by its own controller (SignupController) and has its own route setup.
	// If UserController.Signup is a placeholder, ensure it's not the primary signup mechanism.
	group.POST("/signup", signupController.Signup)                    // This might be handled by SignupController
	group.POST("/users/reset-password", userController.ResetPassword) // Changed route to be more specific

	// Admin or specific permission routes for managing users
	// These routes should be protected by an appropriate authorization middleware (e.g., admin only)
	adminRoutes := group.Group("/users")                                 // Grouping user management routes
	adminRoutes.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret)) // Apply auth middleware, consider adding admin check middleware
	{
		adminRoutes.GET("/", userController.FetchUsers)
		adminRoutes.POST("/", userController.CreateUser) // Consider admin-only access
		adminRoutes.GET("/:user_id", userController.GetUserByID)
		adminRoutes.DELETE("/:user_id", userController.DeleteUser) // Consider admin-only access
	}
}
