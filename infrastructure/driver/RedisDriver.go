package driver

import (
	"github.com/fun-dev/ccms-poc/infrastructure/config"
	"github.com/go-redis/redis/v7"
	"log"
)

var RedisDriverImpl *RedisDriver

func init() {
	RedisDriverImpl = &RedisDriver{}
	err := RedisDriverImpl.Init()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[debug] establised connection on RedisDriver.Init()\n")
}

type RedisDriver struct {
	Client *redis.Client
	Config *config.AppVariableOnRedis
}

func (d *RedisDriver) Init() error {
	d.Client = redis.NewClient(&redis.Options{
		Addr:     d.Config.Address,
		Password: d.Config.Password,
		DB:       d.Config.DBName,
	})
	_, err := d.Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
