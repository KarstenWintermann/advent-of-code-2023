package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func findNumberOfWays(input string) int {
	lines := strings.Split(input, "\r\n")
	times := strings.Fields(lines[0])
	distances := strings.Fields(lines[1])
	result := 1

	for race := 1; race < len(times); race++ {
		ways := 0
		time, _ := strconv.Atoi(times[race])
		distance, _ := strconv.Atoi(distances[race])

		for duration := 1; duration < time; duration++ {
			travelled := (time - duration) * duration

			if travelled > distance {
				ways += 1
			}
		}

		result *= ways
	}

	return result
}

func findNumberOfWays2(input string) int {
	lines := strings.Split(input, "\r\n")
	times := strings.Fields(lines[0])
	distances := strings.Fields(lines[1])
	result := 1

	timestr := ""
	distancestr := ""

	for race := 1; race < len(times); race++ {
		timestr += times[race]
		distancestr += distances[race]
	}

	ways := 0
	time, _ := strconv.Atoi(timestr)
	distance, _ := strconv.Atoi(distancestr)

	for duration := 1; duration < time; duration++ {
		travelled := (time - duration) * duration
		if travelled > distance {
			ways += 1
		}
	}

	result *= ways

	return result
}

func main() {
	defer timeTrack(time.Now(), "2023.6")
	bytes, err := os.ReadFile("2023.6.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("numberOfWays: %v\n", findNumberOfWays(string(bytes)))
	fmt.Printf("numberOfWays2: %v\n", findNumberOfWays2(string(bytes)))
}
