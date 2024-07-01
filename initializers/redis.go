package initializers

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client
var Ctx context.Context

func InitRedis() {
	Ctx = context.Background()
	opt := redis.Options{
		Addr:     "localhost:6379", // 这里填写你的Redis服务器地址
		Password: "",               // 如果有密码，填写密码
		DB:       0,                // 使用哪个数据库，通常填0                                                                                        // if you want to use default DB,set to 0
	}
	RedisDB = redis.NewClient(&opt)
}
