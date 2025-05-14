package initialize

import "github.com/nghiatk54/goEcommerceApi/global"

func Run() {
	LoadConfig()
	InitLogger()
	InitPostgres()
	InitRedis(global.Config.RedisSetting)
	r := InitRouter()
	r.Run(":" + global.Config.ServerSetting.Port)
}