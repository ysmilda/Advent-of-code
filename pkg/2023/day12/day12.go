package aoc2023day12

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/pkg/utils/char"
	"github.com/ysmilda/Advent-of-code/pkg/utils/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input []hotspring
}

type hotspring struct {
	condition string
	groups    []int
}

func (h hotspring) Arrangements() int {
	existingGroups := strings.Split(h.condition, ".")

	return 0
}

func (h hotspring) Validate() bool {
	indexDamagedGroups := 0
	count := 0

	for _, char := range h.condition {
		switch char {
		case '#':
			count++
			if count > h.groups[indexDamagedGroups] {
				return false
			}

		case '.':
			if count == h.groups[indexDamagedGroups] {
				indexDamagedGroups++
			}
			count = 0
		}
	}
	return true
}

func MustGetSolver() solver.Solver {
	return puzzle{
		input: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 12
}

func (s puzzle) Part1() (int, error) {
	sum := 0

	for _, hotspring := range s.input {
		sum += hotspring.Arrangements()
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	return sum, nil
}

func parse(input string) []hotspring {
	lines := strings.Split(input, "\n")
	hotsprings := []hotspring{}

	for _, line := range lines {
		parts := strings.Fields(line)
		condition := parts[0]
		damagedGroups := []int{}

		groups := strings.Split(parts[1], ",")
		for _, group := range groups {
			damagedGroups = append(damagedGroups, char.MustToInt(group))
		}

		hotsprings = append(hotsprings, hotspring{
			condition: condition,
			groups:    damagedGroups,
		})
	}
	return hotsprings
}
