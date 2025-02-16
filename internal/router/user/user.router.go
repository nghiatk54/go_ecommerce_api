package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiatk54/go_ecommerce_api/internal/controller/account"
	"github.com/nghiatk54/go_ecommerce_api/internal/middleware"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// public router
	userRouterPublic := router.Group("/user")
	{
		userRouterPublic.POST("/register", account.Login.Register)
		userRouterPublic.POST("/verify_account", account.Login.VerifyOtp)
		userRouterPublic.POST("/update_password_register", account.Login.UpdatePasswordRegister)
		userRouterPublic.POST("/login", account.Login.Login)
	}
	// private router use middleware: limiter, authentication, permission
	userRouterPrivate := router.Group("/user")
	userRouterPrivate.Use(middleware.AuthenMiddleware)
	{
		userRouterPrivate.GET("/get_info")
		userRouterPrivate.POST("/two_factor/setup", account.TwoFa.SetUpTwoFactorAuth)
		userRouterPrivate.POST("/two_factor/verify", account.TwoFa.VerifyTwoFactorAuth)
	}
}
