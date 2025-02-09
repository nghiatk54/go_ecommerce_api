package main

import (
	_ "github.com/nghiatk54/go_ecommerce_api/cmd/swagger/docs"
	"github.com/nghiatk54/go_ecommerce_api/internal/initialize"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           API Documentation Ecommerce Backend ShopDevGo
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  github.com/nghiatk54/go_ecommerce_api

// @contact.name   Team nghiatk54
// @contact.url    github.com/nghiatk54/go_ecommerce_api
// @contact.email  nghiatk54@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8002
// @BasePath  /v1/2025
// @schemes   http

func main() {
	r := initialize.Run()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8002")
}
