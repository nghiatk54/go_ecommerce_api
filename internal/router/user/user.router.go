package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiatk54/go_ecommerce_api/internal/controller/account"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// public router
	userRouterPublic := router.Group("/user")
	{
		userRouterPublic.POST("/register", account.Login.Register)
		userRouterPublic.POST("/login", account.Login.Login)
	}
	// private router use middleware: limiter, authentication, permission
	userRouterPrivate := router.Group("/user")
	{
		userRouterPrivate.GET("/get_info")
	}
}
