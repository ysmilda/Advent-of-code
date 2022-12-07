package main

import (
	"flag"
	"time"

	"advent-of-code-2022/day_1"
	"advent-of-code-2022/day_2"
	"advent-of-code-2022/day_3"
	"advent-of-code-2022/day_4"
	"advent-of-code-2022/day_5"
	"advent-of-code-2022/day_6"
	"advent-of-code-2022/day_7"

	solver "advent-of-code-2022/SolverInterface"

	log "github.com/sirupsen/logrus"
)

var solvers []solver.SolverInterface = []solver.SolverInterface{
	day_1.Solver{},
	day_2.Solver{},
	day_3.Solver{},
	day_4.Solver{},
	day_5.Solver{},
	day_6.Solver{},
	day_7.Solver{},
}

func main() {
	dayFlag := flag.Uint("d", 0, "the day to run, when set to 0 all days will be run")
	flag.Parse()

	// Normalize flag input
	day := int(*dayFlag)

	switch day {
	case 0:
		for _, solver := range solvers {
			executeSolver(solver)
		}

	default:
		for _, solver := range solvers {
			if solver.GetDay() == day {
				executeSolver(solver)
				return
			}
		}

		log.WithFields(log.Fields{
			"day": day,
		}).Fatal("Day not created yet")
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
