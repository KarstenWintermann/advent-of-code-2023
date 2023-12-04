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

func findFirstDigit(input string) string {
	for i := 0; i < len(input); i++ {
		rest := input[i:]
		switch {
		case strings.HasPrefix(rest, "1"):
			return "1"
		case strings.HasPrefix(rest, "one"):
			return "1"
		case strings.HasPrefix(rest, "2"):
			return "2"
		case strings.HasPrefix(rest, "two"):
			return "2"
		case strings.HasPrefix(rest, "3"):
			return "3"
		case strings.HasPrefix(rest, "three"):
			return "3"
		case strings.HasPrefix(rest, "4"):
			return "4"
		case strings.HasPrefix(rest, "four"):
			return "4"
		case strings.HasPrefix(rest, "5"):
			return "5"
		case strings.HasPrefix(rest, "five"):
			return "5"
		case strings.HasPrefix(rest, "6"):
			return "6"
		case strings.HasPrefix(rest, "six"):
			return "6"
		case strings.HasPrefix(rest, "7"):
			return "7"
		case strings.HasPrefix(rest, "seven"):
			return "7"
		case strings.HasPrefix(rest, "8"):
			return "8"
		case strings.HasPrefix(rest, "eight"):
			return "8"
		case strings.HasPrefix(rest, "9"):
			return "9"
		case strings.HasPrefix(rest, "nine"):
			return "9"
		}
	}

	return ""
}

func findLastDigit(input string) string {
	for i := len(input); i >= 0; i-- {
		rest := input[i:]
		switch {
		case strings.HasPrefix(rest, "1"):
			return "1"
		case strings.HasPrefix(rest, "one"):
			return "1"
		case strings.HasPrefix(rest, "2"):
			return "2"
		case strings.HasPrefix(rest, "two"):
			return "2"
		case strings.HasPrefix(rest, "3"):
			return "3"
		case strings.HasPrefix(rest, "three"):
			return "3"
		case strings.HasPrefix(rest, "4"):
			return "4"
		case strings.HasPrefix(rest, "four"):
			return "4"
		case strings.HasPrefix(rest, "5"):
			return "5"
		case strings.HasPrefix(rest, "five"):
			return "5"
		case strings.HasPrefix(rest, "6"):
			return "6"
		case strings.HasPrefix(rest, "six"):
			return "6"
		case strings.HasPrefix(rest, "7"):
			return "7"
		case strings.HasPrefix(rest, "seven"):
			return "7"
		case strings.HasPrefix(rest, "8"):
			return "8"
		case strings.HasPrefix(rest, "eight"):
			return "8"
		case strings.HasPrefix(rest, "9"):
			return "9"
		case strings.HasPrefix(rest, "nine"):
			return "9"
		}
	}

	return ""
}

func findCalib2(input string) int {
	res, _ := strconv.Atoi(findFirstDigit(input) + findLastDigit(input))
	return res
}

func findCalib(input string) int {
	res, _ := strconv.Atoi(string([]rune(input)[strings.IndexAny(input, "123456789")]) + string([]rune(input)[strings.LastIndexAny(input, "123456789")]))
	return res
}

func findCalibSum(input string) int {
	defer timeTrack(time.Now(), "Part 1")
	calibSum := 0
	values := strings.Split(input, "\r\n")
	for _, v := range values {
		calibSum += findCalib(v)
	}
	return calibSum
}

func findCalib2Sum(input string) int {
	defer timeTrack(time.Now(), "Part 2")
	calibSum := 0
	values := strings.Split(input, "\r\n")
	for _, v := range values {
		calibSum += findCalib2(v)
	}
	return calibSum
}

func main() {
	defer timeTrack(time.Now(), "Overall")
	bytes, err := os.ReadFile("2023.1.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("Calibration sum: %v, with number strings: %v\n", findCalibSum(string(bytes)), findCalib2Sum(string(bytes)))
}
