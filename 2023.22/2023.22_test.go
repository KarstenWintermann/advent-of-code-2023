package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	bytes, err := os.ReadFile("2023.22.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := part1(string(bytes))
	if want != 5 {
		t.Fatalf("expected 5 but got %v", want)
	}
}

func TestPart2(t *testing.T) {
	bytes, err := os.ReadFile("2023.22.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := part2(string(bytes))
	if want != 7 {
		t.Fatalf("expected 7 but got %v", want)
	}
}
