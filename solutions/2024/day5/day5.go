package aoc2024day5

import (
	_ "embed"
	"slices"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/aocstrconv"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type rule struct {
	first, last int
}

type input struct {
	rules   map[rule]struct{}
	updates [][]int
}

type puzzle struct {
	input input
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
	sum := 0

	for _, update := range s.input.updates {
		sorted := make([]int, len(update))
		copy(sorted, update)
		slices.SortFunc(sorted, s.input.compare)

		if slices.Equal(sorted, update) {
			sum += sorted[len(sorted)/2]
		}
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	for _, update := range s.input.updates {
		sorted := make([]int, len(update))
		copy(sorted, update)
		slices.SortFunc(sorted, s.input.compare)

		if !slices.Equal(sorted, update) {
			sum += sorted[len(sorted)/2]
		}
	}

	return sum, nil
}

func (i input) compare(x, y int) int {
	if _, ok := i.rules[rule{x, y}]; ok {
		return -1
	} else if _, ok := i.rules[rule{y, x}]; ok {
		return 1
	}
	return 0
}

func parse(in string) input {
	lines := strings.Split(in, "\n")

	result := input{
		rules: make(map[rule]struct{}),
	}
	for i, line := range lines {
		if line == "" {
			lines = lines[i+1:]
			break
		}

		parts := strings.Split(line, "|")
		rule := rule{
			aocstrconv.MustAtoi(parts[0]),
			aocstrconv.MustAtoi(parts[1]),
		}
		result.rules[rule] = struct{}{}
	}

	for _, line := range lines {
		if line == "" {
			break
		}

		parts := strings.Split(line, ",")
		update := make([]int, len(parts))
		for i := range parts {
			update[i] = aocstrconv.MustAtoi(parts[i])
		}
		result.updates = append(result.updates, update)
	}

	return result
}
