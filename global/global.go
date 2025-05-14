package global

import (
	"github.com/nghiatk54/goEcommerceApi/pkg/logger"
	"github.com/nghiatk54/goEcommerceApi/pkg/setting"
	"github.com/redis/go-redis/v9"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb *redis.Client
)