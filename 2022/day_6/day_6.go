package day_6

import (
	_ "embed"
	"strconv"

	solver "advent-of-code-2022/SolverInterface"
)

//go:embed input.txt
var inputFile string

type Solver struct {
	input string
}

func (s Solver) GetSolver() (solver.SolverInterface, error) {
	return Solver{input: inputFile}, nil
}

func (s Solver) GetDay() int {
	return 6
}

func (s Solver) Part1() (string, error) {
	return strconv.Itoa(findFirstUniqueGroup(s.input, 4)), nil
}

func (s Solver) Part2() (string, error) {
	return strconv.Itoa(findFirstUniqueGroup(s.input, 14)), nil
}

func findFirstUniqueGroup(input string, size int) int {
	for i := size; i < len(input); i++ {
		seen := make(map[rune]bool)
		for _, char := range input[i-size : i] {
			if seen[char] {
				goto skip
			}

			seen[char] = true
		}
		return i
	skip:
	}

	return 0
}
