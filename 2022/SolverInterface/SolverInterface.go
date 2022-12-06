package solver

type SolverInterface interface {
	GetSolver() (SolverInterface, error)
	// ParseInput(string) (SolverInterface, error)

	GetDay() int

	Part1() (string, error)

	Part2() (string, error)
}
