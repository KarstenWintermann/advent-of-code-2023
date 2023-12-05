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

type intRange struct {
	destination int
	source      int
	count       int
}

func findInRange(sectionRange *[]intRange, from int) int {
	for _, elem := range *sectionRange {
		if from >= elem.source && from <= elem.source+elem.count {
			return elem.destination + (from - elem.source)
		}
	}
	return from
}

func findLowestLocation(input string) int {
	sections := strings.Split(input, "\r\n\r\n")

	seedToSoil := make([]intRange, 0, 10)
	soilToFertilizer := make([]intRange, 0, 10)
	fertilizerToWater := make([]intRange, 0, 10)
	waterToLight := make([]intRange, 0, 10)
	lightToTemperature := make([]intRange, 0, 10)
	temperatureToHumidity := make([]intRange, 0, 10)
	humidityToLocation := make([]intRange, 0, 10)

	var seeds []string

	for _, section := range sections {
		if strings.HasPrefix(section, "seeds:") {
			seeds = strings.Fields(section[len("seeds:"):])
		} else {
			lines := strings.Split(section, "\r\n")
			var sectionRange *[]intRange
			for _, line := range lines {
				fields := strings.Fields(line)
				if fields[1] == "map:" {
					switch fields[0] {
					case "seed-to-soil":
						sectionRange = &seedToSoil
					case "soil-to-fertilizer":
						sectionRange = &soilToFertilizer
					case "fertilizer-to-water":
						sectionRange = &fertilizerToWater
					case "water-to-light":
						sectionRange = &waterToLight
					case "light-to-temperature":
						sectionRange = &lightToTemperature
					case "temperature-to-humidity":
						sectionRange = &temperatureToHumidity
					case "humidity-to-location":
						sectionRange = &humidityToLocation
					}
				} else if len(fields) == 3 {
					destination, _ := strconv.Atoi(fields[0])
					source, _ := strconv.Atoi(fields[1])
					numberRange, _ := strconv.Atoi(fields[2])

					*sectionRange = append(*sectionRange, intRange{destination, source, numberRange})
				}
			}
		}
	}

	minLocation := 0
	for _, seed := range seeds {
		seedNumber, _ := strconv.Atoi(seed)
		soil := findInRange(&seedToSoil, seedNumber)
		fertilizer := findInRange(&soilToFertilizer, soil)
		water := findInRange(&fertilizerToWater, fertilizer)
		light := findInRange(&waterToLight, water)
		temperature := findInRange(&lightToTemperature, light)
		humidity := findInRange(&temperatureToHumidity, temperature)
		location := findInRange(&humidityToLocation, humidity)

		if minLocation == 0 || location < minLocation {
			minLocation = location
		}
	}

	return minLocation
}

func findLowestLocation2(input string) int {
	sections := strings.Split(input, "\r\n\r\n")

	seedToSoil := make([]intRange, 0, 10)
	soilToFertilizer := make([]intRange, 0, 10)
	fertilizerToWater := make([]intRange, 0, 10)
	waterToLight := make([]intRange, 0, 10)
	lightToTemperature := make([]intRange, 0, 10)
	temperatureToHumidity := make([]intRange, 0, 10)
	humidityToLocation := make([]intRange, 0, 10)

	var seeds []string

	for _, section := range sections {
		if strings.HasPrefix(section, "seeds:") {
			seeds = strings.Fields(section[len("seeds:"):])
		} else {
			lines := strings.Split(section, "\r\n")
			var sectionRange *[]intRange
			for _, line := range lines {
				fields := strings.Fields(line)
				if fields[1] == "map:" {
					switch fields[0] {
					case "seed-to-soil":
						sectionRange = &seedToSoil
					case "soil-to-fertilizer":
						sectionRange = &soilToFertilizer
					case "fertilizer-to-water":
						sectionRange = &fertilizerToWater
					case "water-to-light":
						sectionRange = &waterToLight
					case "light-to-temperature":
						sectionRange = &lightToTemperature
					case "temperature-to-humidity":
						sectionRange = &temperatureToHumidity
					case "humidity-to-location":
						sectionRange = &humidityToLocation
					}
				} else if len(fields) == 3 {
					destination, _ := strconv.Atoi(fields[0])
					source, _ := strconv.Atoi(fields[1])
					numberRange, _ := strconv.Atoi(fields[2])

					*sectionRange = append(*sectionRange, intRange{destination, source, numberRange})
				}
			}
		}
	}

	minLocation := 0
	inRange := false
	seedStart := 0
	seedEnd := 0
	for _, seedEntry := range seeds {
		seedNumber, _ := strconv.Atoi(seedEntry)
		if !inRange {
			inRange = true
			seedStart = seedNumber
		} else {
			inRange = false
			seedEnd = seedStart + seedNumber

			fmt.Printf("Range from %v to %v\n", seedStart, seedEnd)

			for seed := seedStart; seed < seedEnd; seed++ {

				soil := findInRange(&seedToSoil, seed)
				fertilizer := findInRange(&soilToFertilizer, soil)
				water := findInRange(&fertilizerToWater, fertilizer)
				light := findInRange(&waterToLight, water)
				temperature := findInRange(&lightToTemperature, light)
				humidity := findInRange(&temperatureToHumidity, temperature)
				location := findInRange(&humidityToLocation, humidity)

				if minLocation == 0 || location < minLocation {
					minLocation = location
				}
			}
		}
	}

	return minLocation
}

func main() {
	defer timeTrack(time.Now(), "2023.5")
	bytes, err := os.ReadFile("2023.5.input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	fmt.Printf("test: %v\n", findLowestLocation(string(bytes)))
	fmt.Printf("test: %v\n", findLowestLocation2(string(bytes)))
}
