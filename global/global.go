package global

import (
	"github.com/nghiatk54/go_ecommerce_api/pkg/logger"
	"github.com/nghiatk54/go_ecommerce_api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
