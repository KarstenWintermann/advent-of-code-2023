package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	bytes, err := os.ReadFile("2023.13.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findReflections(string(bytes))
	if want != 405 {
		t.Fatalf("expected 405 but got %v", want)
	}
}

func TestPart2(t *testing.T) {
	bytes, err := os.ReadFile("2023.13.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findSmudgedReflections(string(bytes))
	if want != 400 {
		t.Fatalf("expected 400 but got %v", want)
	}
}
