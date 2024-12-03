package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	sessionToken := flag.String("session-token", "", "the session token to use for downloading the input")
	flag.Parse()

	years, err := os.ReadDir("./solutions")
	if err != nil {
		panic(err)
	}

	for _, year := range years {
		if !year.IsDir() {
			continue
		}

		yearPath := filepath.Join("./solutions", year.Name())
		days, err := os.ReadDir(yearPath)
		if err != nil {
			panic(err)
		}

		for _, day := range days {
			dayPath := filepath.Join(yearPath, day.Name())
			dayFolderContents, err := os.ReadDir(dayPath)
			if err != nil {
				panic(err)
			}

			skip := false
			for _, file := range dayFolderContents {
				if file.Name() == "input.txt" {
					skip = true
				}
			}
			if skip {
				continue
			}

			yearInt, err := strconv.Atoi(year.Name())
			if err != nil {
				panic(err)
			}

			dayInt, err := strconv.Atoi(day.Name()[3:])
			if err != nil {
				panic(err)
			}

			fmt.Printf("Fetching input for %d day %d\n", yearInt, dayInt)

			req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", yearInt, dayInt), nil)
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

			inputFile, err := os.Create(dayPath + "/input.txt")
			if err != nil {
				panic(err)
			}

			input, err := io.ReadAll(resp.Body)
			if err != nil {
				inputFile.Close()
				panic(err)
			}

			if input[len(input)-1] == '\n' {
				input = input[:len(input)-1]
			}

			inputFile.Write(input)
			inputFile.Close()
		}
	}
}

func scan(path string) {

}
