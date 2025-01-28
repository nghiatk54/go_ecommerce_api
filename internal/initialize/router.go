package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiatk54/go_ecommerce_api/global"
	"github.com/nghiatk54/go_ecommerce_api/internal/router"
)

func InitRouter() *gin.Engine {
	// Set gin mode environment
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// Set middleware: logger, cross, limiter
	// r.Use()
	// r.Use()
	// r.Use()
	// Initialize router
	managerRouter := router.RouterGroupApp.Manager
	userRouter := router.RouterGroupApp.User
	mainGroup := r.Group("/v1/2025")
	{
		mainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitUserRouter(mainGroup)
		userRouter.InitProductRouter(mainGroup)
	}
	{
		managerRouter.InitUserRouter(mainGroup)
		managerRouter.InitAdminRouter(mainGroup)
	}

	return r
}
