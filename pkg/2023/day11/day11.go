package aoc2023day11

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/pkg/utils/grid"
	"github.com/ysmilda/Advent-of-code/pkg/utils/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	grid grid.Grid[bool]
}

func (s puzzle) ExpandGrid(step int) grid.Grid[bool] {
	g := s.grid
	for i := 0; i < int(g.GetWidth()); i++ {
		row := g.GetRow(uint(i))
		for _, galaxy := range row {
			if galaxy {
				goto skiprow
			}
		}
		g = g.AddRow(make([]bool, g.GetWidth()), uint(i))
		i++
	skiprow:
	}

	for i := 0; i < int(g.GetHeight()); i++ {
		column := g.GetColumn(uint(i))
		for _, galaxy := range column {
			if galaxy {
				goto skipcolumn
			}
		}
		g = g.AddColumn(make([]bool, g.GetHeight()), uint(i))
		i++
	skipcolumn:
	}
	return g
}

func MustGetSolver() solver.Solver {
	return puzzle{
		grid: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 11
}

func (s puzzle) Part1() (int, error) {
	sum := 0

	g := s.ExpandGrid(1)
	galaxies := make([]grid.Coordinate, 0)
	for y, row := range g {
		for x, galaxy := range row {
			if galaxy {
				galaxies = append(galaxies, grid.Coordinate{X: x, Y: y})
			}
		}
	}
	for i := 0; i < len(galaxies); i++ {
		for j := 0; j < len(galaxies); j++ {
			if galaxies[i] == galaxies[j] {
				continue
			}
			distance := int(galaxies[i].ManhattanDistance(galaxies[j]))
			sum += distance
		}
		galaxies = append(galaxies[:i], galaxies[i+1:]...)
		i--
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	return sum, nil
}

func parse(input string) grid.Grid[bool] {
	lines := strings.Split(input, "\n")
	output := grid.NewGrid[bool](uint(len(lines[0])), uint(len(lines)))

	for i, line := range lines {
		for j, char := range line {
			output.Set(uint(j), uint(i), char == '#')
		}
	}
	return output
}
