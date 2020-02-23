package cache

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"strconv"
	"strings"
)

var client *redis.Client

func InitRedisCache() error {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func AddOnlineDevice(user int64, device int64) error {
	key := "online:" + string(user)
	_, err := client.Append(key, string(device)+",").Result()

	if err != nil {
		fmt.Printf("set online device failed, err:%v\n", err)
		return err
	}
	return nil
}

func GetAllOnlineDevice(user int64) ([]int64, error) {
	key := "online:" + string(user)
	val, err := client.Get(key).Result()

	if err != nil {
		fmt.Printf("get online device failed, err:%v\n", err)
		return nil, err
	}

	devices := strings.Split(val, ",")
	res := make([]int64, len(val))
	for _, d := range devices {
		i, _ := strconv.ParseInt(d, 10, 64)
		res = append(res, i)
	}
	return res, nil
}
