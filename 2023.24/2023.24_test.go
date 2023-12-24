package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	bytes, err := os.ReadFile("2023.24.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := part1(string(bytes), 7, 27, 7, 27)
	if want != 2 {
		t.Fatalf("expected 2 but got %v", want)
	}
}
