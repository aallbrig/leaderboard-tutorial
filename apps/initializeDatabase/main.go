package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"os"
)

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
	// Read the SQL query from the file
	sqlQuery, err := os.ReadFile("query.sql")
	if err != nil {
		fmt.Println("Error reading SQL file:", err)
		return
	}
	output, err := executeQuery(string(sqlQuery))
	if err != nil {
		fmt.Println("Error executing statement:", err)
	} else {
		fmt.Println("Table created successfully!")
		fmt.Println(output)
	}
}
