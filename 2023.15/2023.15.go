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

type box_t struct {
	label  string
	number int
}

func hashMap(input string) int {
	boxes := make([][]box_t, 256)
	steps := strings.Split(input, ",")
	for _, step := range steps {
		pos_op := strings.IndexAny(step, "-=")
		label := step[0:pos_op]
		hashvalue := hash(label)
		op := step[pos_op : pos_op+1]
		box := boxes[hashvalue]
		if op == "=" {
			opnum, _ := strconv.Atoi(step[pos_op+1:])
			found_box := false
			for i := 0; i < len(box); i++ {
				if box[i].label == label {
					box[i] = box_t{label, opnum}
					found_box = true
				}
			}
			if !found_box {
				box = append(box, box_t{label, opnum})
				boxes[hashvalue] = box
			}
		} else if op == "-" {
			for i := 0; i < len(box); i++ {
				if box[i].label == label {
					boxes[hashvalue] = slices.Delete(box, i, i+1)
				}
			}
		}
	}

	sum := 0
	for boxnum, box := range boxes {
		for i, elem := range box {
			sum += (boxnum + 1) * (i + 1) * elem.number
		}
	}
	return sum
}

func hash(input string) int {
	current := 0
	for i := 0; i < len(input); i++ {
		current = ((int(input[i]) + current) * 17) % 256
	}
	return current
}

func hashSteps(input string) int {
	steps := strings.Split(input, ",")
	sum := 0
	for _, step := range steps {
		sum += hash(step)
	}
	return sum
}

func main() {
	defer timeTrack(time.Now(), "2023.15")
	bytes, err := os.ReadFile("2023.15.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("test: %v\n", hash("HASH"))
	fmt.Printf("test: %v\n", hashSteps(string(bytes)))
	fmt.Printf("test: %v\n", hashMap(string(bytes)))
}
