package initialize

import (
	"context"
	"fmt"

	"github.com/nghiatk54/go_ecommerce_api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: r.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis", zap.Error(err))
	}
	fmt.Println("Initialize redis successfully")
	global.Logger.Info("Initialize redis successfully")
	global.Rdb = rdb
	redisExpamle()
}

func redisExpamle() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		fmt.Println("Failed to set score", zap.Error(err))
		return
	}
	value, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		fmt.Println("Failed to get score", zap.Error(err))
		return
	}
	fmt.Println("Value Score is::", value)
	global.Logger.Info("Value Score is::", zap.String("score", value))
}
