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

func transpose(a []string) []string {
	newArr := make([]string, len(a[0]))
	for i := 0; i < len(a[0]); i++ {
		for j := 0; j < len(a); j++ {
			newArr[i] += a[j][i : i+1]
		}
	}

	return newArr
}

func tiltLeft(input []string) []string {
	output := make([]string, len(input))
	for line_no, line := range input {
		for true {
			moved := false
			for i := 0; i < len(line)-1; i++ {
				if line[i:i+1] == "." && line[i+1:i+2] == "O" {
					line = line[0:i] + "O." + line[i+2:]
					moved = true
				}
			}
			output[line_no] = line
			if !moved {
				break
			}
		}
	}
	return output
}

func tiltRight(input []string) []string {
	output := make([]string, len(input))
	for line_no, line := range input {
		for true {
			moved := false
			for i := len(line) - 2; i >= 0; i-- {
				if line[i:i+1] == "O" && line[i+1:i+2] == "." {
					line = line[0:i] + ".O" + line[i+2:]
					moved = true
				}
			}
			output[line_no] = line
			if !moved {
				break
			}
		}
	}
	return output
}

func tiltUp(input []string) []string {
	return transpose(tiltLeft(transpose(input)))
}

func tiltDown(input []string) []string {
	return transpose(tiltRight(transpose(input)))
}

var cache map[string][]string

func cycle(input []string) []string {
	return tiltRight(tiltDown(tiltLeft(tiltUp(input))))
}

func weight(input []string) int {
	total_load := 0
	for line_no, line := range input {
		total_load += strings.Count(line, "O") * (len(input) - line_no)
	}
	return total_load
}

func equals(a, b []string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func findTotalLoad2(input string) int {
	lines := strings.Split(input, "\r\n")
	loads := make([]int, 0)
	cycled := make([][]string, 0)

	for i := 0; ; i++ {
		lines = cycle(lines)

		for j, previous := range cycled {
			if equals(previous, lines) {
				offset := j
				cycle_length := i - j
				return loads[offset+((1000000000-offset)%cycle_length)-1]
			}
		}

		loads = append(loads, weight(lines))
		cycled = append(cycled, lines)
	}
}

func findTotalLoad(input string) int {
	lines := strings.Split(input, "\r\n")

	tilted := transpose(lines)
	for line_no, line := range tilted {
		for true {
			moved := false
			for i := 0; i < len(line)-1; i++ {
				if line[i:i+1] == "." && line[i+1:i+2] == "O" {
					line = line[0:i] + "O." + line[i+2:]
					moved = true
				}
			}
			if !moved {
				break
			} else {
				tilted[line_no] = line
			}
		}
	}

	lines = transpose(tilted)
	total_load := 0
	for line_no, line := range lines {
		total_load += strings.Count(line, "O") * (len(lines) - line_no)
	}
	return total_load
}

func main() {
	defer timeTrack(time.Now(), "2023.14")
	bytes, err := os.ReadFile("2023.14.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("findTotalLoad: %v\n", findTotalLoad(string(bytes)))
	fmt.Printf("findTotalLoad2: %v\n", findTotalLoad2(string(bytes)))
}
