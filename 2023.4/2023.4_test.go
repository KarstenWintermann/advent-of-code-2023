package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFindScore(t *testing.T) {
	bytes, err := os.ReadFile("2023.4.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findScore(string(bytes))
	if want != 13 {
		t.Fatalf("expected 13 but got %v", want)
	}
}

func TestFindCopies(t *testing.T) {
	bytes, err := os.ReadFile("2023.4.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findCopies(string(bytes))
	if want != 30 {
		t.Fatalf("expected 30 but got %v", want)
	}
}
