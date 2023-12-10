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

func findPrediction(input []int) int {
	sameValues := true
	for i := 0; i < len(input); i++ {
		if input[i] != input[0] {
			sameValues = false
			break
		}
	}
	if sameValues {
		return input[0]
	} else {
		nextinput := make([]int, len(input)-1)
		for i := 0; i < len(nextinput); i++ {
			nextinput[i] = input[i+1] - input[i]
		}
		return input[len(input)-1] + findPrediction(nextinput)
	}
}

func findPrediction2(input []int) int {
	sameValues := true
	for i := 0; i < len(input); i++ {
		if input[i] != input[0] {
			sameValues = false
			break
		}
	}
	if sameValues {
		return input[0]
	} else {
		nextinput := make([]int, len(input)-1)
		for i := 0; i < len(nextinput); i++ {
			nextinput[i] = input[i+1] - input[i]
		}
		return input[0] - findPrediction2(nextinput)
	}
}

func findPredictionsSum(input string) int {
	defer timeTrack(time.Now(), "Part 1")
	lines := strings.Split(input, "\r\n")
	result := 0
	for _, line := range lines {
		nums := strings.Fields(line)
		ints := make([]int, len(nums))
		for i, num := range nums {
			ints[i], _ = strconv.Atoi(num)
		}
		result += findPrediction(ints)
	}
	return result
}

func findPredictionsSum2(input string) int {
	defer timeTrack(time.Now(), "Part 2")
	lines := strings.Split(input, "\r\n")
	result := 0
	for _, line := range lines {
		nums := strings.Fields(line)
		ints := make([]int, len(nums))
		for i, num := range nums {
			ints[i], _ = strconv.Atoi(num)
		}
		result += findPrediction2(ints)
	}
	return result
}

func main() {
	defer timeTrack(time.Now(), "2023.9")
	bytes, err := os.ReadFile("2023.9.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("findPredictionsSum: %v\n", findPredictionsSum(string(bytes)))
	fmt.Printf("findPredictionsSum2: %v\n", findPredictionsSum2(string(bytes)))
}
