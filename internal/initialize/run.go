package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitMysqlC()
	InitServiceInterface()
	InitRedis()
	InitKafkaProducer()
	r := InitRouter()
	r.Run(":8002")
}
