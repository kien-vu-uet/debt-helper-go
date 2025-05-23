package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kien-vu-uet/debt-helper-go/api/controller"
	"github.com/kien-vu-uet/debt-helper-go/bootstrap"
	"github.com/kien-vu-uet/debt-helper-go/domain"
	"github.com/kien-vu-uet/debt-helper-go/repository"
	"github.com/kien-vu-uet/debt-helper-go/usecase"
	"gorm.io/gorm"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login/access-token", lc.Login) // Changed from /login to /login/access-token
	group.POST("/login/access-token/extend", lc.ExtendAccessToken)
	group.POST("/login/refresh-token/revoke", lc.RevokeRefreshToken)
	group.POST("/login/verify-username", lc.VerifyUsername)
}
