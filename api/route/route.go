package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kien-vu-uet/debt-helper-go/api/middleware" // Corrected import path
	"github.com/kien-vu-uet/debt-helper-go/bootstrap"
	"gorm.io/gorm"
)

// Setup configures the routes for the application.
// It now accepts *gorm.DB instead of mongo.Database.
func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewLoginRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewUserRouter(env, timeout, db, protectedRouter)
}
