package redis

import (
    "app/configs"
    "github.com/go-redis/redis/v8"
)

type handle struct {
    Conn *redis.Client
}

func ConnectRedis(c *configs.RedisConfig) (*handle, error) {
    cli := redis.NewClient(&redis.Options{
        Addr:     c.Host + ":" + c.Port,
        Password: c.Pass,
        DB:       c.Name,
    })
    return &handle{
        Conn: cli,
    }, nil
}
