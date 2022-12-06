package day_4

import (
	_ "embed"
	"strconv"
	"strings"

	solver "advent-of-code-2022/SolverInterface"
)

//go:embed input.txt
var inputFile string

type Solver struct {
	Input []RangeGroup
}

type RangeGroup struct {
	rangeA Range
	rangeB Range
}

func (r RangeGroup) FullyContained() bool {
	return r.rangeA.min <= r.rangeB.min && r.rangeA.max >= r.rangeB.max ||
		r.rangeB.min <= r.rangeA.min && r.rangeB.max >= r.rangeA.max
}

func (r RangeGroup) Overlap() bool {
	return r.rangeA.Overlap(r.rangeB) || r.rangeB.Overlap(r.rangeA)
}

type Range struct {
	min, max int
}

func (r Range) Overlap(other Range) bool {
	return r.min >= other.min && r.min <= other.max || r.max >= other.min && r.max <= other.max
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
	return 4
}

func (s Solver) Part1() (string, error) {
	count := 0
	for _, ranges := range s.Input {
		if ranges.FullyContained() {
			count++
		}
	}
	return strconv.Itoa(count), nil
}

func (s Solver) Part2() (string, error) {
	count := 0
	for _, ranges := range s.Input {
		if ranges.Overlap() {
			count++
		}
	}
	return strconv.Itoa(count), nil
}

func parseInput() ([]RangeGroup, error) {
	lines := strings.Split(inputFile, "\n")

	output := make([]RangeGroup, len(lines))
	for i, line := range lines {
		ranges := strings.Split(line, ",")

		rangeA := strings.Split(ranges[0], "-")
		rangeAMin, _ := strconv.Atoi(rangeA[0])
		rangeAMax, _ := strconv.Atoi(rangeA[1])

		rangeB := strings.Split(ranges[1], "-")
		rangeBMin, _ := strconv.Atoi(rangeB[0])
		rangeBMax, _ := strconv.Atoi(rangeB[1])

		output[i] = RangeGroup{
			rangeA: Range{
				min: rangeAMin,
				max: rangeAMax,
			},
			rangeB: Range{
				min: rangeBMin,
				max: rangeBMax,
			},
		}
	}

	return output, nil
}
