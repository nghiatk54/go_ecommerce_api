package initialize

import (
	"github.com/nghiatk54/goEcommerceApi/global"
	"github.com/nghiatk54/goEcommerceApi/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.LoggerSetting)
	global.Logger.Info("Init logger success!")
}