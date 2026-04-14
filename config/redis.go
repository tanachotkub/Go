// database/redis.go
package database

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

func ConnectRedis() error {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0, // ใช้ DB 0 เป็นค่าเริ่มต้น
	})

	// ทดสอบการเชื่อมต่อ (Ping)
	_, err := RedisClient.Ping(ctx).Result()
	return err
}
