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

type state_t struct {
	x, m, a, s int
}

type range_t struct {
	from, to int
}

type ranges_t struct {
	x, m, a, s range_t
}

func evaluate_all(state_in ranges_t, workflow *map[string][]string, step string) []ranges_t {
	ret := make([]ranges_t, 0)
	for _, rule := range (*workflow)[step] {
		if strings.ContainsAny(rule, "<>") {
			lessthan := false
			if strings.Contains(rule, "<") {
				lessthan = true
			}
			rule = strings.ReplaceAll(rule, "<", " ")
			rule = strings.ReplaceAll(rule, ">", " ")
			rule = strings.ReplaceAll(rule, ":", " ")
			fields := strings.Fields(rule)
			variable := fields[0]
			operand, _ := strconv.Atoi(fields[1])
			target := fields[2]
			state_true := state_in
			state_false := state_in
			switch variable {
			case "x":
				if lessthan {
					if state_true.x.to > operand {
						state_true.x.to = operand - 1
					}
					if state_false.x.from < operand {
						state_false.x.from = operand
					}
				} else {
					if state_true.x.from < operand {
						state_true.x.from = operand + 1
					}
					if state_false.x.to > operand {
						state_false.x.to = operand
					}
				}
			case "m":
				if lessthan {
					if state_true.m.to > operand {
						state_true.m.to = operand - 1
					}
					if state_false.m.from < operand {
						state_false.m.from = operand
					}
				} else {
					if state_true.m.from < operand {
						state_true.m.from = operand + 1
					}
					if state_false.m.to > operand {
						state_false.m.to = operand
					}
				}
			case "a":
				if lessthan {
					if state_true.a.to > operand {
						state_true.a.to = operand - 1
					}
					if state_false.a.from < operand {
						state_false.a.from = operand
					}
				} else {
					if state_true.a.from < operand {
						state_true.a.from = operand + 1
					}
					if state_false.a.to > operand {
						state_false.a.to = operand
					}
				}
			case "s":
				if lessthan {
					if state_true.s.to > operand {
						state_true.s.to = operand - 1
					}
					if state_false.s.from < operand {
						state_false.s.from = operand
					}
				} else {
					if state_true.s.from < operand {
						state_true.s.from = operand + 1
					}
					if state_false.s.to > operand {
						state_false.s.to = operand
					}
				}
			}
			if state_true.x.from <= state_true.x.to &&
				state_true.m.from <= state_true.m.to &&
				state_true.a.from <= state_true.a.to &&
				state_true.s.from <= state_true.s.to {
				if target == "A" {
					ret = append(ret, state_true)
				} else {
					for _, s := range evaluate_all(state_true, workflow, target) {
						ret = append(ret, s)
					}
				}
			}

			state_in = state_false
		} else if rule == "A" {
			ret = append(ret, state_in)
		} else if rule != "R" {
			for _, s := range evaluate_all(state_in, workflow, rule) {
				ret = append(ret, s)
			}
		}
	}

	return ret
}

func evaluate(state state_t, rules []string) string {
	for _, rule := range rules {
		if strings.Contains(rule, "<") {
			rule = strings.ReplaceAll(rule, "<", " ")
			rule = strings.ReplaceAll(rule, ":", " ")
			fields := strings.Fields(rule)
			variable := fields[0]
			operand, _ := strconv.Atoi(fields[1])
			target := fields[2]

			switch variable {
			case "x":
				if state.x < operand {
					return target
				}
			case "m":
				if state.m < operand {
					return target
				}
			case "a":
				if state.a < operand {
					return target
				}
			case "s":
				if state.s < operand {
					return target
				}
			}
		} else if strings.Contains(rule, ">") {
			rule = strings.ReplaceAll(rule, ">", " ")
			rule = strings.ReplaceAll(rule, ":", " ")
			fields := strings.Fields(rule)
			variable := fields[0]
			operand, _ := strconv.Atoi(fields[1])
			target := fields[2]

			switch variable {
			case "x":
				if state.x > operand {
					return target
				}
			case "m":
				if state.m > operand {
					return target
				}
			case "a":
				if state.a > operand {
					return target
				}
			case "s":
				if state.s > operand {
					return target
				}
			}
		} else {
			return rule
		}
	}
	return ""
}

func part2(input string) int {
	sections := strings.Split(input, "\r\n\r\n")
	workflow := make(map[string][]string)

	rules := strings.Split(sections[0], "\r\n")
	for _, rule := range rules {
		name := rule[:strings.Index(rule, "{")]
		steps := strings.Split(rule[strings.Index(rule, "{")+1:strings.Index(rule, "}")], ",")
		for _, step := range steps {
			workflow[name] = append(workflow[name], step)
		}
	}

	sum := 0
	for _, r := range evaluate_all(ranges_t{range_t{1, 4000}, range_t{1, 4000}, range_t{1, 4000}, range_t{1, 4000}}, &workflow, "in") {
		sum += (r.x.to - r.x.from + 1) * (r.m.to - r.m.from + 1) * (r.a.to - r.a.from + 1) * (r.s.to - r.s.from + 1)
	}

	return sum
}

func part1(input string) int {
	sections := strings.Split(input, "\r\n\r\n")
	workflow := make(map[string][]string)

	rules := strings.Split(sections[0], "\r\n")
	for _, rule := range rules {
		name := rule[:strings.Index(rule, "{")]
		steps := strings.Split(rule[strings.Index(rule, "{")+1:strings.Index(rule, "}")], ",")
		for _, step := range steps {
			workflow[name] = append(workflow[name], step)
		}
	}

	sum := 0
	inputs := strings.Split(sections[1], "\r\n")
	for _, input := range inputs {
		input = strings.ReplaceAll(input, "{", "")
		input = strings.ReplaceAll(input, "}", "")
		vars := strings.Split(input, ",")
		state := state_t{0, 0, 0, 0}
		for _, v := range vars {
			switch v[0] {
			case 'x':
				state.x, _ = strconv.Atoi(v[2:])
			case 'm':
				state.m, _ = strconv.Atoi(v[2:])
			case 'a':
				state.a, _ = strconv.Atoi(v[2:])
			case 's':
				state.s, _ = strconv.Atoi(v[2:])
			}
		}

		step := evaluate(state, workflow["in"])
		for step != "A" && step != "R" {
			step = evaluate(state, workflow[step])
		}

		if step == "A" {
			sum += state.x + state.m + state.a + state.s
		}
	}

	return sum
}

func main() {
	defer timeTrack(time.Now(), "2023.19")
	bytes, err := os.ReadFile("2023.19.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("part1: %v\n", part1(string(bytes)))
	fmt.Printf("part2: %v\n", part2(string(bytes)))
}
