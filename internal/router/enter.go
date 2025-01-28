package router

import (
	"github.com/nghiatk54/go_ecommerce_api/internal/router/manager"
	"github.com/nghiatk54/go_ecommerce_api/internal/router/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
