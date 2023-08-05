package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"testing"
)

func TestRedisConnection(t *testing.T) {
	ctx := context.Background()

	hostname := os.Getenv("REDIS_HOSTNAME")
	port := os.Getenv("REDIS_PORT")
	redisURL := fmt.Sprintf("%s:%s", hostname, port)

	if redisURL == "" {
		t.Error("Expected non-empty Redis URL")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()

	if err != nil {
		t.Error(fmt.Print("Error when sending ping: ", err))
	} else {
		fmt.Println("Redis ping response:", pong)
	}
}

func main() {
	testing.Init()

	redisConnectionTest := testing.InternalTest{
		Name: "TestRedisConnection",
		F:    TestRedisConnection,
	}

	tests := []testing.InternalTest{
		redisConnectionTest,
	}

	matchFullName := func(pat, str string) (bool, error) { return true, nil }
	testing.Main(matchFullName, tests, nil, nil)
}
