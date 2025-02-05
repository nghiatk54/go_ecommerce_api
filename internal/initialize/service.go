package initialize

import (
	"github.com/nghiatk54/go_ecommerce_api/global"
	"github.com/nghiatk54/go_ecommerce_api/internal/database"
	"github.com/nghiatk54/go_ecommerce_api/internal/service"
	"github.com/nghiatk54/go_ecommerce_api/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	// User service interface
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
}
