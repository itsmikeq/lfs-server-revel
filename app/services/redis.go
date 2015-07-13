package services
import (
	"gopkg.in/redis.v3"
	"github.com/revel/revel"
	"strconv"
	"fmt"
)

type RedisConfig struct {
	Addr string
	Password string
	DB int64
}

func SetRedisConfig(rc *RedisConfig) (*RedisConfig) {
	if len(rc.Addr) < 1 {rc.Addr, _ = revel.Config.String("redis.server")}
	if len(rc.Password) < 1 {rc.Password, _ = revel.Config.String("redis.password")}
	// What a nightmare
	_db, _ := revel.Config.String("redis.db")
	fmt.Println("DB", _db)
	db, _ := strconv.ParseInt(_db, 0, 0)
	if rc.DB < 0 {rc.DB =  db}
	return rc
}

func NewRedisClient(rc *RedisConfig) (*redis.Client) {
	SetRedisConfig(rc)
	client := redis.NewClient(&redis.Options{Addr:rc.Addr, Password:rc.Password, DB: rc.DB})
	_, err := client.Ping().Result()
	perror(err)
	return client
}
