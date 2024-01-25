package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	bytes, err := os.ReadFile("2023.21.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := part1(string(bytes), 6)
	if want != 16 {
		t.Fatalf("expected 16 but got %v", want)
	}
}

func TestPart2(t *testing.T) {
	bytes, err := os.ReadFile("2023.21.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := part2(string(bytes), 100)
	if want != 6536 {
		t.Fatalf("expected 6536 but got %v", want)
	}
}
