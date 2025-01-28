package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiatk54/go_ecommerce_api/internal/wire"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// use dependency injection by wire
	userController, _ := wire.InitUserRouterHandler()
	// public router
	userRouterPublic := router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}
	// private router use middleware: limiter, authentication, permission
	userRouterPrivate := router.Group("/user")
	{
		userRouterPrivate.GET("/get_info")
	}
}
