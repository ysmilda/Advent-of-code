package day_3

import (
	solver "advent-of-code-2021/SolverInterface"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

type Solver struct {
	Input  []int
	Length int
}

func (s Solver) GetSolver() (solver.SolverInterface, error) {
	input, length, err := parseInput()
	if err != nil {
		return Solver{}, err
	}

	return Solver{
		Input:  input,
		Length: length,
	}, nil
}

func (s Solver) GetDay() int {
	return 3
}

func (s Solver) Part1() (int, error) {
	gammaRate := s.findGammaRate()
	epsilonRate := 0xFFFF ^ gammaRate

	epsilonRate &= ^(0xFFFF << s.Length)

	return gammaRate * epsilonRate, nil
}

func (s Solver) Part2() (int, error) {
	// gammaRate := s.findGammaRate()

	// for i := s.Length - 1; i >= 0; i-- {

	// }

	return 0, nil
}

func (s Solver) findGammaRate() int {
	gammaRate := 0

	for i := 0; i < s.Length; i++ {
		count := 0
		for j := 0; j < len(s.Input); j++ {
			if (s.Input[j]>>i)&0x1 == 1 {
				count++
			}
		}

		if count > (len(s.Input) - count) {
			gammaRate += 1 << i
		}
	}

	return gammaRate
}

func parseInput() ([]int, int, error) {
	lines := strings.Split(inputFile, "\n")

	output := []int{}

	for _, line := range lines {
		value, err := strconv.ParseInt(line, 2, 64)
		if err != nil {
			return nil, 0, err
		}

		output = append(output, int(value))
	}

	return output, len(lines[0]), nil
}
