package redis

import "C"
import (
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

type Redis_Client struct {
	Client *redis.Client
	sync.Once
}

type Config struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"DB"`
}

var Redis_client *Redis_Client

func NewClient() (*Redis_Client, error) {
	bytes, err := os.ReadFile("internal/config/redis.yaml")
	if err != nil {
		return nil, err
	}
	c := new(Config)
	err = yaml.Unmarshal(bytes, c)
	if err != nil {
		return nil, err
	}
	client := new(Redis_Client)

	client.Once.Do(func() {
		client.Client = redis.NewClient(&redis.Options{
			Addr:     c.Addr,
			Password: c.Password,
			DB:       c.DB,
		})
		Redis_client = client
	})
	return client, nil
}
