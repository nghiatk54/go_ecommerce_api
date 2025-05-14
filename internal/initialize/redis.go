package initialize

import (
	"context"
	"fmt"

	"github.com/nghiatk54/goEcommerceApi/global"
	"github.com/nghiatk54/goEcommerceApi/pkg/setting"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis(config setting.RedisSetting) {
	rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%v", config.Host, config.Port),
        Password: config.Password, // no password set
        DB:       config.Db,  // use default DB
		PoolSize: config.PoolSize,
    })

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis", zap.Error(err))
	}
	global.Logger.Info("Connected to Redis")
	global.Rdb = rdb
	redisExample()
}

func redisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		global.Logger.Error("Failed to set score", zap.Error(err))
		return
	}
	value, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		global.Logger.Error("Failed to get score", zap.Error(err))
		return
	}
	global.Logger.Info("Score", zap.String("score", value))
}