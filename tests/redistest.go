package tests

import (
	"github.com/revel/revel/testing"
	"github.com/memikequinn/lfs-server-go/app/services"
)

type RedisTest struct {
	testing.TestSuite
}

func (t *RedisTest) Before() {
	println("Set up")
}

func (t *RedisTest) TestRedisTestLoads() {
	t.AssertEqual(true,true)
}

func (t *RedisTest) TestDefaultRedisConfig() {
	config := &services.RedisConfig{}
	services.SetRedisConfig(config)
	t.AssertEqual("localhost:6379",config.Addr)
	t.AssertEqual("",config.Password)
	t.AssertEqual(0,config.DB)
}

func (t *RedisTest) TestRedisConfigWithArgs() {
	config := &services.RedisConfig{Addr: "somehost:6379"}
	services.SetRedisConfig(config)
	t.AssertEqual("somehost:6379",config.Addr)
	t.AssertEqual("",config.Password)
	t.AssertEqual(0,config.DB)
}

func (t *RedisTest) TestRedisNewClient() {
	config := &services.RedisConfig{}
	services.SetRedisConfig(config)
	client := services.NewRedisClient(config)
	r, err := client.Ping().Result()
	t.AssertEqual("PONG", r)
	t.AssertEqual(nil, err)
}


func (t *RedisTest) After() {
	println("Tear Down")
}
