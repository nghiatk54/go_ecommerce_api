package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nghiatk54/go_ecommerce_api/global"
	"github.com/redis/go-redis/v9"
)

func GetCache(ctx context.Context, key string, obj interface{}) error {
	rs, err := global.Rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return fmt.Errorf("key %s not found", key)
	} else if err != nil {
		return fmt.Errorf("get cache failed %v", err)
	}
	// convert rs json to obj
	if err := json.Unmarshal([]byte(rs), obj); err != nil {
		return fmt.Errorf("convert json to obj failed %v", err)
	}
	return nil
}
