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

type coord_t struct {
	x, y, z int
}

type brick_t struct {
	from, to coord_t
}

func parseCoord(input string) coord_t {
	var ret coord_t
	coords := strings.Split(input, ",")
	ret.x, _ = strconv.Atoi(coords[0])
	ret.y, _ = strconv.Atoi(coords[1])
	ret.z, _ = strconv.Atoi(coords[2])
	return ret
}

func makeBrick(a, b coord_t) brick_t {
	if a.x < b.x {
		return brick_t{a, b}
	}
	if a.x == b.x && a.y < b.y {
		return brick_t{a, b}
	}
	if a.x == b.x && a.y == b.y && a.z < b.z {
		return brick_t{a, b}
	}
	return brick_t{b, a}
}

func intersectsXY(a, b brick_t) bool {
	if a.to.x >= b.from.x && b.to.x >= a.from.x &&
		a.to.y >= b.from.y && b.to.y >= a.from.y {
		return true
	}
	return false
}

func build_graph(input string, bricks *[]brick_t, is_supported_by *map[int]map[int]bool, supports *map[int]map[int]bool) {
	lines := strings.Split(input, "\r\n")
	(*bricks) = make([]brick_t, len(lines))
	(*is_supported_by) = make(map[int]map[int]bool)
	(*supports) = make(map[int]map[int]bool)
	highest_z := 0
	for i, line := range lines {
		leftright := strings.Split(line, "~")
		(*bricks)[i] = makeBrick(parseCoord(leftright[0]), parseCoord(leftright[1]))
		(*is_supported_by)[i] = make(map[int]bool)
		(*supports)[i] = make(map[int]bool)
		if (*bricks)[i].to.z > highest_z {
			highest_z = (*bricks)[i].to.z
		}
	}

	bottoms := make([][]int, highest_z+1)
	tops := make([][]int, highest_z+1)

	for i := 0; i <= highest_z; i++ {
		bottoms[i] = make([]int, 0)
		tops[i] = make([]int, 0)
	}

	for i, brick := range *bricks {
		bottoms[brick.from.z] = append(bottoms[brick.from.z], i)
		tops[brick.to.z] = append(tops[brick.to.z], i)
	}

	dropped := true
	for dropped {
		dropped = false
		for z := 0; z <= highest_z; z++ {
			to_delete := make([]int, 0)
			for _, i := range bottoms[z] {
				drops := 0
				for support_z := z - 1; support_z > 0; support_z-- {
					found_support := false
					for _, j := range tops[support_z] {
						if i != j && intersectsXY((*bricks)[i], (*bricks)[j]) {
							found_support = true
							(*supports)[j][i] = true
							(*is_supported_by)[i][j] = true
						}
					}
					if found_support {
						break
					}
					if !found_support {
						drops += 1
						dropped = true
					}
				}
				if drops > 0 {
					tops[(*bricks)[i].to.z] = slices.DeleteFunc(tops[(*bricks)[i].to.z], func(b int) bool { return b == i })
					to_delete = append(to_delete, i)
					(*bricks)[i].from.z -= drops
					(*bricks)[i].to.z -= drops
					tops[(*bricks)[i].to.z] = append(tops[(*bricks)[i].to.z], i)
					bottoms[(*bricks)[i].from.z] = append(bottoms[(*bricks)[i].from.z], i)
				}
			}
			if len(to_delete) > 0 {
				bottoms[z] = slices.DeleteFunc(bottoms[z], func(b int) bool {
					return slices.Contains(to_delete, b)
				})
			}
		}
	}
}

func part1(input string) int {
	var bricks []brick_t
	var is_supported_by map[int]map[int]bool
	var supports map[int]map[int]bool

	build_graph(input, &bricks, &is_supported_by, &supports)

	removable := 0
	for i := range bricks {
		can_remove := true
		if len(supports[i]) > 0 {
			for j := range supports[i] {
				if len(is_supported_by[j]) == 1 {
					can_remove = false
				}
			}
		}
		if can_remove {
			removable += 1
		}
	}

	return removable
}

func findFallingBricks(brick int, fallen_bricks *map[int]bool, is_supported_by *map[int]map[int]bool, supports *map[int]map[int]bool) int {
	falling := 0
	if len((*supports)[brick]) > 0 {
		new_fallen := make(map[int]bool)
		for supported := range (*supports)[brick] {
			has_other_support := false
			for other_support := range (*is_supported_by)[supported] {
				_, ok := (*fallen_bricks)[other_support]
				if !ok {
					has_other_support = true
				}
			}
			if !has_other_support {
				new_fallen[supported] = true
				(*fallen_bricks)[supported] = true
			}
		}

		for fallen := range new_fallen {
			falling += 1 + findFallingBricks(fallen, fallen_bricks, is_supported_by, supports)
		}
	}

	return falling
}

func part2(input string) int {
	var bricks []brick_t
	var is_supported_by map[int]map[int]bool
	var supports map[int]map[int]bool

	build_graph(input, &bricks, &is_supported_by, &supports)

	falling := 0
	for i := range bricks {
		fallen_bricks := make(map[int]bool)
		fallen_bricks[i] = true
		findFallingBricks(i, &fallen_bricks, &is_supported_by, &supports)
		falling += len(fallen_bricks) - 1
	}

	return falling
}

func main() {
	defer timeTrack(time.Now(), "2023.22")
	bytes, err := os.ReadFile("2023.22.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("part1: %v\n", part1(string(bytes)))
	fmt.Printf("part2: %v\n", part2(string(bytes)))
}
