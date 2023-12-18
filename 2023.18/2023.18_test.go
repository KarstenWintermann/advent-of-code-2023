package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	bytes, err := os.ReadFile("2023.18.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := part1(string(bytes))
	if want != 62 {
		t.Fatalf("expected 62 but got %v", want)
	}
}

func TestPart2(t *testing.T) {
	bytes, err := os.ReadFile("2023.18.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := part2(string(bytes))
	if want != 952408144115 {
		t.Fatalf("expected 952408144115 but got %v", want)
	}
}
