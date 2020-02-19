package db

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func initRedis() {
	// redis conn
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + config.Redis.Port,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
	logrus.Info("Redis connect successful.")
}
