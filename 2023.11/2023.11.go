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

func transpose(a [][]bool) [][]bool {
	newArr := make([][]bool, len(a[0]))
	for i := 0; i < len(a[0]); i++ {
		newArr[i] = make([]bool, len(a))
		for j := 0; j < len(a); j++ {
			newArr[i][j] = a[j][i]
		}
	}

	return newArr
}

type pos_t struct {
	x, y int
}

func distance(a, b pos_t) int {
	distance_x := b.x - a.x
	if distance_x < 0 {
		distance_x = -distance_x
	}
	distance_y := b.y - a.y
	if distance_y < 0 {
		distance_y = -distance_y
	}
	return distance_x + distance_y
}

func distance2(a pos_t, b pos_t, nonempty_lines []bool, nonempty_columns []bool, expansion_factor int) int {
	distance := 0
	lower_x := 0
	higher_x := 0
	if a.x > b.x {
		lower_x = b.x
		higher_x = a.x
	} else {
		lower_x = a.x
		higher_x = b.x
	}

	for i := lower_x; i < higher_x; i++ {
		if !nonempty_columns[i] {
			distance += expansion_factor
		} else {
			distance += 1
		}
	}

	lower_y := 0
	higher_y := 0
	if a.y > b.y {
		lower_y = b.y
		higher_y = a.y
	} else {
		lower_y = a.y
		higher_y = b.y
	}

	for i := lower_y; i < higher_y; i++ {
		if !nonempty_lines[i] {
			distance += expansion_factor
		} else {
			distance += 1
		}
	}

	return distance
}

func findAllDistances2(input string, expansion_factor int) int {
	lines := strings.Split(input, "\r\n")
	nonempty_lines := make([]bool, len(lines))
	nonempty_columns := make([]bool, len(lines[0]))
	stars := make([]pos_t, 0)
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			if line[j:j+1] == "#" {
				stars = append(stars, pos_t{j, i})
				nonempty_columns[j] = true
				nonempty_lines[i] = true
			}
		}
	}

	sum_distances := 0
	for i := 0; i < len(stars); i++ {
		for j := i + 1; j < len(stars); j++ {
			sum_distances += distance2(stars[i], stars[j], nonempty_lines, nonempty_columns, expansion_factor)
		}
	}

	return sum_distances
}

func findAllDistances(input string) int {
	lines := strings.Split(input, "\r\n")
	bitmap := make([][]bool, 0)
	i := 0
	for _, line := range lines {
		bitmap = append(bitmap, make([]bool, len(line)))
		all_empty := true
		for j := 0; j < len(line); j++ {
			if line[j:j+1] == "#" {
				bitmap[i][j] = true
				all_empty = false
			}
		}
		if all_empty {
			bitmap = append(bitmap, make([]bool, len(line)))
			i++
		}
		i++
	}

	bitmap2 := transpose(bitmap)
	bitmap = make([][]bool, 0)
	for _, line := range bitmap2 {
		all_empty := true
		for j := 0; j < len(line); j++ {
			if line[j] {
				all_empty = false
			}
		}
		if all_empty {
			bitmap = append(bitmap, line)
		}
		bitmap = append(bitmap, line)
	}

	bitmap = transpose(bitmap)

	pos_array := make([]pos_t, 0)
	for i, line := range bitmap {
		for j := 0; j < len(line); j++ {
			if line[j] {
				pos_array = append(pos_array, pos_t{j, i})
			}
		}
	}

	sum_distances := 0
	for i := 0; i < len(pos_array); i++ {
		for j := i + 1; j < len(pos_array); j++ {
			sum_distances += distance(pos_array[i], pos_array[j])
		}
	}

	return sum_distances
}

func main() {
	defer timeTrack(time.Now(), "2023.11")
	bytes, err := os.ReadFile("2023.11.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("test: %v\n", findAllDistances(string(bytes)))
	fmt.Printf("test: %v\n", findAllDistances2(string(bytes), 1000000))
}
