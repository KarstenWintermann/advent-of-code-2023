package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	bytes, err := os.ReadFile("2023.15.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := hashSteps(string(bytes))
	if want != 1320 {
		t.Fatalf("expected 1320 but got %v", want)
	}
}

func TestPart2(t *testing.T) {
	bytes, err := os.ReadFile("2023.15.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := hashMap(string(bytes))
	if want != 145 {
		t.Fatalf("expected 145 but got %v", want)
	}
}
