package aoc2024day11

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/aocmath"
	"github.com/ysmilda/Advent-of-code/foundation/aocstrconv"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input []int
}

func MustGetSolver() solver.Solver {
	return puzzle{
		input: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 11
}

func (s puzzle) Part1() (int, error) {

	return blink(25, s.input), nil
}

func (s puzzle) Part2() (int, error) {
	return blink(75, s.input), nil
}

func parse(input string) []int {
	parts := strings.Split(input, " ")
	return aocstrconv.MustAtoiSlice(parts)
}

func blink(times int, in []int) int {
	stones := map[int]int{}
	next := map[int]int{}
	for _, in := range in {
		stones[in]++
	}

	for range times {
		for stone, value := range stones {
			if stone == 0 {
				next[1] += value
			} else {
				if aocmath.NumberOfDigits(stone)%2 == 0 {
					a, b := aocmath.Split(stone)
					next[a] += value
					next[b] += value
				} else {
					next[stone*2024] += value
				}
			}
		}

		stones = map[int]int{}
		for k, v := range next {
			stones[k] = v
		}
		next = map[int]int{}
	}

	sum := 0
	for _, count := range stones {
		sum += count
	}
	return sum
}
