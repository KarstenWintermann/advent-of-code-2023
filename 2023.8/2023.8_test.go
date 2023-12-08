package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFindSteps(t *testing.T) {
	bytes, err := os.ReadFile("2023.8.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findSteps(string(bytes))
	if want != 2 {
		t.Fatalf("expected 2 but got %v", want)
	}
}

func TestFindSteps2(t *testing.T) {
	bytes, err := os.ReadFile("2023.8.example2.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findSteps2(string(bytes))
	if want != 6 {
		t.Fatalf("expected 6 but got %v", want)
	}
}
