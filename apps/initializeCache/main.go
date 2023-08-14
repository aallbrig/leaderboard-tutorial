package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

type Game struct {
	Username string `json:"username"`
	Gamedate string `json:"gamedate"`
	Level    int    `json:"level"`
	Score    int    `json:"score"`
}

func main() {
	ctx := context.Background()

	filePath := "games.json"
	raw, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var games []Game
	err = json.Unmarshal(raw, &games)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	opt, err := redis.ParseURL("redis://" + os.Getenv("REDIS_HOSTNAME"))
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	client := redis.NewClient(opt)
	defer client.Close()

	for _, game := range games {
		gametime, err := time.Parse("2006-01-02T15:04:05", game.Gamedate)
		if err != nil {
			fmt.Println("Error parsing game date:", err)
			return
		}

		key := fmt.Sprintf("%s|%s|%s", game.Username, game.Gamedate, game.Level)
		pipe := client.TxPipeline()
		pipe.ZAdd(ctx, "Overall Leaderboard", &redis.Z{Score: float64(game.Score), Member: key})
		pipe.ZAdd(ctx, fmt.Sprintf("Monthly Leaderboard|%d-%d", gametime.Month(), gametime.Year()), &redis.Z{Score: float64(game.Score), Member: key})
		pipe.ZAdd(ctx, fmt.Sprintf("Daily Leaderboard|%d-%d-%d", gametime.Day(), gametime.Month(), gametime.Year()), &redis.Z{Score: float64(game.Score), Member: key})

		_, err = pipe.Exec(ctx)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	fmt.Println("Loaded data!")
}
