package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFindLowestLocation(t *testing.T) {
	bytes, err := os.ReadFile("2023.5.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findLowestLocation(string(bytes))
	if want != 35 {
		t.Fatalf("expected 35 but got %v", want)
	}
}

func TestFindLowestLocation2(t *testing.T) {
	bytes, err := os.ReadFile("2023.5.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findLowestLocation2(string(bytes))
	if want != 46 {
		t.Fatalf("expected 46 but got %v", want)
	}
}
