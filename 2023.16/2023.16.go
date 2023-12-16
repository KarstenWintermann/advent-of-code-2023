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

const DOWN = 1
const UP = 2
const RIGHT = 4
const LEFT = 8

type beam_t struct {
	heading int
	x       int
	y       int
	dead    bool
}

func findCoverage(lines []string, start beam_t) int {
	visited := make([][]int, len(lines))
	for i := 0; i < len(lines); i++ {
		visited[i] = make([]int, len(lines[0]))
	}

	beams := make([]beam_t, 1)
	beams[0] = start

	for len(beams) > 0 {
		new_beams := make([]beam_t, 0)
		for beam_no, beam := range beams {
			visited[beam.y][beam.x] |= beam.heading
			switch lines[beam.y][beam.x] {
			case '|':
				switch beam.heading {
				case LEFT:
					beam.heading = DOWN
					new_beams = append(new_beams, beam_t{UP, beam.x, beam.y, false})
				case RIGHT:
					beam.heading = DOWN
					new_beams = append(new_beams, beam_t{UP, beam.x, beam.y, false})
				}
			case '-':
				switch beam.heading {
				case UP:
					beam.heading = LEFT
					new_beams = append(new_beams, beam_t{RIGHT, beam.x, beam.y, false})
				case DOWN:
					beam.heading = LEFT
					new_beams = append(new_beams, beam_t{RIGHT, beam.x, beam.y, false})
				}
			case '\\':
				switch beam.heading {
				case LEFT:
					beam.heading = UP
				case RIGHT:
					beam.heading = DOWN
				case UP:
					beam.heading = LEFT
				case DOWN:
					beam.heading = RIGHT
				}
			case '/':
				switch beam.heading {
				case LEFT:
					beam.heading = DOWN
				case RIGHT:
					beam.heading = UP
				case UP:
					beam.heading = RIGHT
				case DOWN:
					beam.heading = LEFT
				}
			}

			switch beam.heading {
			case UP:
				beam.y--
			case DOWN:
				beam.y++
			case LEFT:
				beam.x--
			case RIGHT:
				beam.x++
			}

			if beam.x < 0 ||
				beam.x >= len(lines[0]) ||
				beam.y < 0 ||
				beam.y >= len(lines) ||
				(visited[beam.y][beam.x]&beam.heading) != 0 {
				beams[beam_no].dead = true
			} else {
				beams[beam_no].x = beam.x
				beams[beam_no].y = beam.y
				beams[beam_no].heading = beam.heading
			}
		}

		beams = slices.DeleteFunc(beams, func(beam beam_t) bool {
			return beam.dead
		})
		for _, beam := range new_beams {
			beams = append(beams, beam)
		}
	}

	sum := 0
	for _, line := range visited {
		for _, pos := range line {
			if pos != 0 {
				sum++
			}
		}
	}

	return sum
}

func part2(input string) int {
	maxsum := 0
	lines := strings.Split(input, "\r\n")
	for i := 0; i < len(lines); i++ {
		sum := findCoverage(lines, beam_t{RIGHT, 0, i, false})
		if sum > maxsum {
			maxsum = sum
		}
		sum = findCoverage(lines, beam_t{LEFT, len(lines[0]) - 1, i, false})
		if sum > maxsum {
			maxsum = sum
		}
	}
	for i := 0; i < len(lines[0]); i++ {
		sum := findCoverage(lines, beam_t{DOWN, i, 0, false})
		if sum > maxsum {
			maxsum = sum
		}
		sum = findCoverage(lines, beam_t{UP, i, len(lines) - 1, false})
		if sum > maxsum {
			maxsum = sum
		}
	}
	return maxsum
}

func part1(input string) int {
	return findCoverage(strings.Split(input, "\r\n"), beam_t{RIGHT, 0, 0, false})
}

func main() {
	defer timeTrack(time.Now(), "2023.16")
	bytes, err := os.ReadFile("2023.16.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("part1: %v\n", part1(string(bytes)))
	fmt.Printf("part2: %v\n", part2(string(bytes)))
}
