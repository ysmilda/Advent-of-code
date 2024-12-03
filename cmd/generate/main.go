package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"text/template"
)

type definition struct {
	Day  uint
	Year uint
}

//go:embed day.tmpl
var dayTemplateContent string
var dayTemplate = template.Must(template.New("day").Parse(dayTemplateContent))

//go:embed day_test.tmpl
var dayTestTemplateContent string
var dayTestTemplate = template.Must(template.New("day_test").Parse(dayTestTemplateContent))

func main() {
	year := flag.Uint("year", 2024, "the year to generate for")
	day := flag.Uint("day", 1, "the day to generate for")
	flag.Parse()

	path := fmt.Sprintf("./solutions/%d/day%d", *year, *day)
	if _, err := os.ReadDir(path); err == nil {
		panic(fmt.Errorf("day %d already exists", *day))
	}

	err := os.MkdirAll(path, 0o755)
	if err != nil {
		panic(err)
	}

	solutionFile, err := os.Create(fmt.Sprintf("%s/day%d.go", path, *day))
	if err != nil {
		panic(err)
	}
	defer solutionFile.Close()

	testFile, err := os.Create(fmt.Sprintf("%s/day%d_test.go", path, *day))
	if err != nil {
		panic(err)
	}
	defer testFile.Close()

	def := definition{
		Day:  *day,
		Year: *year,
	}

	err = dayTemplate.Execute(solutionFile, def)
	if err != nil {
		panic(err)
	}

	err = dayTestTemplate.Execute(testFile, def)
	if err != nil {
		panic(err)
	}
}
