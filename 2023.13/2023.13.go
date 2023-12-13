package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func transpose(a []string) []string {
	newArr := make([]string, len(a[0]))
	for i := 0; i < len(a[0]); i++ {
		for j := 0; j < len(a); j++ {
			newArr[i] += a[j][i : i+1]
		}
	}

	return newArr
}

func findReflection(lines []string) int {
	for i, line := range lines {
		if len(lines) > i+1 && line == lines[i+1] {
			for j := 1; j <= len(lines); j++ {
				if i+1+j >= len(lines) || i-j < 0 {
					return i + 1
				}
				if lines[i+1+j] != lines[i-j] {
					break
				}
			}
		}
	}
	return 0
}

func difference(a, b string) int {
	difference := 0
	for i := 0; i < len(a); i++ {
		if a[i:i+1] != b[i:i+1] {
			difference++
		}
	}
	return difference
}

func findSmudgedReflection(lines []string) int {
	for i, _ := range lines {
		smudge := 0
		for j := 0; j <= len(lines); j++ {
			if i+j+1 >= len(lines) || i-j < 0 {
				if smudge == 1 {
					return i + 1
				}
				break
			}
			smudge += difference(lines[i+j+1], lines[i-j])
		}
	}
	return 0
}

func findReflections(input string) int {
	defer timeTrack(time.Now(), "Part 1")
	sections := strings.Split(input, "\r\n\r\n")
	sum := 0
	for _, section := range sections {
		lines := strings.Split(section, "\r\n")
		sum += findReflection(lines) * 100
		sum += findReflection(transpose(lines))
	}
	return sum
}

func findSmudgedReflections(input string) int {
	defer timeTrack(time.Now(), "Part 2")
	sections := strings.Split(input, "\r\n\r\n")
	sum := 0
	for _, section := range sections {
		lines := strings.Split(section, "\r\n")
		sum += findSmudgedReflection(lines) * 100
		sum += findSmudgedReflection(transpose(lines))
	}
	return sum
}

func main() {
	defer timeTrack(time.Now(), "2023.13")
	bytes, err := os.ReadFile("2023.13.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("findReflections: %v\n", findReflections(string(bytes)))
	fmt.Printf("findSmudgedReflections: %v\n", findSmudgedReflections(string(bytes)))
}
