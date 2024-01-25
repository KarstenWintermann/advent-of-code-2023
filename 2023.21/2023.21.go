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

func printMap(lines []string, places map[pos_t]bool) {
	printlines := make([]string, len(lines))
	for i, line := range lines {
		printlines[i] = line
	}
	for pos := range places {
		printlines[pos.y] = printlines[pos.y][:pos.x] + "O" + printlines[pos.y][pos.x+1:]
	}
	for _, line := range printlines {
		fmt.Println(line)
	}
}

func setInTile1(tile1 []string, pos pos_t) {
	if pos.x >= 0 &&
		pos.x < len(tile1[0]) &&
		pos.y >= 0 &&
		pos.y < len(tile1) {
		tile1[pos.y] = tile1[pos.y][:pos.x] + "0" + tile1[pos.y][pos.x+1:]
	}
}

type pos_t struct {
	x, y int
}

func positive_modulo(i, n int) int {
	return (i%n + n) % n
}

func part1(input string, days int) int {
	lines := strings.Split(input, "\r\n")
	places := make(map[pos_t]bool)
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if line[x] == 'S' {
				places[pos_t{x, y}] = true
			}
		}
	}

	for day := 0; day < days; day++ {
		new_places := make(map[pos_t]bool)
		for pos := range places {
			if pos.x > 0 {
				if lines[pos.y][pos.x-1] != '#' {
					new_places[pos_t{pos.x - 1, pos.y}] = true
				}
			}
			if pos.y > 0 {
				if lines[pos.y-1][pos.x] != '#' {
					new_places[pos_t{pos.x, pos.y - 1}] = true
				}
			}
			if pos.x+1 < len(lines[0]) {
				if lines[pos.y][pos.x+1] != '#' {
					new_places[pos_t{pos.x + 1, pos.y}] = true
				}
			}
			if pos.y+1 < len(lines) {
				if lines[pos.y+1][pos.x] != '#' {
					new_places[pos_t{pos.x, pos.y + 1}] = true
				}
			}
		}
		places = new_places
	}

	return len(places)
}

func part2(input string, days int) int {
	lines := strings.Split(input, "\r\n")
	places := make(map[pos_t]bool)
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if line[x] == 'S' {
				places[pos_t{x, y}] = true
			}
		}
	}

	for day := 0; day < days; day++ {
		new_places := make(map[pos_t]bool)
		tile1 := make([]string, len(lines))
		for i := range tile1 {
			tile1[i] = strings.Repeat(" ", len(lines[0]))
		}
		for pos := range places {
			if lines[positive_modulo(pos.y, len(lines))][positive_modulo(pos.x-1, len(lines[0]))] != '#' {
				new_places[pos_t{pos.x - 1, pos.y}] = true
				setInTile1(tile1, pos_t{pos.x - 1, pos.y})
			}
			if lines[positive_modulo(pos.y-1, len(lines))][positive_modulo(pos.x, len(lines[0]))] != '#' {
				new_places[pos_t{pos.x, pos.y - 1}] = true
				setInTile1(tile1, pos_t{pos.x, pos.y - 1})
			}
			if lines[positive_modulo(pos.y, len(lines))][positive_modulo(pos.x+1, len(lines[0]))] != '#' {
				new_places[pos_t{pos.x + 1, pos.y}] = true
				setInTile1(tile1, pos_t{pos.x + 1, pos.y})
			}
			if lines[positive_modulo(pos.y+1, len(lines))][positive_modulo(pos.x, len(lines[0]))] != '#' {
				new_places[pos_t{pos.x, pos.y + 1}] = true
				setInTile1(tile1, pos_t{pos.x, pos.y + 1})
			}
		}
		places = new_places
	}

	return len(places)
}

func calc(in float64) float64 {
	var res float64

	res = (in * in * float64(3936288)) / float64(4496182)

	res = res + (in*float64(7025268))/float64(4496182)

	res = res + float64(60978928)/float64(4496182)

	return res
}

func main() {
	defer timeTrack(time.Now(), "2023.21")
	bytes, err := os.ReadFile("2023.21.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("part1: %v\n", part1(string(bytes), 64))
}
