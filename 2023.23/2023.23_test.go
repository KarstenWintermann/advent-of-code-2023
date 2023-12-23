package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	bytes, err := os.ReadFile("2023.23.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := part1(string(bytes))
	if want != 94 {
		t.Fatalf("expected 94 but got %v", want)
	}
}

func TestPart2(t *testing.T) {
	bytes, err := os.ReadFile("2023.23.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := part2(string(bytes))
	if want != 154 {
		t.Fatalf("expected 154 but got %v", want)
	}
}
