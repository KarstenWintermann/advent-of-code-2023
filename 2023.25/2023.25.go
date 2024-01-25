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

func part1(input string) int {
	lines := strings.Split(input, "\r\n")
	names := make(map[string]bool)
	graph := make(map[string][]string)
	for _, line := range lines {
		line = strings.ReplaceAll(line, ":", "")
		items := strings.Fields(line)
		graph[items[0]] = make([]string, 0)
		for n, item := range items {
			names[item] = true
			if n > 0 {
				graph[items[0]] = append(graph[items[0]], item)
			}
		}
	}

	for elem := range names {
		fmt.Printf("%v %v\n", elem, elem)
	}
	fmt.Println("#")

	for elem, edgelist := range graph {
		for _, edge := range edgelist {
			fmt.Printf("%v %v\n", elem, edge)
		}
	}

	return 0
}

func part2(input string) int {
	return 0
}

func main() {
	defer timeTrack(time.Now(), "2023.25")
	bytes, err := os.ReadFile("2023.25.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("part1: %v\n", part1(string(bytes)))
	fmt.Printf("part2: %v\n", part2(string(bytes)))
}
