package main

import (
	"fmt"
	"testing"
)

func TestExecuteQuery(t *testing.T) {
	query := "SELECT * FROM games;"
	output, err := executeQuery(query)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
		return
	}
	fmt.Println(output)
}
