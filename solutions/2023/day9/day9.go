package aoc2023day9

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/char"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input []sequence
}

type sequence struct {
	input []int
}

func (s sequence) ExtrapolateEnd() int {
	ends := []int{}
	input := s.input
	for {
		next := []int{}
		allZero := true
		ends = append(ends, input[len(input)-1])
		for i := 1; i < len(input); i++ {
			value := input[i] - input[i-1]
			if value != 0 {
				allZero = false
			}
			next = append(next, value)
		}
		input = next
		if allZero {
			break
		}
	}

	sum := ends[len(ends)-1]
	for i := len(ends) - 2; i >= 0; i-- {
		sum += ends[i]
	}

	return sum
}

func (s sequence) ExtrapolateStart() int {
	starts := []int{}
	input := s.input
	for {
		next := []int{}
		allZero := true
		starts = append(starts, input[0])
		for i := 1; i < len(input); i++ {
			value := input[i-1] - input[i]
			if value != 0 {
				allZero = false
			}
			next = append(next, value)
		}
		input = next
		if allZero {
			break
		}
	}

	sum := starts[len(starts)-1]
	for i := len(starts) - 2; i >= 0; i-- {
		sum += starts[i]
	}

	return sum
}

func MustGetSolver() solver.Solver {
	return puzzle{
		input: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 9
}

func (s puzzle) Part1() (int, error) {
	sum := 0

	for _, seq := range s.input {
		sum += seq.ExtrapolateEnd()
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	for _, seq := range s.input {
		sum += seq.ExtrapolateStart()
	}

	return sum, nil
}

func parse(input string) (output []sequence) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		output = append(output, sequence{
			input: toInts(strings.Fields(line)),
		})
	}
	return output
}

func toInts(input []string) (output []int) {
	for _, i := range input {
		output = append(output, char.MustToInt(i))
	}
	return output
}
