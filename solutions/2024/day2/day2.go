package aoc2024day2

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/aocmath"
	"github.com/ysmilda/Advent-of-code/foundation/aocslices"
	"github.com/ysmilda/Advent-of-code/foundation/aocstrconv"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	reports [][]int
}

func MustGetSolver() solver.Solver {
	return puzzle{
		reports: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 2
}

func (s puzzle) Part1() (int, error) {
	sum := 0

	for _, report := range s.reports {
		if isSafe(report) {
			sum++
		}
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	for _, report := range s.reports {
		if isSafe(report) {
			sum++
			continue
		}

		for i := range report {
			if isSafe(aocslices.Remove(append([]int{}, report...), i)) {
				sum++
				break
			}
		}
	}

	return sum, nil
}

func parse(input string) (reports [][]int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		reports = append(reports, aocstrconv.MustAtoiSlice(strings.Split(line, " ")))
	}

	return reports
}

func isSafe(report []int) bool {
	direction := 0
	unsafe := false

	for i := 0; i < len(report)-1; i++ {
		step := report[i] - report[i+1]

		if direction == 0 {
			direction = step
		}

		// Verify the direction and stepsize.
		if !aocmath.SameSign(step, direction) || !aocmath.BetweenInclusive(aocmath.Abs(step), 1, 3) {
			unsafe = true
			break
		}
	}
	return !unsafe
}
