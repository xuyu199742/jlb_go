package boot

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/xinzf/gokit/logger"
	"jlb_shop_go/global"
)

func Redis() {
	redisCfg := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.DefaultLogger.Error("redis connect ping failed, err:", err)
	} else {
		logger.DefaultLogger.Info("redis connect ping response:", pong)
		global.Redis = client
	}
}
