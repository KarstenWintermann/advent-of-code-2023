package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func findCopies(input string) int {
	defer timeTrack(time.Now(), "Part 2")
	lines := strings.Split(input, "\r\n")
	copies := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		copies[i] = 1
	}

	for i, line := range lines {
		leftright := strings.Split(line, "|")
		left := leftright[0]
		right := leftright[1]
		winning := strings.Split(strings.TrimSpace(left), ":")[1]
		winningnumbers := strings.Fields(winning)
		scratchcard := strings.Fields(right)
		linescore := 0
		for _, number := range scratchcard {
			if len(strings.TrimSpace(number)) > 0 && slices.Contains(winningnumbers, number) {
				linescore += 1
			}
		}

		for j := 1; j <= linescore; j++ {
			if i+j < len(lines) {
				copies[i+j] += copies[i]
			}
		}
	}

	sumCopies := 0
	for i := 0; i < len(lines); i++ {
		sumCopies += copies[i]
	}

	return sumCopies
}

func findScore(input string) int {
	defer timeTrack(time.Now(), "Part 1")
	lines := strings.Split(input, "\r\n")
	score := 0
	for _, line := range lines {
		leftright := strings.Split(line, "|")
		left := leftright[0]
		right := leftright[1]
		winning := strings.Split(strings.TrimSpace(left), ":")[1]
		winningnumbers := strings.Fields(winning)
		scratchcard := strings.Fields(right)
		linescore := 0
		for _, number := range scratchcard {
			if len(strings.TrimSpace(number)) > 0 && slices.Contains(winningnumbers, number) {
				if linescore == 0 {
					linescore = 1
				} else {
					linescore *= 2
				}
			}
		}
		score += linescore
	}
	return score
}

func main() {
	defer timeTrack(time.Now(), "2023.4")
	bytes, err := os.ReadFile("2023.4.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("score: %v\n", findScore(string(bytes)))
	fmt.Printf("copies: %v\n", findCopies(string(bytes)))
}
