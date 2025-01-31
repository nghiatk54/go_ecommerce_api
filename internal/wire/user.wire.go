//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/nghiatk54/go_ecommerce_api/internal/controller"
	"github.com/nghiatk54/go_ecommerce_api/internal/repo"
	"github.com/nghiatk54/go_ecommerce_api/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepo,
		repo.NewUserAuthRepo,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}
