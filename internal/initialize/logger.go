package initialize

import (
	"github.com/nghiatk54/go_ecommerce_api/global"
	"github.com/nghiatk54/go_ecommerce_api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
