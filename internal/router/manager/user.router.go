package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// private router use middleware: limiter, authentication, permission
	userRouterPrivate := router.Group("/admin/user")
	{
		userRouterPrivate.POST("/active_user")
	}
}
