package global

import (
	"database/sql"

	"github.com/nghiatk54/go_ecommerce_api/pkg/logger"
	"github.com/nghiatk54/go_ecommerce_api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	Rdb           *redis.Client
	Mdbc          *sql.DB
	KafkaProducer *kafka.Writer
)
