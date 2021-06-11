package configs

import (
    "os"
    "strconv"
)

type RedisConfig struct {
    Host string
    Port string
    Pass string
    Name int
}

func NewRedisConfig() *RedisConfig {
    db, _ := strconv.Atoi(os.Getenv("REDIS_NAME"))
    return &RedisConfig{
        Host: os.Getenv("REDIS_HOST"),
        Port: os.Getenv("REDIS_PORT"),
        Pass: "",
        Name: db,
    }
}

type MongoConfig struct {
    Host string
    Port string
    Name string
}

func NewMongoConfig() *MongoConfig {
    return &MongoConfig{
        Host: os.Getenv("MONGO_HOST"),
        Port: os.Getenv("MONGO_PORT"),
        Name: os.Getenv("MONGO_NAME"),
    }
}
