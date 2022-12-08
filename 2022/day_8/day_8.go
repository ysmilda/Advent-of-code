package day_8

import (
	_ "embed"
	"strconv"
	"strings"

	solver "advent-of-code-2022/SolverInterface"
)

//go:embed input.txt
var inputFile string

type Solver struct {
	input Forest
}

type Forest struct {
	Grid [][]int
}

func (f Forest) CountVisible() int {
	count := 0
	for x, row := range f.Grid {
		for y := range row {
			if f.Visible(x, y) {
				count++
			}
		}
	}
	return count
}

func (f Forest) BestScenicScore() int {
	count := 0
	for x := 1; x < len(f.Grid)-1; x++ {
		for y := 1; y < len(f.Grid[0])-1; y++ {
			score := f.ScenicScore(x, y)
			if score > count {
				count = score
			}
		}
	}
	return count
}

func (f Forest) Visible(x, y int) bool {
	height := f.Grid[x][y]
	invisible := make([]bool, 4)
	for i := x - 1; i >= 0; i-- {
		if f.Grid[i][y] >= height {
			invisible[0] = true
			break
		}
	}
	for i := x + 1; i < len(f.Grid); i++ {
		if f.Grid[i][y] >= height {
			invisible[1] = true
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		if f.Grid[x][i] >= height {
			invisible[2] = true
			break
		}
	}
	for i := y + 1; i < len(f.Grid[0]); i++ {
		if f.Grid[x][i] >= height {
			invisible[3] = true
			break
		}
	}
	for _, val := range invisible {
		if !val {
			return true
		}
	}
	return false
}

func (f Forest) ScenicScore(x, y int) int {
	height := f.Grid[x][y]
	distance := make([]int, 4)
	for i := x - 1; i >= 0; i-- {
		distance[0]++
		if f.Grid[i][y] >= height {
			break
		}
	}
	for i := x + 1; i < len(f.Grid); i++ {
		distance[1]++
		if f.Grid[i][y] >= height {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		distance[2]++
		if f.Grid[x][i] >= height {
			break
		}
	}
	for i := y + 1; i < len(f.Grid[0]); i++ {
		distance[3]++
		if f.Grid[x][i] >= height {
			break
		}
	}
	count := 1
	for _, dist := range distance {
		if dist == 0 {
			continue
		}
		count *= dist
	}
	return count
}

func (s Solver) GetSolver() (solver.SolverInterface, error) {
	forest, err := parseInput()
	if err != nil {
		return Solver{}, err
	}

	return Solver{input: forest}, nil
}

func (s Solver) GetDay() int {
	return 8
}

func (s Solver) Part1() (string, error) {
	return strconv.Itoa(s.input.CountVisible()), nil
}

func (s Solver) Part2() (string, error) {
	return strconv.Itoa(s.input.BestScenicScore()), nil
}

func parseInput() (Forest, error) {
	var (
		lines  = strings.Split(inputFile, "\n")
		output = [][]int{}
	)
	for _, line := range lines {
		temp := []int{}
		for _, char := range line {
			i, err := strconv.Atoi(string(char))
			if err != nil {
				return Forest{}, err
			}
			temp = append(temp, i)
		}
		output = append(output, temp)
	}

	return Forest{Grid: output}, nil
}
