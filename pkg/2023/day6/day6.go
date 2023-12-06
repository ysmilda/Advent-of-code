package aoc2023day6

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/ysmilda/Advent-of-code/pkg/solver"
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
	lowest := 0

	for i := r.distance / r.time; i < r.time; i++ {
		distance := (r.time - i) * i
		if distance > r.distance {
			lowest = i
			break
		}
	}

	return ((r.time - lowest) - lowest) + 1
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
		time:     toInt(tempTime),
		distance: toInt(tempDistance),
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
			time:     toInt(times[i]),
			distance: toInt(distances[i]),
		})
	}

	return output
}

func toInt(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return output
}
