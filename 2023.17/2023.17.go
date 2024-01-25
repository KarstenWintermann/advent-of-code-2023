package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

var lines []string
var bestHeatloss map[state_t]int
var bestSum int

type pos_t struct {
	x, y int
}

type vector_t struct {
	x, y int
}

type state_t struct {
	heading   vector_t
	pos       pos_t
	stepsleft int
}

var RIGHT = vector_t{1, 0}
var DOWN = vector_t{0, 1}
var LEFT = vector_t{-1, 0}
var UP = vector_t{0, -1}

func turnright(vec vector_t) vector_t {
	switch vec {
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	case UP:
		return RIGHT
	}

	return vector_t{0, 0}
}

func turnleft(vec vector_t) vector_t {
	switch vec {
	case RIGHT:
		return UP
	case DOWN:
		return RIGHT
	case LEFT:
		return DOWN
	case UP:
		return LEFT
	}

	return vector_t{0, 0}
}

func move(pos pos_t, direction vector_t) pos_t {
	return pos_t{pos.x + direction.x, pos.y + direction.y}
}

func isOutside(pos pos_t) bool {
	return (pos.x > len(lines[0])-1 ||
		pos.y > len(lines)-1 ||
		pos.x < 0 ||
		pos.y < 0)
}

func findHeatloss(pos pos_t) int {
	return int(lines[pos.y][pos.x] - '0')
}

func marker(direction vector_t) string {
	switch direction {
	case RIGHT:
		return ">"
	case DOWN:
		return "v"
	case LEFT:
		return "<"
	case UP:
		return "^"
	}
	return ""
}

func markerpos(pos pos_t) string {
	return string(lines[pos.y][pos.x])
}

func beststep(heatloss int, pos pos_t, direction vector_t, stepsLeft int) int {
	if isOutside(pos) {
		return math.MaxInt
	}

	heatloss += findHeatloss(pos)
	if bestSum != 0 && bestSum <= heatloss {
		return math.MaxInt
	}

	if pos.x == len(lines[0])-1 && pos.y == len(lines)-1 {
		if bestSum == 0 || bestSum > heatloss {
			bestSum = heatloss
			fmt.Printf("bestsum: %v\n", bestSum)
		}

		return heatloss
	}

	state := state_t{direction, pos, stepsLeft}
	existingHeatloss, foundExistingHeatloss := bestHeatloss[state]

	if foundExistingHeatloss && existingHeatloss <= heatloss {
		return math.MaxInt
	}

	bestHeatloss[state] = heatloss

	leftPath := beststep(heatloss, move(pos, turnleft(direction)), turnleft(direction), 2)
	rightPath := beststep(heatloss, move(pos, turnright(direction)), turnright(direction), 2)

	straightPath := math.MaxInt
	if stepsLeft > 0 {
		straightPath = beststep(heatloss, move(pos, direction), direction, stepsLeft-1)
	}

	return min(leftPath, straightPath, rightPath)
}

func beststep2(heatloss int, pos pos_t, direction vector_t, straightSteps int) int {
	if isOutside(pos) {
		return math.MaxInt
	}

	heatloss += findHeatloss(pos)
	if bestSum != 0 && bestSum <= heatloss {
		return math.MaxInt
	}

	if pos.x == len(lines[0])-1 && pos.y == len(lines)-1 {
		if straightSteps < 3 {
			return math.MaxInt
		}

		if bestSum == 0 || bestSum > heatloss {
			bestSum = heatloss
			fmt.Printf("bestsum: %v\n", bestSum)
		}

		return heatloss
	}

	state := state_t{direction, pos, straightSteps}
	existingHeatloss, foundExistingHeatloss := bestHeatloss[state]

	if foundExistingHeatloss && existingHeatloss <= heatloss {
		return math.MaxInt
	}

	bestHeatloss[state] = heatloss

	straightPath := math.MaxInt
	if straightSteps < 9 {
		straightPath = beststep2(heatloss, move(pos, direction), direction, straightSteps+1)
	}

	leftPath := math.MaxInt
	rightPath := math.MaxInt
	if straightSteps >= 3 {
		leftPath = beststep2(heatloss, move(pos, turnleft(direction)), turnleft(direction), 0)
		rightPath = beststep2(heatloss, move(pos, turnright(direction)), turnright(direction), 0)
	}

	return min(leftPath, straightPath, rightPath)
}

func findPath(input string) int {
	lines = strings.Split(input, "\r\n")
	bestHeatloss = make(map[state_t]int)
	bestSum = math.MaxInt

	start := pos_t{0, 0}

	ret := min(beststep(0, move(start, RIGHT), RIGHT, 2), beststep(0, move(start, DOWN), DOWN, 2))
	return ret
}

func findPath2(input string) int {
	lines = strings.Split(input, "\r\n")
	bestHeatloss = make(map[state_t]int)
	bestSum = math.MaxInt

	start := pos_t{0, 0}

	ret := min(beststep2(0, move(start, RIGHT), RIGHT, 0), beststep2(0, move(start, DOWN), DOWN, 0))
	return ret
}

func main() {
	defer timeTrack(time.Now(), "2023.17")
	bytes, err := os.ReadFile("2023.17.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	//fmt.Printf("part1: %v\n", findPath(string(bytes)))
	fmt.Printf("part2: %v\n", findPath2(string(bytes)))
}
