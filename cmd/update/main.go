package main

import (
	_ "embed"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"slices"

	"github.com/ysmilda/Advent-of-code/foundation/aocstrconv"
)

//go:embed days.tmpl
var daysTemplateContent string
var daysTemplate = template.Must(template.New("days").Parse(daysTemplateContent))

func main() {

	// Update the main days.go file to include all defined days.
	solutions := make(map[int][]int)

	years, err := os.ReadDir("./solutions")
	if err != nil {
		panic(err)
	}

	for _, year := range years {
		if !year.IsDir() {
			continue
		}

		yearInt := aocstrconv.MustAtoi(year.Name())

		yearPath := filepath.Join("./solutions", year.Name())
		days, err := os.ReadDir(yearPath)
		if err != nil {
			panic(err)
		}

		for _, day := range days {
			if !day.IsDir() {
				continue
			}

			dayInt := aocstrconv.MustAtoi(day.Name()[3:])
			solutions[yearInt] = append(solutions[yearInt], dayInt)
		}
	}

	for year := range solutions {
		slices.Sort(solutions[year])
	}

	daysFile, err := os.Create("./days.go")
	if err != nil {
		panic(err)
	}
	daysTemplate.Execute(daysFile, solutions)

	err = exec.Command("gofmt", "-w", "days.go").Run()
	if err != nil {
		panic(err)
	}
}
