package initialize

import (
	"github.com/gin-gonic/gin"
	c "github.com/nghiatk54/go_ecommerce_api/internal/controller"
	"github.com/nghiatk54/go_ecommerce_api/internal/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// use middleware
	r.Use(middleware.AuthenMiddleware)
	v1 := r.Group("/v1/2025")
	{
		v1.GET("/user", c.NewUserController().GetUserById)
	}
	return r
}
