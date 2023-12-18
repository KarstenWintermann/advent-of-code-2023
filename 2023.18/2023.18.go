package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

const INSIDE = 0
const FENCE = 1
const OUTSIDE = 2

func floodfill(field *[][]int, pos_x int, pos_y int) {
	if pos_x < 0 ||
		pos_y < 0 ||
		pos_x >= len((*field)[0]) ||
		pos_y >= len(*field) {
		return
	}
	if (*field)[pos_y][pos_x] != 0 {
		return
	}
	(*field)[pos_y][pos_x] = OUTSIDE
	floodfill(field, pos_x+1, pos_y)
	floodfill(field, pos_x-1, pos_y)
	floodfill(field, pos_x, pos_y+1)
	floodfill(field, pos_x, pos_y-1)
}

type instruction_t struct {
	direction string
	steps     int
	color     int
}

func dimensions(instructions []instruction_t) (int, int, int, int) {
	max_x := 0
	max_y := 0

	min_x := 0
	min_y := 0

	pos_x := 0
	pos_y := 0

	for _, instruction := range instructions {
		switch instruction.direction {
		case "R":
			pos_x += instruction.steps
		case "D":
			pos_y += instruction.steps
		case "L":
			pos_x -= instruction.steps
		case "U":
			pos_y -= instruction.steps
		}

		max_x = max(max_x, pos_x)
		max_y = max(max_y, pos_y)
		min_x = min(min_x, pos_x)
		min_y = min(min_y, pos_y)
	}

	return max_x + 1, max_y + 1, min_x, min_y
}

type point_t struct {
	x, y int
}

func part2(input string) int {
	lines := strings.Split(input, "\r\n")
	instructions := make([]instruction_t, 0)
	for _, line := range lines {
		fields := strings.Fields(line)
		steps := fields[2]
		steps = strings.TrimPrefix(steps, "(#")
		steps = strings.TrimSuffix(steps, ")")
		n_steps, _ := strconv.ParseInt(steps[:5], 16, 32)
		color := 1
		direction := ""
		switch steps[5:] {
		case "0":
			direction = "R"
		case "1":
			direction = "D"
		case "2":
			direction = "L"
		case "3":
			direction = "U"
		}
		instructions = append(instructions, instruction_t{direction, int(n_steps), color})
	}

	x_values := make(map[int]bool, 0)
	y_values := make(map[int]bool, 0)

	max_x := 0
	max_y := 0

	min_x := 0
	min_y := 0

	pos_x := 0
	pos_y := 0

	x_values[0] = true
	y_values[0] = true
	for _, instruction := range instructions {
		switch instruction.direction {
		case "R":
			pos_x += instruction.steps
		case "D":
			pos_y += instruction.steps
		case "L":
			pos_x -= instruction.steps
		case "U":
			pos_y -= instruction.steps
		}
		x_values[pos_x] = true
		y_values[pos_y] = true

		max_x = max(max_x, pos_x)
		max_y = max(max_y, pos_y)
		min_x = min(min_x, pos_x)
		min_y = min(min_y, pos_y)
	}

	x_len := make([]int, 0)
	for key := range x_values {
		x_len = append(x_len, key-1)
		x_len = append(x_len, key)
		x_len = append(x_len, key+1)
	}
	sort.Ints(x_len)

	y_len := make([]int, 0)
	for key := range y_values {
		y_len = append(y_len, key-1)
		y_len = append(y_len, key)
		y_len = append(y_len, key+1)
	}
	sort.Ints(y_len)

	field := make([][]int, len(y_len)+2)
	for i := range field {
		field[i] = make([]int, len(x_len)+2)
	}

	real_x := 0
	real_y := 0

	field_x := 0
	field_y := 0

	for x_len[field_x] != real_x {
		field_x++
	}
	for y_len[field_y] != real_y {
		field_y++
	}
	field[field_y+1][field_x+1] = 1
	for _, instruction := range instructions {
		field[field_y+1][field_x+1] = FENCE
		switch instruction.direction {
		case "R":
			for x_len[field_x] != real_x+instruction.steps {
				field_x++
				field[field_y+1][field_x+1] = FENCE
			}
			real_x += instruction.steps
		case "D":
			for y_len[field_y] != real_y+instruction.steps {
				field_y++
				field[field_y+1][field_x+1] = FENCE
			}
			real_y += instruction.steps
		case "L":
			for x_len[field_x] != real_x-instruction.steps {
				field_x--
				field[field_y+1][field_x+1] = FENCE
			}
			real_x -= instruction.steps
		case "U":
			for y_len[field_y] != real_y-instruction.steps {
				field_y--
				field[field_y+1][field_x+1] = FENCE
			}
			real_y -= instruction.steps
		}
	}

	floodfill(&field, 0, 0)

	sum := 1
	for i := range field {
		for j := range field[i] {
			if field[i][j] != 2 {
				if field[i+1][j] != OUTSIDE &&
					field[i][j+1] != OUTSIDE &&
					field[i+1][j+1] != OUTSIDE {
					area := (x_len[j] - x_len[j-1]) * (y_len[i] - y_len[i-1])
					if area < 0 {
						area = -area
					}
					sum += area
				} else {
					if field[i+1][j] != OUTSIDE {
						sum += y_len[i] - y_len[i-1]
					}
					if field[i][j+1] != OUTSIDE {
						sum += x_len[j] - x_len[j-1]
					}
				}
			}
		}
	}

	return sum
}

func part1(input string) int {
	lines := strings.Split(input, "\r\n")
	instructions := make([]instruction_t, 0)
	for _, line := range lines {
		fields := strings.Fields(line)
		direction := fields[0]
		steps, _ := strconv.Atoi(fields[1])
		color := 1
		instructions = append(instructions, instruction_t{direction, steps, color})
	}

	max_x, max_y, min_x, min_y := dimensions(instructions)
	field := make([][]int, (max_y-min_y)+2)
	for i := range field {
		field[i] = make([]int, (max_x-min_x)+2)
	}

	pos_x := 1 - min_x
	pos_y := 1 - min_y
	field[pos_y][pos_x] = 1
	for _, instruction := range instructions {
		for i := 0; i < instruction.steps; i++ {
			switch instruction.direction {
			case "R":
				pos_x++
			case "D":
				pos_y++
			case "L":
				pos_x--
			case "U":
				pos_y--
			}
			field[pos_y][pos_x] = instruction.color
		}
	}

	floodfill(&field, 0, 0)

	sum := 0
	for i := range field {
		for j := range field[i] {
			if field[i][j] != OUTSIDE {
				sum += 1
			}
		}
	}

	return sum
}

func main() {
	defer timeTrack(time.Now(), "2023.18")
	bytes, err := os.ReadFile("2023.18.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("part1: %v\n", part1(string(bytes)))
	fmt.Printf("part2: %v\n", part2(string(bytes)))
}
