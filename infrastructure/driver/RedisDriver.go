package driver

import (
	"github.com/go-redis/redis/v7"
	"log"
	"os"
)

var (
	_address   = os.Getenv("REDIS_ADDRESS")
	_password  = os.Getenv("REDIS_PASSWORD")
	_defaultDB = 0
)

type RedisDriver struct {
	Client *redis.Client
}

func NewRedisDriver() *RedisDriver {
	result := &RedisDriver{}
	if err := result.Init(); err != nil {
		log.Fatal(err)
	}
	return result
}

func (d *RedisDriver) Init() error {
	d.Client = redis.NewClient(&redis.Options{
		Addr:     _address,
		Password: _password,
		DB:       _defaultDB,
	})
	if _, err := d.Client.Ping().Result(); err != nil {
		return err
	}
	return nil
}
