package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"web_app/settings"
)

// 声明一个全局的rdb变量
var (
	client *redis.Client
	Nil = redis.Nil
)

// 初始化连接

func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.DB, // use default DB
		PoolSize: cfg.PoolSize,
	})

	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func Close(){
	_ = client.Close()
}


