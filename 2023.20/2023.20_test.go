package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	bytes, err := os.ReadFile("2023.19.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := part1(string(bytes))
	if want != 19114 {
		t.Fatalf("expected 19114 but got %v", want)
	}
}

func TestPart2(t *testing.T) {
	bytes, err := os.ReadFile("2023.19.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := part2(string(bytes))
	if want != 167409079868000 {
		t.Fatalf("expected 167409079868000 but got %v", want)
	}
}
