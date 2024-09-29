package arpc

import (
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var Redis map[int]*redis.Client

func InitRedis() {
	if viper.GetString("redis.addr") == "" && viper.GetBool("redis.disable") {
		return
	}

	Redis = make(map[int]*redis.Client)

	dbs := viper.GetIntSlice("redis.dbs")
	for _, db := range dbs {
		Redis[db] = redis.NewClient(&redis.Options{
			Addr:     viper.GetString("redis.addr"),
			Username: viper.GetString("redis.username"),
			Password: viper.GetString("redis.password"),
			DB:       db,
		})
		log.Printf("redis connect success: %d", db)
	}
}
