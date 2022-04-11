package db

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var Ctx context.Context

func init() {
	Ctx = context.Background()

	RDB = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: 0,
	})
}