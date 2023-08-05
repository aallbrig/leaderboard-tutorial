package main

import (
	"testing"
)

func TestSelectGames(t *testing.T) {
	result, err := executeQuery("SELECT * FROM games;")

	if err != nil {
		t.Fatalf("Error executing statement: %v", err)
	}

	if len(result.Records) == 0 {
		t.Fatalf("Expected record set above 0, but got 0")
	}
}
