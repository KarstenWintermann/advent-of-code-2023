package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFindNumberOfWays(t *testing.T) {
	bytes, err := os.ReadFile("2023.6.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findNumberOfWays(string(bytes))
	if want != 288 {
		t.Fatalf("expected 288 but got %v", want)
	}
}

func TestFindNumberOfWays2(t *testing.T) {
	bytes, err := os.ReadFile("2023.6.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findNumberOfWays2(string(bytes))
	if want != 71503 {
		t.Fatalf("expected 71503 but got %v", want)
	}
}
