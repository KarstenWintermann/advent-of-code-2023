package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSumPartNumbers(t *testing.T) {
	bytes, err := os.ReadFile("2023.3.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findSumPartNumbers(string(bytes))
	if want != 4361 {
		t.Fatalf("expected 4361 but got %v", want)
	}
}

func TestGearRatios(t *testing.T) {
	bytes, err := os.ReadFile("2023.3.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findGearRatios(string(bytes))
	if want != 467835 {
		t.Fatalf("expected 467835 but got %v", want)
	}
}
