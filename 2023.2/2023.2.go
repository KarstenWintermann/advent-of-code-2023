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

func findCubePowers(input string) int {
	defer timeTrack(time.Now(), "Part 2")
	sumPowers := 0
	games := strings.Split(input, "\r\n")
	for _, game := range games {
		moves := strings.Split(game, ":")
		move := strings.Split(moves[1], ";")
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		for _, color := range move {
			colorMove := strings.Split(color, ",")
			for _, move := range colorMove {
				moves := strings.Split(strings.TrimSpace(move), " ")
				numColor, _ := strconv.Atoi(moves[0])
				colorName := moves[1]

				switch colorName {
				case "blue":
					if maxBlue < numColor {
						maxBlue = numColor
					}
				case "red":
					if maxRed < numColor {
						maxRed = numColor
					}
				case "green":
					if maxGreen < numColor {
						maxGreen = numColor
					}
				}
			}
		}
		sumPowers += maxBlue * maxRed * maxGreen
	}
	return sumPowers
}

func findPossibleGameIDs(input string) int {
	defer timeTrack(time.Now(), "Part 1")
	possibleGameIDs := 0
	games := strings.Split(input, "\r\n")
	for _, game := range games {
		skipGame := false
		moves := strings.Split(game, ":")
		gameId := strings.Split(moves[0], " ")
		gameNum, _ := strconv.Atoi(gameId[1])
		move := strings.Split(moves[1], ";")
		for _, color := range move {
			colorMove := strings.Split(color, ",")
			for _, move := range colorMove {
				moves := strings.Split(strings.TrimSpace(move), " ")
				numColor, _ := strconv.Atoi(moves[0])
				colorName := moves[1]

				switch {
				case colorName == "red" && numColor > 12:
					skipGame = true
				case colorName == "green" && numColor > 13:
					skipGame = true
				case colorName == "blue" && numColor > 14:
					skipGame = true
				}
			}
		}
		if !skipGame {
			possibleGameIDs += gameNum
		}
	}
	return possibleGameIDs
}

func main() {
	defer timeTrack(time.Now(), "2023.2")
	bytes, err := os.ReadFile("2023.2.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("possibleGameIDs: %v\n", findPossibleGameIDs(string(bytes)))
	fmt.Printf("sumPowers: %v\n", findCubePowers(string(bytes)))
}
