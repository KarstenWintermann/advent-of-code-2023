package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFindPossibleGameIDs(t *testing.T) {
	bytes, err := os.ReadFile("2023.2.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findPossibleGameIDs(string(bytes))
	if want != 8 {
		t.Fatalf("expected 8 but got %v", want)
	}
}

func TestFindCubePowers(t *testing.T) {
	bytes, err := os.ReadFile("2023.2.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findCubePowers(string(bytes))
	if want != 2286 {
		t.Fatalf("expected 2286 but got %v", want)
	}
}
