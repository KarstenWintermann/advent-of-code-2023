package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFindCalib(t *testing.T) {
	bytes, err := os.ReadFile("2023.1.example.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findCalibSum(string(bytes))
	if want != 142 {
		t.Fatalf("expected 142 but got %v", want)
	}
}

func TestFindCalib2(t *testing.T) {
	bytes, err := os.ReadFile("2023.1.example2.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	want := findCalib2Sum(string(bytes))
	if want != 281 {
		t.Fatalf("expected 281 but got %v", want)
	}
}
