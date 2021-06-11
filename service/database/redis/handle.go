package redis

import (
    "app/configs"
    "github.com/go-redis/redis/v8"
)

type handle struct {
    Cli *redis.Client
}

func ConnectRedis(c *configs.RedisConfig) (*handle, error) {
    rdb := redis.NewClient(&redis.Options{
        Addr:     c.Host + ":" + c.Port,
        Password: c.Pass,
        DB:       c.Name,
    })
    return &handle{
        Cli: rdb,
    }, nil
}
