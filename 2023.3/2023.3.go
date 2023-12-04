package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func containsSymbol(c rune) bool {
	return !unicode.IsDigit(c) && c != rune('.')
}

func findGear(c rune) bool {
	return c == rune('*')
}

type pos_t struct {
	line int
	col  int
}

func findGearRatios(input string) int {
	lines := strings.Split(input, "\r\n")
	gears := make(map[pos_t][]int)
	for i, line := range lines {
		j := 0
		inNumber := false
		firstDigit := 0
		for j < len(line) {
			switch {
			case !inNumber && unicode.IsDigit([]rune(line)[j]):
				firstDigit = j
				inNumber = true
			case inNumber && (!unicode.IsDigit([]rune(line)[j]) || j == len(line)-1):
				lastDigit := j
				if lastDigit == len(line)-1 && unicode.IsDigit([]rune(line)[j]) {
					lastDigit = len(line)
				}
				number, _ := strconv.Atoi(line[firstDigit:lastDigit])
				if firstDigit-1 > 0 && line[firstDigit-1:firstDigit] == "*" {
					gears[pos_t{i, firstDigit - 1}] = append(gears[pos_t{i, firstDigit - 1}], number)
				}
				if j < len(line) && line[j:j+1] == "*" {
					gears[pos_t{i, j}] = append(gears[pos_t{i, j}], number)
				}
				correctLeft := firstDigit - 1
				if correctLeft < 0 {
					correctLeft = 0
				}
				correctRight := j + 1
				if correctRight > len(line) {
					correctRight = len(line)
				}
				if i > 0 {
					gearpos := strings.IndexFunc(lines[i-1][correctLeft:correctRight], findGear)
					if gearpos != -1 {
						gears[pos_t{i - 1, correctLeft + gearpos}] = append(gears[pos_t{i - 1, correctLeft + gearpos}], number)
					}
				}
				if (i + 1) < len(lines) {
					gearpos := strings.IndexFunc(lines[i+1][correctLeft:correctRight], findGear)
					if gearpos != -1 {
						gears[pos_t{i + 1, correctLeft + gearpos}] = append(gears[pos_t{i + 1, correctLeft + gearpos}], number)
					}
				}
				inNumber = false
			}
			j++
		}
	}

	gearRatios := 0
	for _, numbers := range gears {
		if len(numbers) == 2 {
			gearRatios += numbers[0] * numbers[1]
		}
	}

	return gearRatios
}

func findSumPartNumbers(input string) int {
	lines := strings.Split(input, "\r\n")
	sumPartNumbers := 0
	for i, line := range lines {
		j := 0
		inNumber := false
		firstDigit := 0
		for j < len(line) {
			switch {
			case !inNumber && unicode.IsDigit([]rune(line)[j]):
				firstDigit = j
				inNumber = true
			case inNumber && (!unicode.IsDigit([]rune(line)[j]) || j == len(line)-1):
				lastDigit := j
				if lastDigit == len(line)-1 && unicode.IsDigit([]rune(line)[j]) {
					lastDigit = len(line)
				}
				number, _ := strconv.Atoi(line[firstDigit:lastDigit])
				foundSymbol := false
				if firstDigit-1 > 0 && strings.IndexFunc(line[firstDigit-1:firstDigit], containsSymbol) != -1 {
					foundSymbol = true
				}
				if j < len(line) && strings.IndexFunc(line[j:j+1], containsSymbol) != -1 {
					foundSymbol = true
				}
				correctLeft := firstDigit - 1
				if correctLeft < 0 {
					correctLeft = 0
				}
				correctRight := j + 1
				if correctRight > len(line) {
					correctRight = len(line)
				}
				if i > 0 && strings.IndexFunc(lines[i-1][correctLeft:correctRight], containsSymbol) != -1 {
					foundSymbol = true
				}
				if (i+1) < len(lines) && strings.IndexFunc(lines[i+1][correctLeft:correctRight], containsSymbol) != -1 {
					foundSymbol = true
				}
				if foundSymbol {
					sumPartNumbers += number
				}
				inNumber = false
			}
			j++
		}
	}

	return sumPartNumbers
}

func main() {
	defer timeTrack(time.Now(), "2023.3")
	bytes, err := os.ReadFile("2023.3.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("sumPartNumbers: %v\n", findSumPartNumbers(string(bytes)))
	fmt.Printf("gearRatios: %v\n", findGearRatios(string(bytes)))
}
