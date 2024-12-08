package aoc2024day8

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/grid"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	antennas map[byte][]grid.Coordinate
	size     grid.Coordinate
}

func MustGetSolver() solver.Solver {
	antennas, size := parse(inputFile)
	return puzzle{
		antennas: antennas,
		size:     size,
	}
}

func (s puzzle) GetDay() int {
	return 8
}

func (s puzzle) Part1() (int, error) {
	antinodes := make(map[grid.Coordinate]bool)

	for _, antennas := range s.antennas {
		if len(antennas) <= 1 {
			continue
		}
		for i, antenna := range antennas {
			for _, other := range antennas[i+1:] {
				diff := antenna.Subtract(other)

				if coord := antenna.Add(diff); coord.Within(s.size) {
					antinodes[coord] = true
				}
				if coord := other.Add(diff.Invert()); coord.Within(s.size) {
					antinodes[coord] = true
				}
			}
		}
	}

	return len(antinodes), nil
}

func (s puzzle) Part2() (int, error) {
	antinodes := make(map[grid.Coordinate]bool)

	for _, antennas := range s.antennas {
		if len(antennas) <= 1 {
			continue
		}
		for i, antenna := range antennas {
			antinodes[antenna] = true

			for _, other := range antennas[i+1:] {
				diff := antenna.Subtract(other)

				coord := antenna
				for {
					if coord = coord.Add(diff); coord.Within(s.size) {
						antinodes[coord] = true
					} else {
						break
					}
				}

				coord = other
				for {
					if coord = coord.Add(diff.Invert()); coord.Within(s.size) {
						antinodes[coord] = true
					} else {
						break
					}
				}
			}
		}
	}

	return len(antinodes), nil
}

func parse(input string) (map[byte][]grid.Coordinate, grid.Coordinate) {
	lines := strings.Split(input, "\n")
	antennas := make(map[byte][]grid.Coordinate)

	for y, line := range lines {
		for x, c := range line {
			if c == '.' {
				continue
			}

			antennas[byte(c)] = append(antennas[byte(c)], grid.NewCoordinate(x, y))
		}
	}

	return antennas, grid.NewCoordinate(len(lines[0]), len(lines))
}
