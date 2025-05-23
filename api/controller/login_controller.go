package controller

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/kien-vu-uet/debt-helper-go/bootstrap"
	"github.com/kien-vu-uet/debt-helper-go/domain"
	// "golang.org/x/oauth2" // Example: Google OAuth2
	// "golang.org/x/oauth2/google" // Example: Google OAuth2
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

// Login godoc
// @Summary User login
// @Description User login with username and password.
// @Tags login
// @Accept  json
// @Produce  json
// @Param login body domain.LoginRequest true "Login Request"
// @Success 200 {object} domain.LoginResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 401 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Router /login [post]
// Login handles user login.
// It accepts a username and password, verifies them, and returns an access token if successful.
// It uses the LoginUsecase to interact with the user repository and handle token creation.
// The function also handles errors and returns appropriate HTTP status codes and messages.
// The function is documented using Swagger annotations for API documentation.
func (lc *LoginController) Login(c *gin.Context) {
	var request domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := lc.LoginUsecase.GetUserByUsername(c, request.Username) // Changed from GetUserByEmail
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given username"}) // Changed from email to username
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := domain.LoginResponse{
		Token: accessToken, // Changed from AccessToken and RefreshToken
	}

	c.JSON(http.StatusOK, loginResponse)
}


// ExtendAccessToken godoc
// @Summary Extend access token
// @Description Extends the validity of an existing access token.
// @Tags login
// @Accept  json
// @Produce  json
// @Param token body domain.ExtendAccessTokenRequest true "Extend Access Token Request"
// @Success 200 {object} domain.LoginResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 401 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /login/access-token/extend [post]
func (lc *LoginController) ExtendAccessToken(c *gin.Context) {
	var request domain.ExtendAccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID, err := lc.LoginUsecase.ExtractIDFromToken(request.Token, lc.Env.AccessTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid or expired token"})
		return
	}

	user, err := lc.LoginUsecase.GetUserByID(c, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found"})
		return
	}

	newAccessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.LoginResponse{Token: newAccessToken})
}

// RevokeRefreshToken godoc
// @Summary Revoke refresh token
// @Description Revokes an existing refresh token.
// @Tags login
// @Accept  json
// @Produce  json
// @Param token body domain.RevokeRefreshTokenRequest true "Revoke Refresh Token Request"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 401 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /login/refresh-token/revoke [post]
func (lc *LoginController) RevokeRefreshToken(c *gin.Context) {
	var request domain.RevokeRefreshTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := lc.LoginUsecase.RevokeRefreshToken(c, request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Failed to revoke token: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Refresh token revoked successfully"})
}

// VerifyUsername godoc
// @Summary Verify username
// @Description Verifies if a username is available.
// @Tags login
// @Accept  json
// @Produce  json
// @Param username body domain.VerifyUsernameRequest true "Verify Username Request"
// @Success 200 {object} domain.VerifyUsernameResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /login/verify-username [post]
func (lc *LoginController) VerifyUsername(c *gin.Context) {
	var request domain.VerifyUsernameRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err := lc.LoginUsecase.GetUserByUsername(c, request.Username)
	available := err != nil

	c.JSON(http.StatusOK, domain.VerifyUsernameResponse{Available: available})
}
