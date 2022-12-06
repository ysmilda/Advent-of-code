package day_1

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"

	solver "advent-of-code-2022/SolverInterface"
)

//go:embed input.txt
var inputFile string

type Solver struct {
	Input []int
}

func (s Solver) GetSolver() (solver.SolverInterface, error) {
	input, err := parseInput()
	if err != nil {
		return Solver{}, err
	}

	return Solver{
		Input: input,
	}, nil
}

func (s Solver) GetDay() int {
	return 1
}

func (s Solver) Part1() (string, error) {
	return strconv.Itoa(s.Input[len(s.Input)-1]), nil
}

func (s Solver) Part2() (string, error) {
	lastIndex := len(s.Input) - 1
	count := s.Input[lastIndex] + s.Input[lastIndex-1] + s.Input[lastIndex-2]
	return strconv.Itoa(count), nil
}

func parseInput() ([]int, error) {
	lines := strings.Split(inputFile, "\n")

	index := 0
	output := []int{0}
	for _, line := range lines {
		if line == "" {
			output = append(output, 0)
			index = len(output) - 1
			continue
		}
		value, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		output[index] += value
	}

	sort.Ints(output)

	return output, nil
}
