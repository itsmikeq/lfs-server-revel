package services
import (
	"gopkg.in/redis.v3"
)

type RedisConfig struct {
	Addr string
	Password string
	DB int64
}

func SetRedisConfig(rc *RedisConfig) (*RedisConfig) {
	if len(rc.Addr) < 1 {rc.Addr = "localhost:6379"}
	if len(rc.Password) < 1 {rc.Password = ""}
	if rc.DB < 0 {rc.DB = 0}
	return rc
}

func NewRedisClient(rc *RedisConfig) (*redis.Client) {
	SetRedisConfig(rc)
	client := redis.NewClient(&redis.Options{Addr:rc.Addr, Password:rc.Password, DB: rc.DB})
	_, err := client.Ping().Result()
	perror(err)
	return client
}
