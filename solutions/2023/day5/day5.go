package aoc2023day5

import (
	_ "embed"
	"strings"
	"sync"

	"github.com/ysmilda/Advent-of-code/foundation/aocstrconv"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input garden
}

type garden struct {
	seeds    []int
	mappings [][]mapping
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
		input := seed
		for _, m := range s.input.mappings {
			input = getDestination(input, m)
		}

		if input < lowest {
			lowest = input
		}
	}

	return lowest, nil
}

func (s puzzle) Part2() (int, error) {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	lowest := 1 << 32

	for i := 0; i < len(s.input.seeds); i += 2 {
		wg.Add(1)
		go func(seed, rng int) {
			localLowest := 1 << 32
			for j := 0; j < rng; j++ {
				input := seed + j
				for _, m := range s.input.mappings {
					input = getDestination(input, m)
				}

				if input < localLowest {
					localLowest = input
				}
			}

			mu.Lock()
			if localLowest < lowest {
				lowest = localLowest
			}
			mu.Unlock()
			wg.Done()
		}(s.input.seeds[i], s.input.seeds[i+1])
	}

	wg.Wait()

	return lowest, nil
}

func parse(input string) garden {
	lines := strings.Split(input, "\n")

	garden := garden{
		seeds: aocstrconv.MustAtoiSlice(strings.Fields(strings.Split(lines[0], ":")[1])),
	}

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		if line[len(line)-1] == ':' {
			garden.mappings = append(garden.mappings, []mapping{})
			continue
		}
		garden.mappings[len(garden.mappings)-1] = append(garden.mappings[len(garden.mappings)-1], parseRange(line))
	}

	return garden
}

func parseRange(input string) mapping {
	parts := aocstrconv.MustAtoiSlice(strings.Fields(input))
	if len(parts) != 3 {
		panic("invalid input")
	}

	return mapping{
		destination: parts[0],
		source:      parts[1],
		spread:      parts[2],
	}
}

func getDestination(source int, mapping []mapping) int {
	for _, m := range mapping {
		if source >= m.source && source <= m.source+m.spread {
			return m.destination + (source - m.source)
		}
	}
	return source
}
