package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFindAllDistances(t *testing.T) {
	bytes, err := os.ReadFile("2023.11.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findAllDistances(string(bytes))
	if want != 374 {
		t.Fatalf("expected 374 but got %v", want)
	}
}

func TestFindAllDistances2(t *testing.T) {
	bytes, err := os.ReadFile("2023.11.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findAllDistances2(string(bytes), 100)
	if want != 8410 {
		t.Fatalf("expected 8410 but got %v", want)
	}
}
