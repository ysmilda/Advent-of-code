package aoc2023day1

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/char"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input []string
}

func MustGetSolver() solver.Solver {
	return puzzle{
		input: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 1
}

func (s puzzle) Part1() (int, error) {
	sum := 0
	for _, line := range s.input {
		digits := findDigits(line)
		sum += digits[0]*10 + digits[len(digits)-1]
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	for _, line := range s.input {
		line := replace(line)
		digits := findDigits(line)
		sum += digits[0]*10 + digits[len(digits)-1]
	}
	return sum, nil
}

func findDigits(line string) []int {
	digits := []int{}
	for i := 0; i < len(line); i++ {
		c := line[i]
		if char.IsDigit(c) {
			digits = append(digits, char.ToInt(c))
		}
	}
	return digits
}

func replace(line string) string {
	length := len(line)
	for i := 0; i < length; i++ {
		if length-i >= 3 {
			switch line[i : i+3] {
			case "one":
				line = line[:i] + "1" + line[i+1:]
			case "two":
				line = line[:i] + "2" + line[i+1:]
			case "six":
				line = line[:i] + "6" + line[i+1:]
			}
		}

		if length-i >= 4 {
			switch line[i : i+4] {
			case "four":
				line = line[:i] + "4" + line[i+1:]
			case "five":
				line = line[:i] + "5" + line[i+1:]
			case "nine":
				line = line[:i] + "9" + line[i+1:]
			}
		}

		if length-i >= 5 {
			switch line[i : i+5] {
			case "three":
				line = line[:i] + "3" + line[i+1:]
			case "seven":
				line = line[:i] + "7" + line[i+1:]
			case "eight":
				line = line[:i] + "8" + line[i+1:]
			}
		}
	}
	return line
}

func parse(input string) []string {
	return strings.Split(input, "\n")
}
