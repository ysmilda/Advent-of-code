package aoc2023day5

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/ysmilda/Advent-of-code/pkg/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input garden
}

type garden struct {
	seeds                 []int
	seedToSoil            []mapping
	soilToFertilizer      []mapping
	fertilizerToWater     []mapping
	waterToLight          []mapping
	lightToTemperature    []mapping
	temperatureToHumidity []mapping
	humidityToLocation    []mapping
}

type mapping struct {
	source      int
	destination int
	spread      int
}

func MustGetSolver() solver.Solver {
	return puzzle{
		input: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 5
}

func (s puzzle) Part1() (int, error) {
	lowest := 1 << 32

	for _, seed := range s.input.seeds {
		soil := getDestination(seed, s.input.seedToSoil)
		fertilizer := getDestination(soil, s.input.soilToFertilizer)
		water := getDestination(fertilizer, s.input.fertilizerToWater)
		light := getDestination(water, s.input.waterToLight)
		temperature := getDestination(light, s.input.lightToTemperature)
		humidity := getDestination(temperature, s.input.temperatureToHumidity)
		location := getDestination(humidity, s.input.humidityToLocation)

		if location < lowest {
			lowest = location
		}
	}

	return lowest, nil
}

func (s puzzle) Part2() (int, error) {
	lowest := 1 << 32

	for i := 0; i < len(s.input.seeds); i += 2 {
		for j := 0; j < s.input.seeds[i+1]; j++ {
			soil := getDestination(s.input.seeds[i]+j, s.input.seedToSoil)
			fertilizer := getDestination(soil, s.input.soilToFertilizer)
			water := getDestination(fertilizer, s.input.fertilizerToWater)
			light := getDestination(water, s.input.waterToLight)
			temperature := getDestination(light, s.input.lightToTemperature)
			humidity := getDestination(temperature, s.input.temperatureToHumidity)
			location := getDestination(humidity, s.input.humidityToLocation)

			if location < lowest {
				lowest = location
			}

		}
	}

	return lowest, nil
}

func parse(input string) garden {
	lines := strings.Split(input, "\n")

	garden := garden{}
	var target *[]mapping

	for i, line := range lines {
		if i == 0 {
			garden.seeds = toInt(strings.Fields(strings.Split(line, ":")[1]))
			continue
		}

		if line == "" {
			target = nil
			continue
		}

		switch strings.Fields(line)[0] {
		case "seed-to-soil":
			target = &garden.seedToSoil
		case "soil-to-fertilizer":
			target = &garden.soilToFertilizer
		case "fertilizer-to-water":
			target = &garden.fertilizerToWater
		case "water-to-light":
			target = &garden.waterToLight
		case "light-to-temperature":
			target = &garden.lightToTemperature
		case "temperature-to-humidity":
			target = &garden.temperatureToHumidity
		case "humidity-to-location":
			target = &garden.humidityToLocation
		default:
			if target == nil {
				panic("invalid input")
			}
			*target = append(*target, parseRange(line))

		}
	}

	return garden
}

func parseRange(input string) mapping {
	parts := toInt(strings.Fields(input))
	if len(parts) != 3 {
		panic("invalid input")
	}

	return mapping{
		destination: parts[0],
		source:      parts[1],
		spread:      parts[2],
	}
}

func toInt(input []string) (output []int) {
	for _, value := range input {
		val, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		output = append(output, val)
	}
	return output
}

func getDestination(source int, mapping []mapping) int {
	for _, m := range mapping {
		if source >= m.source && source <= m.source+m.spread {
			return m.destination + (source - m.source)
		}
	}
	return source
}
