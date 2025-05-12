package initialize

import "github.com/nghiatk54/goEcommerceApi/global"

func Run() {
	LoadConfig()
	InitLogger()
	global.Logger.Info("Start server at port 8002")
	InitMysql()
	InitRedis()
	r := InitRouter()
	r.Run(":8002")
}