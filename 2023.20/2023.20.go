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

const NORMAL = 0
const FLIPFLOP = 1
const CONJUNCTION = 2

const LOW = 1
const HIGH = 2

type module_t struct {
	modtype    int
	successors []string
}

type pulse_t struct {
	target string
	pulse  int
}

func sendPulse(sender string, targets *[]string, config *map[string]module_t, conjunction *map[string]map[string]int, received *[]pulse_t, pulse int) int {
	ret := 0
	for _, s := range *targets {
		ret += 1
		(*received) = append((*received), pulse_t{s, pulse})
		if (*config)[s].modtype == CONJUNCTION {
			(*conjunction)[s][sender] = pulse
		}
	}
	return ret
}

func part1(input string) int {
	lines := strings.Split(input, "\r\n")
	config := make(map[string]module_t)
	received := make([]pulse_t, 0)
	state := make(map[string]bool)
	conjunction := make(map[string]map[string]int, 0)
	for _, line := range lines {
		rightleft := strings.Split(line, "->")
		modtype := NORMAL
		name := strings.Trim(rightleft[0], " ")
		if strings.HasPrefix(name, "%") {
			name = name[1:]
			modtype = FLIPFLOP
		} else if strings.HasPrefix(name, "&") {
			name = name[1:]
			modtype = CONJUNCTION
			conjunction[name] = make(map[string]int)
		}

		targets := strings.ReplaceAll(rightleft[1], " ", "")
		config[name] = module_t{modtype, strings.Split(targets, ",")}
	}

	for name, module := range config {
		for _, successor := range module.successors {
			if config[successor].modtype == CONJUNCTION {
				conjunction[successor][name] = LOW
			}
		}
	}

	sent_low := 0
	sent_high := 0
	button_pressed := 0
	for true {
		if len(received) == 0 {
			received = append(received, pulse_t{"broadcaster", LOW})
			button_pressed += 1
			if button_pressed > 10000 {
				break
			}
			sent_low += 1
		}

		new_received := make([]pulse_t, 0)
		for _, pulse := range received {
			module := config[pulse.target]
			switch module.modtype {
			case NORMAL:
				sent_low += sendPulse(pulse.target, &module.successors, &config, &conjunction, &new_received, pulse.pulse)

			case FLIPFLOP:
				if pulse.pulse == LOW {
					if state[pulse.target] {
						state[pulse.target] = false
						sent_low += sendPulse(pulse.target, &module.successors, &config, &conjunction, &new_received, LOW)
					} else {
						state[pulse.target] = true
						sent_high += sendPulse(pulse.target, &module.successors, &config, &conjunction, &new_received, HIGH)
					}
				}

			case CONJUNCTION:
				all_high := true
				for _, c := range conjunction[pulse.target] {
					if c == LOW {
						all_high = false
					}
				}
				if all_high {
					sent_low += sendPulse(pulse.target, &module.successors, &config, &conjunction, &new_received, LOW)
				} else {
					if pulse.target == "lb" ||
						pulse.target == "rp" ||
						pulse.target == "cl" ||
						pulse.target == "nj" {
						fmt.Printf("%v fired at step %v\n", pulse.target, button_pressed)
					}
					sent_high += sendPulse(pulse.target, &module.successors, &config, &conjunction, &new_received, HIGH)
				}
			}
		}
		received = new_received
	}

	return sent_high * sent_low
}

func part2(input string) int {
	return 0
}

func main() {
	defer timeTrack(time.Now(), "2023.20")
	bytes, err := os.ReadFile("2023.20.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("part1: %v\n", part1(string(bytes)))
	fmt.Printf("part2: %v\n", part2(string(bytes)))
}
