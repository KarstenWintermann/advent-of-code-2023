package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

type pos_t struct {
	x, y, z int
}

type velocity_t struct {
	x, y, z int
}

type star_t struct {
	pos      pos_t
	velocity velocity_t
}

// taken from https://stackoverflow.com/questions/563198/how-do-you-detect-where-two-line-segments-intersect
// Returns 1 if the lines intersect, otherwise 0. In addition, if the lines
// intersect the intersection point may be stored in the floats i_x and i_y.
func get_line_intersectionXY(a, b, c, d pos_t) (x, y float64) {
	var s1_x, s1_y, s2_x, s2_y float64
	s1_x = float64(b.x) - float64(a.x)
	s1_y = float64(b.y) - float64(a.y)
	s2_x = float64(d.x) - float64(c.x)
	s2_y = float64(d.y) - float64(c.y)

	if -s2_x*s1_y+s1_x*s2_y == 0 {
		return math.MaxFloat64, math.MaxFloat64 // No collision
	}

	s := (-s1_y*(float64(a.x)-float64(c.x)) + s1_x*(float64(a.y)-float64(c.y))) / (-s2_x*s1_y + s1_x*s2_y)
	t := (s2_x*(float64(a.y)-float64(c.y)) - s2_y*(float64(a.x)-float64(c.x))) / (-s2_x*s1_y + s1_x*s2_y)

	if s < 0 || t < 0 {
		return math.MaxFloat64, math.MaxFloat64 // collision in the past
	}

	// Collision detected
	return float64(a.x) + (t * s1_x), float64(a.y) + (t * s1_y)
}

func collideXY(a star_t, b star_t, min_x int, max_x int, min_y int, max_y int) bool {
	x, y := get_line_intersectionXY(
		a.pos, pos_t{a.pos.x + a.velocity.x, a.pos.y + a.velocity.y, 0},
		b.pos, pos_t{b.pos.x + b.velocity.x, b.pos.y + b.velocity.y, 0})

	if x >= float64(min_x) &&
		x <= float64(max_x) &&
		y >= float64(min_y) &&
		y <= float64(max_y) {
		return true
	}

	return false
}

func part1(input string, min_x int, max_x int, min_y int, max_y int) int {
	lines := strings.Split(input, "\r\n")
	stars := make([]star_t, 0)

	for _, line := range lines {
		line = strings.ReplaceAll(line, "@", "")
		line = strings.ReplaceAll(line, ",", "")
		fields := strings.Fields(line)

		var star star_t
		star.pos.x, _ = strconv.Atoi(fields[0])
		star.pos.y, _ = strconv.Atoi(fields[1])
		star.pos.z, _ = strconv.Atoi(fields[2])
		star.velocity.x, _ = strconv.Atoi(fields[3])
		star.velocity.y, _ = strconv.Atoi(fields[4])
		star.velocity.z, _ = strconv.Atoi(fields[5])
		stars = append(stars, star)
	}

	sum := 0
	for a := 0; a < len(stars); a++ {
		for b := a + 1; b < len(stars); b++ {
			if collideXY(stars[a], stars[b], min_x, max_x, min_y, max_y) {
				sum++
			}
		}
	}

	return sum
}

func main() {
	defer timeTrack(time.Now(), "2023.24")
	bytes, err := os.ReadFile("2023.24.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("part1: %v\n", part1(string(bytes), 200000000000000, 400000000000000, 200000000000000, 400000000000000))
}
