package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	bytes, err := os.ReadFile("2023.12.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findArrangements(string(bytes))
	if want != 21 {
		t.Fatalf("expected 21 but got %v", want)
	}
}

func TestPart2(t *testing.T) {
	bytes, err := os.ReadFile("2023.12.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findArrangements3(string(bytes))
	if want != 525152 {
		t.Fatalf("expected 525152 but got %v", want)
	}
}
