package aoc2023day11

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/aocmath"
	"github.com/ysmilda/Advent-of-code/foundation/grid"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	grid grid.Grid[bool]
}

func (s puzzle) CalculateDistances(step int) int {
	sum := 0

	emptyRows, emptyColumns := s.FindEmptyRowsAndColumns()
	galaxies := s.GetGalaxies()

	for i := 0; i < len(galaxies); i++ {
		for j := 0; j < len(galaxies); j++ {
			if galaxies[i] == galaxies[j] {
				continue
			}

			rowCount, columnCount := 0, 0
			for _, row := range emptyRows {
				if aocmath.Between(row, int(galaxies[i].Y), int(galaxies[j].Y)) {
					rowCount++
				}
			}
			for _, column := range emptyColumns {
				if aocmath.Between(column, int(galaxies[i].X), int(galaxies[j].X)) {
					columnCount++
				}
			}
			distance := int(galaxies[i].ManhattanDistance(galaxies[j]))
			distance += (rowCount + columnCount) * (step - 1)
			sum += distance
		}
		galaxies = append(galaxies[:i], galaxies[i+1:]...)
		i--
	}
	return sum
}

func (s puzzle) FindEmptyRowsAndColumns() (rows []int, columns []int) {
	for i := 0; i < int(s.grid.GetWidth()); i++ {
		row := s.grid.GetRow(uint(i))
		for _, galaxy := range row {
			if galaxy {
				goto skiprow
			}
		}
		rows = append(rows, i)
	skiprow:
	}

	for i := 0; i < int(s.grid.GetHeight()); i++ {
		column := s.grid.GetColumn(uint(i))
		for _, galaxy := range column {
			if galaxy {
				goto skipcolumn
			}
		}
		columns = append(columns, i)
	skipcolumn:
	}

	return rows, columns
}

func (s puzzle) GetGalaxies() []grid.Coordinate {
	galaxies := []grid.Coordinate{}
	for y, row := range s.grid {
		for x, galaxy := range row {
			if galaxy {
				galaxies = append(galaxies, grid.Coordinate{X: x, Y: y})
			}
		}
	}
	return galaxies
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
	return s.CalculateDistances(2), nil
}

func (s puzzle) Part2() (int, error) {
	return s.CalculateDistances(1000000), nil
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
