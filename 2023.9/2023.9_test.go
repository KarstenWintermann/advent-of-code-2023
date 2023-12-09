package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFindPredictionSums(t *testing.T) {
	bytes, err := os.ReadFile("2023.9.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findPredictionsSum(string(bytes))
	if want != 114 {
		t.Fatalf("expected 114 but got %v", want)
	}
}

func TestFindPredictionSums2(t *testing.T) {
	bytes, err := os.ReadFile("2023.9.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findPredictionsSum2(string(bytes))
	if want != 2 {
		t.Fatalf("expected 2 but got %v", want)
	}
}
