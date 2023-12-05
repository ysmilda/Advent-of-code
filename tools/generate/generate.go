package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io"
	"net/http"
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
	year := flag.Uint("year", 2023, "the year to generate for")
	day := flag.Uint("day", 1, "the day to generate for")
	sessionToken := flag.String("session-token", "", "the session token to use for downloading the input")
	flag.Parse()

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", *year, *day), nil)
	if err != nil {
		panic(err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: *sessionToken,
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	path := fmt.Sprintf("./pkg/%d/day%d", *year, *day)
	err = os.MkdirAll(path, 0o755)
	if err != nil {
		panic(err)
	}

	inputFile, err := os.Create(path + "/input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	input, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if input[len(input)-1] == '\n' {
		input = input[:len(input)-1]
	}

	inputFile.Write(input)

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

	testInputFile, err := os.Create(fmt.Sprintf("%s/testinput.txt", path))
	if err != nil {
		panic(err)
	}
	defer testInputFile.Close()
}
