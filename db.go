package main

import (
	"github.com/go-redis/redis"
)

func SetIpBlocks(client *redis.Client, key string, value string)  {
	client.Set(key, value, 0)
}