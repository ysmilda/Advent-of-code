package main

import (
	day_1 "advent-of-code-2021/Day_1"
	solver "advent-of-code-2021/SolverInterface"
	"flag"
	"time"

	log "github.com/sirupsen/logrus"
)

var solvers []solver.SolverInterface = []solver.SolverInterface{
	day_1.Solver{},
}

func main() {
	solverFlag := flag.Uint("s", 0, "the solver to run, when set to 0 all solvers will be run")
	flag.Parse()

	// Normalize flag input
	solverIndex := int(*solverFlag)

	// Program logic
	// Either do nothing because the requested day is not created yet
	// Or execute all solvers if the requested day == 0
	// Or execute the requested day
	if solverIndex > len(solvers) {
		log.WithFields(log.Fields{
			"solver": solverIndex,
		}).Fatal("Solution not created yet")

	} else if solverIndex == 0 {
		for _, solver := range solvers {
			executeSolver(solver)
		}

	} else {
		executeSolver(solvers[solverIndex-1])
	}
}

func executeSolver(solver solver.SolverInterface) {
	start := time.Now()

	input, err := solver.GetSolver()
	if err != nil {
		log.WithFields(log.Fields{
			"day":   solver.GetDay(),
			"error": err,
		}).Fatal("Failed to parse input")
	}

	part1, err := input.Part1()
	if err != nil {
		log.WithFields(log.Fields{
			"day":   solver.GetDay(),
			"error": err,
		}).Info("Failed to solve part 1")
	}

	part2, err := input.Part2()
	if err != nil {
		log.WithFields(log.Fields{
			"day":   solver.GetDay(),
			"error": err,
		}).Info("Failed to solve part 2")
	}

	log.WithFields(log.Fields{
		"day":     solver.GetDay(),
		"part 1":  part1,
		"part 2":  part2,
		"runtime": time.Since(start).String(),
	}).Info("Solution found")
}
