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

type pos_t struct {
	x, y int
}

func printmap(lines *[]string, visited map[pos_t]bool) {
	for y := 0; y < len((*lines)); y++ {
		for x := 0; x < len((*lines)[0]); x++ {
			if visited[pos_t{x, y}] {
				fmt.Print("0")
			} else {
				fmt.Print(string((*lines)[y][x]))
			}
		}
		fmt.Print("\n")
	}
}

func canstep(pos pos_t, visited map[pos_t]bool, lines *[]string) bool {
	return !visited[pos] &&
		pos.x >= 0 &&
		pos.x < len((*lines)[0]) &&
		pos.y >= 0 &&
		pos.y < len((*lines)) &&
		(*lines)[pos.y][pos.x] != '#'
}

func step(pos pos_t, visited map[pos_t]bool, lines *[]string) int {

	//printmap(lines, visited)

	if !canstep(pos_t{pos.x, pos.y}, visited, lines) {
		return 0
	}
	if pos.x == len((*lines)[0])-2 && pos.y == len(*lines)-1 {
		return 0
	}
	visited_new := make(map[pos_t]bool)
	for id := range visited {
		visited_new[id] = visited[id]
	}
	visited_new[pos] = true
	switch (*lines)[pos.y][pos.x] {
	case '<':
		return 1 + step(pos_t{pos.x - 1, pos.y}, visited_new, lines)
	case '>':
		return 1 + step(pos_t{pos.x + 1, pos.y}, visited_new, lines)
	case 'v':
		return 1 + step(pos_t{pos.x, pos.y + 1}, visited_new, lines)
	case '^':
		return 1 + step(pos_t{pos.x, pos.y - 1}, visited_new, lines)
	case '.':
		return 1 + max(
			step(pos_t{pos.x - 1, pos.y}, visited_new, lines),
			step(pos_t{pos.x + 1, pos.y}, visited_new, lines),
			step(pos_t{pos.x, pos.y + 1}, visited_new, lines),
			step(pos_t{pos.x, pos.y - 1}, visited_new, lines))
	}
	return 0
}

func setUnvisited(visited map[pos_t]bool, pos pos_t) {
	visited[pos] = false
}

func step2(pos pos_t, visited map[pos_t]bool, lines *[]string) int {

	if !canstep(pos_t{pos.x, pos.y}, visited, lines) {
		return 0
	}
	if pos.x == len((*lines)[0])-2 && pos.y == len(*lines)-1 {
		fmt.Printf(".")
		//printmap(lines, visited)
		return len(visited)
	}

	visited[pos] = true

	ret := max(
		step2(pos_t{pos.x - 1, pos.y}, visited, lines),
		step2(pos_t{pos.x + 1, pos.y}, visited, lines),
		step2(pos_t{pos.x, pos.y + 1}, visited, lines),
		step2(pos_t{pos.x, pos.y - 1}, visited, lines))

	delete(visited, pos)

	return ret
}

func part1(input string) int {
	lines := strings.Split(input, "\r\n")
	visited := make(map[pos_t]bool)
	return step(pos_t{1, 0}, visited, &lines)
}

func part2(input string) int {
	lines := strings.Split(input, "\r\n")
	visited := make(map[pos_t]bool)
	return step2(pos_t{1, 0}, visited, &lines)
}

func main() {
	defer timeTrack(time.Now(), "2023.23")
	bytes, err := os.ReadFile("2023.23.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("part1: %v\n", part1(string(bytes)))
	fmt.Printf("part2: %v\n", part2(string(bytes)))
}
