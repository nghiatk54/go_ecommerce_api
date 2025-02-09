package initialize

import "github.com/gin-gonic/gin"

func Run() *gin.Engine {
	LoadConfig()
	InitLogger()
	InitMysqlC()
	InitServiceInterface()
	InitRedis()
	InitKafkaProducer()
	r := InitRouter()
	return r
}
