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

const UP = 1
const RIGHT = 2
const DOWN = 3
const LEFT = 4

const OUTSIDE = 1
const ABOVE = 2
const BELOW = 3
const INSIDE = 4

func findPlacesInside(input string) int {
	defer timeTrack(time.Now(), "Part 2")
	lines := strings.Split(input, "\r\n")

	var posx, posy int
	sline := 0
	visited := make([][]bool, len(lines))

	for y, line := range lines {
		x := strings.Index(line, "S")
		if x != -1 {
			posx = x
			posy = y
			sline = y
		}
		visited[y] = make([]bool, len(line))
	}

	direction := 0
	north := false
	east := false
	south := false
	west := false

	if posy > 0 && strings.IndexAny(lines[posy-1][posx:posx+1], "|7F") != -1 {
		north = true
	}
	if posx < len(lines[posy]) && strings.IndexAny(lines[posy][posx+1:posx+2], "-7J") != -1 {
		east = true
	}
	if posy < len(lines) && strings.IndexAny(lines[posy+1][posx+1:posx+2], "|LJ") != -1 {
		south = true
	}
	if posx > 0 && strings.IndexAny(lines[posy][posx-1:posx], "-LF") != -1 {
		west = true
	}

	if north {
		direction = UP
	} else if east {
		direction = RIGHT
	} else if south {
		direction = DOWN
	} else {
		direction = LEFT
	}

	spiece := ""
	if north && east {
		spiece = "L"
	} else if north && south {
		spiece = "|"
	} else if north && west {
		spiece = "J"
	} else if south && west {
		spiece = "7"
	} else if south && east {
		spiece = "F"
	} else if east && west {
		spiece = "-"
	}

	stop := false
	for !stop {
		visited[posy][posx] = true
		switch direction {
		case UP:
			switch lines[posy-1][posx : posx+1] {
			case "S":
				stop = true
			case "|":
				posy--
			case "7":
				posy--
				direction = LEFT
			case "F":
				posy--
				direction = RIGHT
			}

		case RIGHT:
			switch lines[posy][posx+1 : posx+2] {
			case "S":
				stop = true
			case "-":
				posx++
			case "7":
				posx++
				direction = DOWN
			case "J":
				posx++
				direction = UP
			}
		case DOWN:
			switch lines[posy+1][posx : posx+1] {
			case "S":
				stop = true
			case "|":
				posy++
			case "L":
				posy++
				direction = RIGHT
			case "J":
				posy++
				direction = LEFT
			}
		case LEFT:
			switch lines[posy][posx-1 : posx] {
			case "S":
				stop = true
			case "-":
				posx--
			case "L":
				posx--
				direction = UP
			case "F":
				posx--
				direction = DOWN
			}
		}
	}

	lines[sline] = strings.Replace(lines[sline], "S", spiece, 1)

	count_inside := 0
	for y, line := range lines {
		inside := false
		above := false
		below := false
		for x := 0; x < len(line); x++ {
			if visited[y][x] {
				switch line[x : x+1] {
				case "|":
					inside = !inside
				case "L":
					below = true
				case "F":
					above = true
				case "J":
					if above {
						inside = !inside
						above = false
					} else {
						below = false
					}
				case "7":
					if below {
						inside = !inside
						below = false
					} else {
						above = false
					}
				}
			}
			if !visited[y][x] && inside {
				count_inside += 1
			}
		}
	}

	return count_inside
}

func findFarthestPoint(input string) int {
	defer timeTrack(time.Now(), "Part 1")
	lines := strings.Split(input, "\r\n")

	var posx, posy int

	for y, line := range lines {
		x := strings.Index(line, "S")
		if x != -1 {
			posx = x
			posy = y
			break
		}
	}

	direction := 0

	if posy > 0 && strings.IndexAny(lines[posy-1][posx:posx+1], "|7F") != -1 {
		direction = UP
	}
	if posx < len(lines[posy]) && strings.IndexAny(lines[posy][posx+1:posx+2], "-7J") != -1 {
		direction = RIGHT
	}
	if posy < len(lines) && strings.IndexAny(lines[posy+1][posx+1:posx+2], "|LJ") != -1 {
		direction = DOWN
	}
	if posx > 0 && strings.IndexAny(lines[posy][posx-1:posx], "-LF") != -1 {
		direction = LEFT
	}

	steps := 0
	for true {
		steps++

		switch direction {
		case UP:
			switch lines[posy-1][posx : posx+1] {
			case "S":
				return steps / 2
			case "|":
				posy--
			case "7":
				posy--
				direction = LEFT
			case "F":
				posy--
				direction = RIGHT
			}

		case RIGHT:
			switch lines[posy][posx+1 : posx+2] {
			case "S":
				return steps / 2
			case "-":
				posx++
			case "7":
				posx++
				direction = DOWN
			case "J":
				posx++
				direction = UP
			}
		case DOWN:
			switch lines[posy+1][posx : posx+1] {
			case "S":
				return steps / 2
			case "|":
				posy++
			case "L":
				posy++
				direction = RIGHT
			case "J":
				posy++
				direction = LEFT
			}
		case LEFT:
			switch lines[posy][posx-1 : posx] {
			case "S":
				return steps / 2
			case "-":
				posx--
			case "L":
				posx--
				direction = UP
			case "F":
				posx--
				direction = DOWN
			}
		}
	}

	return 0
}

func main() {
	defer timeTrack(time.Now(), "2023.10")
	bytes, err := os.ReadFile("2023.10.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("findFarthestPoint: %v\n", findFarthestPoint(string(bytes)))
	fmt.Printf("findPlacesInside: %v\n", findPlacesInside(string(bytes)))
}
