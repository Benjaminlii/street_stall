package drivers

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"street_stall/biz/config"
)

var (
	RedisClient *redis.Client
)

func InitRedis(config *config.BasicConfig) {
	addr := fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port)
	password := config.Redis.PassWorld
	db, err := strconv.Atoi(config.Redis.DB)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Print("[system][redis] redis client connection wrong, err:", err)
	}
	log.Print("[system][redis] redis connection success, pong:", pong)
	RedisClient = client
}
