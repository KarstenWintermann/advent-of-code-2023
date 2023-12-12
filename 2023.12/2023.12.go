package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func groupsMatch(pattern string, blocks []int) bool {
	fields := strings.Fields(strings.ReplaceAll(pattern, ".", " "))
	if len(fields) != len(blocks) {
		return false
	}
	for i, block := range blocks {
		if block != len(fields[i]) {
			return false
		}
	}
	return true
}

func patternMatch(pattern string, test string, frompos int) int {
	cachekey := fmt.Sprintf("%v %v %v", pattern, test, frompos)
	ret, ok := cache[cachekey]
	if ok {
		return ret
	}
	for i := frompos; i < len(test); i++ {
		if pattern[i:i+1] != "?" {
			if pattern[i:i+1] != test[i:i+1] {
				cache[cachekey] = 0
				return 0
			}
		}
	}
	cache[cachekey] = 1
	return 1
}

var cache map[string]int

func sumArrangements(pattern string, prefix string, numbers []int, freespaces int) int {
	cachekey := fmt.Sprintf("%v %v %v %v", pattern, prefix, numbers, freespaces)
	ret, ok := cache[cachekey]
	if ok {
		return ret
	}
	ret = 0
	if len(numbers) == 0 {
		if patternMatch(pattern, prefix+strings.Repeat(".", freespaces), len(prefix)) == 1 {
			cache[cachekey] = 1
			return 1
		} else {
			cache[cachekey] = 0
			return 0
		}
	} else {
		for i := 0; i <= freespaces; i++ {
			nextprefix := prefix + strings.Repeat(".", i) + strings.Repeat("#", numbers[0])
			if len(numbers) > 1 {
				nextprefix = nextprefix + "."
			}
			if patternMatch(pattern, nextprefix, len(prefix)) == 1 {
				ret += sumArrangements(pattern, nextprefix, numbers[1:], freespaces-i)
			}
		}
		cache[cachekey] = ret
		return ret
	}
}

func sumArrangements3(spring_blocks []string, numbers []int) int {
	cachekey := fmt.Sprintf("%v %v", spring_blocks, numbers)
	sum, ok := cache[cachekey]
	if ok {
		return sum
	}
	sum = 0
	if len(numbers) == 0 {
		for _, block := range spring_blocks {
			if strings.Contains(block, "#") { // at least one # left, but no more numbers -> return 0
				cache[cachekey] = 0
				return 0
			}
		}

		cache[cachekey] = 1
		return 1
	}
	for i, block := range spring_blocks {
		for j := 0; j < len(block); j++ {
			if len(block[j:]) == numbers[0] { // fits exactly
				if strings.Contains(block[j:], "#") { // must put here
					sum = sum + sumArrangements3(spring_blocks[i+1:], numbers[1:])
					cache[cachekey] = sum
					return sum
				} else { // carry on with next block
					j += numbers[0]
					sum += sumArrangements3(spring_blocks[i+1:], numbers[1:])
				}
			} else if len(block[j:]) >= numbers[0] && (j == 0 || block[j-1:j] == "?") && block[j+numbers[0]:j+numbers[0]+1] == "?" {
				spring_blocks_copy := make([]string, len(spring_blocks)-i)
				copy(spring_blocks_copy, spring_blocks[i:])
				spring_blocks_copy = slices.Replace(spring_blocks_copy, 0, 1, block[j+numbers[0]+1:])
				if block[j:j+1] == "#" { // must put here
					sum = sum + sumArrangements3(spring_blocks_copy, numbers[1:])
					cache[cachekey] = sum
					return sum
				} else { // carry on
					sum += sumArrangements3(spring_blocks_copy, numbers[1:])
				}
			} else if block[j:j+1] == "#" { // can't put anything here but have to
				cache[cachekey] = sum
				return sum
			}
		}
	}

	cache[cachekey] = sum
	return sum
}

func findArrangements3(input string) int {
	defer timeTrack(time.Now(), "Part 2")
	lines := strings.Split(input, "\r\n")
	arrangements := 0
	cache = make(map[string]int, 0)
	for _, line := range lines {
		blocks := strings.Fields(line)
		pattern := blocks[0]
		groups := blocks[1]

		pattern = strings.Repeat(pattern+"?", 5)
		pattern = pattern[0 : len(pattern)-1]
		groups = strings.Repeat(groups+",", 5)
		groups = groups[0 : len(groups)-1]

		numberStrings := strings.Split(groups, ",")
		numbers := make([]int, len(numberStrings))

		patternlength := 0
		for i, number := range numberStrings {
			numbers[i], _ = strconv.Atoi(number)
			patternlength += numbers[i] + 1
		}
		patternlength--

		spring_blocks := strings.Fields(strings.ReplaceAll(pattern, ".", " "))

		sum := sumArrangements3(spring_blocks, numbers)

		//fmt.Printf("line nr %v, pattern %v, arrangements %v\n", line_no, pattern, sum)

		arrangements += sum
	}
	return arrangements
}

func findArrangements2(input string) int {
	lines := strings.Split(input, "\r\n")
	arrangements := 0
	cache = make(map[string]int)
	for _, line := range lines {
		blocks := strings.Fields(line)
		pattern := blocks[0]
		groups := blocks[1]

		pattern = strings.Repeat(pattern+"?", 5)
		pattern = pattern[0 : len(pattern)-1]
		groups = strings.Repeat(groups+",", 5)
		groups = groups[0 : len(groups)-1]

		numberStrings := strings.Split(groups, ",")
		numbers := make([]int, len(numberStrings))
		patternlength := 0
		for i, number := range numberStrings {
			numbers[i], _ = strconv.Atoi(number)
			patternlength += numbers[i] + 1
		}
		patternlength--

		freespaces := len(pattern) - patternlength
		arrangements += sumArrangements(pattern, "", numbers, freespaces)

		//fmt.Printf("line nr %v, pattern %v, arrangements %v\n", line_no, pattern, arrangements)
	}
	return arrangements
}

func findArrangements(input string) int {
	defer timeTrack(time.Now(), "Part 1")
	lines := strings.Split(input, "\r\n")
	sum := 0
	for _, line := range lines {
		blocks := strings.Fields(line)
		pattern := blocks[0]
		groups := blocks[1]
		numberStrings := strings.Split(groups, ",")
		numbers := make([]int, len(numberStrings))
		arrangements := 0
		for i, number := range numberStrings {
			numbers[i], _ = strconv.Atoi(number)
		}

		n_unknown := strings.Count(pattern, "?")
		for j := 0; j < (1 << n_unknown); j++ {
			teststr := pattern
			for k := 0; k < n_unknown; k++ {
				if j&(1<<k) == 0 {
					teststr = strings.Replace(teststr, "?", ".", 1)
				} else {
					teststr = strings.Replace(teststr, "?", "#", 1)
				}
			}
			if groupsMatch(teststr, numbers) {
				arrangements += 1
			}
		}

		//fmt.Printf("line nr %v, pattern %v, arrangements %v\n", line_no, pattern, arrangements)
		sum += arrangements
	}
	return sum
}

func main() {
	defer timeTrack(time.Now(), "2023.12")
	bytes, err := os.ReadFile("2023.12.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("test: %v\n", findArrangements(string(bytes)))
	//fmt.Printf("test: %v\n", findArrangements2(string(bytes)))
	fmt.Printf("test: %v\n", findArrangements3(string(bytes)))
}
