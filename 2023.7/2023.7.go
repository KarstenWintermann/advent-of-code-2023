package main

import (
	"cmp"
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

const FIVE_OF_A_KIND = 7
const FOUR_OF_A_KIND = 6
const FULL_HOUSE = 5
const THREE_OF_A_KIND = 4
const TWO_PAIRS = 3
const ONE_PAIR = 2

func score(cards string) int {
	res := make(map[string]int, len(cards))
	for i := 0; i < len(cards); i++ {
		res[cards[i:i+1]] += 1
	}

	score := 0

	for nums, _ := range res {
		switch res[nums] {
		case 5:
			score = FIVE_OF_A_KIND
		case 4:
			score = FOUR_OF_A_KIND
		case 3:
			if score == ONE_PAIR {
				score = FULL_HOUSE
			} else {
				score = THREE_OF_A_KIND
			}
		case 2:
			if score == THREE_OF_A_KIND {
				score = FULL_HOUSE
			} else if score == ONE_PAIR {
				score = TWO_PAIRS
			} else {
				score = ONE_PAIR
			}
		}
	}

	return score
}

func score2(cards string) int {
	res := make(map[string]int, len(cards))
	for i := 0; i < len(cards); i++ {
		res[cards[i:i+1]] += 1
	}

	score := 0
	num_jokers := res["0"]

	for nums, _ := range res {
		if nums != "0" {
			switch res[nums] {
			case 5:
				score = FIVE_OF_A_KIND
			case 4:
				if num_jokers > 0 {
					score = FIVE_OF_A_KIND
				} else {
					score = FOUR_OF_A_KIND
				}
			case 3:
				if num_jokers == 2 {
					score = FIVE_OF_A_KIND
				} else if num_jokers == 1 {
					score = FOUR_OF_A_KIND
				} else if score == 2 {
					score = FULL_HOUSE
				} else {
					score = THREE_OF_A_KIND
				}
			case 2:
				if score == THREE_OF_A_KIND {
					score = FULL_HOUSE
				} else if score == ONE_PAIR {
					if num_jokers == 1 {
						score = FULL_HOUSE
					} else {
						score = TWO_PAIRS
					}
				} else {
					score = ONE_PAIR
				}
			}
		}
	}

	if score == 2 {
		switch num_jokers {
		case 3:
			score = FIVE_OF_A_KIND
		case 2:
			score = FOUR_OF_A_KIND
		case 1:
			score = THREE_OF_A_KIND
		}
	} else if score == 0 {
		switch num_jokers {
		case 5:
			score = FIVE_OF_A_KIND
		case 4:
			score = FIVE_OF_A_KIND
		case 3:
			score = FOUR_OF_A_KIND
		case 2:
			score = THREE_OF_A_KIND
		case 1:
			score = ONE_PAIR
		}
	}

	return score
}

type handAndBet struct {
	hand  string
	bet   int
	score int
}

func findWinnings(input string) int {
	defer timeTrack(time.Now(), "Part 1")

	lines := strings.Split(input, "\r\n")
	cards_array := make([]handAndBet, 0)
	totalScore := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		hand := fields[0]
		hand = strings.ReplaceAll(hand, "A", "E")
		hand = strings.ReplaceAll(hand, "K", "D")
		hand = strings.ReplaceAll(hand, "Q", "C")
		hand = strings.ReplaceAll(hand, "J", "B")
		hand = strings.ReplaceAll(hand, "T", "A")
		bet, _ := strconv.Atoi(fields[1])
		cards_array = append(cards_array, handAndBet{hand, bet, score(hand)})
	}

	handsCmp := func(a, b handAndBet) int {
		if a.score == b.score {
			return strings.Compare(a.hand, b.hand)
		} else {
			return cmp.Compare(a.score, b.score)
		}
	}

	slices.SortFunc(cards_array, handsCmp)

	for i := 0; i < len(cards_array); i++ {
		totalScore += cards_array[i].bet * (i + 1)
	}

	return totalScore
}

func findWinnings2(input string) int {
	defer timeTrack(time.Now(), "Part 2")
	lines := strings.Split(input, "\r\n")
	cards_array := make([]handAndBet, len(lines))
	totalScore := 0
	for i, line := range lines {
		fields := strings.Fields(line)
		hand := fields[0]
		hand = strings.ReplaceAll(hand, "A", "E")
		hand = strings.ReplaceAll(hand, "K", "D")
		hand = strings.ReplaceAll(hand, "Q", "C")
		hand = strings.ReplaceAll(hand, "T", "A")
		hand = strings.ReplaceAll(hand, "J", "0")
		bet, _ := strconv.Atoi(fields[1])
		cards_array[i] = handAndBet{hand, bet, score2(hand)}
	}

	handsCmp := func(a, b handAndBet) int {
		if a.score == b.score {
			return strings.Compare(a.hand, b.hand)
		} else {
			return cmp.Compare(a.score, b.score)
		}
	}

	slices.SortFunc(cards_array, handsCmp)

	for i := 0; i < len(cards_array); i++ {
		totalScore += cards_array[i].bet * (i + 1)
	}

	return totalScore
}

func main() {
	defer timeTrack(time.Now(), "2023.7")
	bytes, err := os.ReadFile("2023.7.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("test: %v\n", findWinnings(string(bytes)))
	fmt.Printf("test: %v\n", findWinnings2(string(bytes)))
}
