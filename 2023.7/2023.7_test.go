package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFindNumberOfWays(t *testing.T) {
	bytes, err := os.ReadFile("2023.7.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findWinnings(string(bytes))
	if want != 6440 {
		t.Fatalf("expected 6440 but got %v", want)
	}
}

func TestFindNumberOfWays2(t *testing.T) {
	bytes, err := os.ReadFile("2023.7.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findWinnings2(string(bytes))
	if want != 5905 {
		t.Fatalf("expected 5905 but got %v", want)
	}
}
