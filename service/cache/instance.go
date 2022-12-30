package cache

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
)

type handle struct {
	Conn *redis.Client
}

func getHandle() (*handle, error) {
	host := os.Getenv("CACHE_HOST")
	port := os.Getenv("CACHE_PORT")
	name := "0"
	dsn := fmt.Sprintf("redis://%s:%s/%s", host, port, name)
	return ConnectRedis(dsn)
}

// ConnectRedis
// By default, the pool size is 10 connections per every available CPU as reported by runtime.GOMAXPROCS
// redis://<user>:<pass>@localhost:6379/<db>
// https://redis.uptrace.dev/guide/server.html#connecting-to-redis-server
// https://redis.uptrace.dev/guide/go-redis-debugging.html#connection-pool-size
func ConnectRedis(dsn string) (*handle, error) {

	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}
	cli := redis.NewClient(opt)
	return &handle{
		Conn: cli,
	}, nil
}

// Cluster
// https://redis.uptrace.dev/guide/go-redis-cluster.html
//cli := redis.NewClusterClient(&redis.ClusterOptions{
//	Addrs: []string{":7000", ":7001", ":7002"},
//})

// Cache
// https://redis.uptrace.dev/guide/go-redis-cache.html
