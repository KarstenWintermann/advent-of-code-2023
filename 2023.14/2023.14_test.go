package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	bytes, err := os.ReadFile("2023.14.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findTotalLoad(string(bytes))
	if want != 136 {
		t.Fatalf("expected 136 but got %v", want)
	}
}

func TestPart2(t *testing.T) {
	bytes, err := os.ReadFile("2023.14.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findTotalLoad2(string(bytes))
	if want != 64 {
		t.Fatalf("expected 64 but got %v", want)
	}
}
