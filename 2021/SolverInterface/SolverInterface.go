package solver

type SolverInterface interface {
	GetSolver() (SolverInterface, error)

	GetDay() int

	Part1() (int, error)

	Part2() (int, error)
}
