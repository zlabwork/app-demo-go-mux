package redis

import (
	"app"
	"github.com/go-redis/redis/v8"
)

type handle struct {
	Conn *redis.Client
}

func ConnectRedis(c *app.Yaml) (*handle, error) {
	cli := redis.NewClient(&redis.Options{
		Addr: c.Redis.Host + ":" + c.Redis.Port,
		// Password: "",
		// DB:       1,
	})
	return &handle{
		Conn: cli,
	}, nil
}
