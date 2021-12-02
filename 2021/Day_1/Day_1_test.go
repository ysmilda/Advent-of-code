package day_1

import (
	"testing"
)

type input struct {
	solve    Solver
	outcome1 int
	outcome2 int
}

var inputs []input = []input{
	{
		solve: Solver{
			Input: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
		},
		outcome1: 7,
		outcome2: 5,
	},
}

func TestPart1(t *testing.T) {
	for index, input := range inputs {
		outcome, err := input.solve.Part1()

		if err != nil {
			t.Errorf("Failed to solve part 1 for input %d: %s", index, err.Error())
		}

		if outcome != input.outcome1 {
			t.Errorf("Expected %d, but got %d", input.outcome1, outcome)
		}
	}
}

func TestPart2(t *testing.T) {
	for index, input := range inputs {
		outcome, err := input.solve.Part2()

		if err != nil {
			t.Errorf("Failed to solve part 1 for input %d: %s", index, err.Error())
		}

		if outcome != input.outcome2 {
			t.Errorf("Expected %d, but got %d", input.outcome2, outcome)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inputs[0].solve.Part1()
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inputs[0].solve.Part2()
	}
}
