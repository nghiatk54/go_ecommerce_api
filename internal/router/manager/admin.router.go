package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ar *AdminRouter) InitAdminRouter(router *gin.RouterGroup) {
	// public router
	adminRouterPublic := router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}
	// private router use middleware: limiter, authentication, permission
	adminRouterPrivate := router.Group("/admin")
	{
		adminRouterPrivate.POST("/active_user")
	}
}
