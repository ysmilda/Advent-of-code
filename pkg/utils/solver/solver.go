package solver

type Solver interface {
	GetDay() int

	Part1() (int, error)

	Part2() (int, error)
}
