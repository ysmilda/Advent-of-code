package main

import (
	"time"

	aoc2023day1 "github.com/ysmilda/Advent-of-code/pkg/2023/day1"
	aoc2023day10 "github.com/ysmilda/Advent-of-code/pkg/2023/day10"
	aoc2023day2 "github.com/ysmilda/Advent-of-code/pkg/2023/day2"
	aoc2023day3 "github.com/ysmilda/Advent-of-code/pkg/2023/day3"
	aoc2023day4 "github.com/ysmilda/Advent-of-code/pkg/2023/day4"
	aoc2023day5 "github.com/ysmilda/Advent-of-code/pkg/2023/day5"
	aoc2023day6 "github.com/ysmilda/Advent-of-code/pkg/2023/day6"
	aoc2023day7 "github.com/ysmilda/Advent-of-code/pkg/2023/day7"
	aoc2023day8 "github.com/ysmilda/Advent-of-code/pkg/2023/day8"
	aoc2023day9 "github.com/ysmilda/Advent-of-code/pkg/2023/day9"
	"github.com/ysmilda/Advent-of-code/pkg/utils/solver"
)

func init() {
	start := time.Now()
	solvers = map[int]map[int]solver.Solver{
		2023: {
			1:  aoc2023day1.MustGetSolver(),
			2:  aoc2023day2.MustGetSolver(),
			3:  aoc2023day3.MustGetSolver(),
			4:  aoc2023day4.MustGetSolver(),
			5:  aoc2023day5.MustGetSolver(),
			6:  aoc2023day6.MustGetSolver(),
			7:  aoc2023day7.MustGetSolver(),
			8:  aoc2023day8.MustGetSolver(),
			9:  aoc2023day9.MustGetSolver(),
			10: aoc2023day10.MustGetSolver(),
		},
	}
	printParseTime(time.Since(start))
}
