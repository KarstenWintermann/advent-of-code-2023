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

type node struct {
	left, right string
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func findLength(startnode, directions string, tree map[string]node) int {
	steps := 0
	pos := startnode
	for i := 0; i <= len(directions); i++ {
		if i == len(directions) {
			i = 0
		}

		steps += 1
		switch directions[i : i+1] {
		case "L":
			pos = tree[pos].left
		case "R":
			pos = tree[pos].right
		}

		if strings.HasSuffix(pos, "Z") {
			break
		}
	}

	return steps
}

func findSteps2(input string) int {
	defer timeTrack(time.Now(), "Part 2")
	sections := strings.Split(input, "\r\n\r\n")
	directions := sections[0]
	tree := make(map[string]node)
	pos := make([]string, 0)
	for _, line := range strings.Split(sections[1], "\r\n") {
		line = strings.ReplaceAll(line, "=", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		line = strings.ReplaceAll(line, ",", "")
		fields := strings.Fields(line)
		tree[fields[0]] = node{fields[1], fields[2]}

		if strings.HasSuffix(fields[0], "A") {
			pos = append(pos, fields[0])
		}
	}

	result := 1
	for j := 0; j < len(pos); j++ {
		result = LCM(result, findLength(pos[j], directions, tree))
	}

	return result
}

func findSteps(input string) int {
	defer timeTrack(time.Now(), "Part 1")
	sections := strings.Split(input, "\r\n\r\n")
	directions := sections[0]
	tree := make(map[string]node)
	for _, line := range strings.Split(sections[1], "\r\n") {
		line = strings.ReplaceAll(line, "=", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		line = strings.ReplaceAll(line, ",", "")
		fields := strings.Fields(line)
		tree[fields[0]] = node{fields[1], fields[2]}
	}

	pos := "AAA"
	steps := 0
	for i := 0; i <= len(directions); i++ {
		if i == len(directions) {
			i = 0
		}

		steps += 1
		switch directions[i : i+1] {
		case "L":
			pos = tree[pos].left
		case "R":
			pos = tree[pos].right
		}

		if pos == "ZZZ" {
			break
		}
	}

	return steps
}

func main() {
	defer timeTrack(time.Now(), "2023.8")
	bytes, err := os.ReadFile("2023.8.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("findSteps: %v\n", findSteps(string(bytes)))
	fmt.Printf("findSteps2: %v\n", findSteps2(string(bytes)))
}
