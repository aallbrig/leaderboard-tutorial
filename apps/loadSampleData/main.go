package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"os"
	"strings"
)

type Game struct {
	Username string
	Gamedate string
	Score    int
	Level    int
}

func executeQuery(query string) (output *rdsdataservice.ExecuteStatementOutput, err error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("DATABASE_REGION")),
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}
	svc := rdsdataservice.New(sess)

	params := &rdsdataservice.ExecuteStatementInput{
		ResourceArn: aws.String(os.Getenv("DATABASE_ARN")),
		SecretArn:   aws.String(os.Getenv("SECRET_ARN")),
		Database:    aws.String(os.Getenv("DATABASE_NAME")),
		Sql:         aws.String(query),
	}

	output, err = svc.ExecuteStatement(params)

	return
}

func main() {
	filePath := "games.json"
	raw, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var games []Game
	err = json.Unmarshal(raw, &games)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	var values []string
	for _, game := range games {
		values = append(values, fmt.Sprintf("('%s', '%s', %d, %d)", game.Username, game.Gamedate, game.Score, game.Level))
	}
	sql := `INSERT INTO games (username, gamedate, score, level) VALUES ` + strings.Join(values, ",\n")
	output, err := executeQuery(sql)
	if err != nil {
		fmt.Println("Error executing statement:", err)
	} else {
		fmt.Println("Games inserted successfully!")
		fmt.Println(output)
	}
}
