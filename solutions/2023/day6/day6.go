package aoc2023day6

import (
	_ "embed"
	"math"
	"strconv"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/char"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input []race
}

type race struct {
	time     int
	distance int
}

func (r race) findNumberOfWaysToWin() int {
	t := float64(r.time)
	d := float64(r.distance)

	dSqrt := math.Sqrt((t * t) - (4 * d))
	p1 := (-t + dSqrt) / -2
	p2 := (-t - dSqrt) / -2
	return int(math.Ceil(p2) - math.Floor(p1) - 1)
}

func MustGetSolver() solver.Solver {
	return puzzle{
		input: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 6
}

func (s puzzle) Part1() (int, error) {
	sum := 1

	for _, r := range s.input {
		sum *= r.findNumberOfWaysToWin()
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	// Seemed easier than reparsing the input, not the fastest solution though
	tempTime := ""
	tempDistance := ""

	for _, r := range s.input {
		tempTime += strconv.Itoa(r.time)
		tempDistance += strconv.Itoa(r.distance)
	}

	combinedRace := race{
		time:     char.MustToInt(tempTime),
		distance: char.MustToInt(tempDistance),
	}

	return combinedRace.findNumberOfWaysToWin(), nil
}

func parse(input string) (output []race) {
	lines := strings.Split(input, "\n")
	times := strings.Fields(lines[0])
	distances := strings.Fields(lines[1])
	if len(times) != len(distances) {
		panic("invalid input")
	}

	for i := range times {
		if i == 0 {
			continue
		}
		output = append(output, race{
			time:     char.MustToInt(times[i]),
			distance: char.MustToInt(distances[i]),
		})
	}

	return output
}
