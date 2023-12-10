package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFindFarthestPoint(t *testing.T) {
	bytes, err := os.ReadFile("2023.10.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findFarthestPoint(string(bytes))
	if want != 8 {
		t.Fatalf("expected 8 but got %v", want)
	}
}

func TestFindPlacesInside(t *testing.T) {
	bytes, err := os.ReadFile("2023.10.example2.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findPlacesInside(string(bytes))
	if want != 10 {
		t.Fatalf("expected 10 but got %v", want)
	}
}
