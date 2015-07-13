package models

import (
	"github.com/revel/revel"
	"github.com/memikequinn/lfs-server-go/app/services"
	"gopkg.in/redis.v3"
)

func init(){
	revel.OnAppStart(InitRedis)
}

var RedisClient *redis.Client
var InitRedis func() = func(){
	config := &services.RedisConfig{}
	services.SetRedisConfig(config)
	RedisClient = services.NewRedisClient(config)
}