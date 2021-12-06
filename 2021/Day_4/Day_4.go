package day_4

import (
	solver "advent-of-code-2021/SolverInterface"
	_ "embed"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

type Solver struct {
	Input Bingo
}

type Bingo struct {
	Drawn  []int
	Boards []Board
}

func (b Bingo) FindFastestSolve() Board {
	smallest := Board{DrawsToComplete: math.MaxInt64}

	for _, board := range b.Boards {
		if board.DrawsToComplete < smallest.DrawsToComplete {
			smallest = board
		}
	}

	return smallest
}

type Board struct {
	Fields          [][]Field
	DrawsToComplete int
}

func (b *Board) Solve(input []int) {
	for index, value := range input {
		b.Mark(value)

		if b.CompleteRows() || b.CompleteColumns() {
			b.DrawsToComplete = index
			return
		}
	}
}

func (b *Board) Mark(value int) {
	for row := range b.Fields {
		for column := range b.Fields[0] {
			if b.Fields[row][column].Value == value {
				b.Fields[row][column].Marked = true
				return
			}
		}
	}
}

func (b *Board) CompleteRows() bool {
	for column := range b.Fields {
		complete := true
		for row := range b.Fields[0] {
			if !b.Fields[row][column].Marked {
				complete = false
				break
			}
		}

		if complete {
			return true
		}
	}

	return false
}

func (b *Board) CompleteColumns() bool {
	for row := range b.Fields {
		complete := true
		for column := range b.Fields[0] {
			if !b.Fields[row][column].Marked {
				complete = false
				break
			}
		}

		if complete {
			return true
		}
	}

	return false
}

func (b *Board) SumUnmarkedFields() int {
	output := 0

	for row := range b.Fields {
		for column := range b.Fields[0] {
			if !b.Fields[row][column].Marked {
				output += b.Fields[row][column].Value
			}
		}
	}

	return output
}

type Field struct {
	Value  int
	Marked bool
}

func (s Solver) GetSolver() (solver.SolverInterface, error) {
	input, err := parseInput()
	if err != nil {
		return Solver{}, err
	}

	return Solver{
		Input: input,
	}, nil
}

func (s Solver) GetDay() int {
	return 3
}

func (s Solver) Part1() (int, error) {
	for i := range s.Input.Boards {
		s.Input.Boards[i].Solve(s.Input.Drawn)
	}

	bestBoard := s.Input.FindFastestSolve()

	unmarkedSum := bestBoard.SumUnmarkedFields()

	return unmarkedSum * s.Input.Drawn[bestBoard.DrawsToComplete], nil
}

func (s Solver) Part2() (int, error) {

	return 0, nil
}

func parseInput() (Bingo, error) {
	lines := strings.Split(inputFile, "\n")

	lines = removeEmtpy(lines)

	output := Bingo{}

	draws := strings.Split(lines[0], ",")

	for _, draw := range draws {
		value, err := strconv.Atoi(draw)
		if err != nil {
			return Bingo{}, err
		}

		output.Drawn = append(output.Drawn, value)
	}

	for i := 1; i < len(lines); i += 5 {
		board := Board{
			Fields: make([][]Field, 0),
		}

		for j := 0; j < 5; j++ {
			boardInputs := strings.Split(lines[i+j], " ")

			boardInputs = removeEmtpy(boardInputs)

			fields := []Field{}

			for _, input := range boardInputs {
				value, err := strconv.Atoi(input)
				if err != nil {
					return Bingo{}, err
				}

				fields = append(fields, Field{
					Value: value,
				})
			}

			board.Fields = append(board.Fields, fields)
		}

		output.Boards = append(output.Boards, board)
	}

	return output, nil
}

func removeEmtpy(input []string) (output []string) {
	for _, value := range input {
		if value != "" {
			output = append(output, value)
		}
	}

	return output
}
