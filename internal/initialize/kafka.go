package initialize

import (
	"fmt"
	"log"

	"github.com/nghiatk54/go_ecommerce_api/global"
	kafka "github.com/segmentio/kafka-go"
)

// init kafka producer
var KafkaProducer *kafka.Writer

func InitKafkaProducer() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP(fmt.Sprintf("%s:%d", global.Config.Kafka.Host, global.Config.Kafka.Port)),
		Topic:    global.Config.Kafka.Topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// Close kafka producer
func CloseKafkaProducer() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatalf("Failed to close kafka producer: %v", err)
	}
}
