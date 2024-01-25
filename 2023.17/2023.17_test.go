package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	bytes, err := os.ReadFile("2023.17.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findPath(string(bytes))
	if want != 102 {
		t.Fatalf("expected 102 but got %v", want)
	}
}

func TestPart2(t *testing.T) {
	bytes, err := os.ReadFile("2023.17.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findPath2(string(bytes))
	if want != 94 {
		t.Fatalf("expected 94 but got %v", want)
	}
}
