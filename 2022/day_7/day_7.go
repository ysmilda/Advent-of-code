package day_7

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	solver "advent-of-code-2022/SolverInterface"
)

//go:embed input.txt
var inputFile string

type Solver struct {
	input Directory
}

func NewRootDirectory() Directory {
	return Directory{
		Name:   "/",
		Parent: nil,
	}
}

type Directory struct {
	Name        string
	Parent      *Directory
	Directories []Directory
	Files       []File
}

func (d Directory) FindDir(name string) (*Directory, error) {
	for i := 0; i < len(d.Directories); i++ {
		if d.Directories[i].Name == name {
			return &d.Directories[i], nil
		}
	}

	return nil, fmt.Errorf("directory %s does not exist in folder %s", name, d.Name)
}

func (d *Directory) AddDir(name string) {
	for _, directory := range d.Directories {
		if directory.Name == name {
			return
		}
	}
	d.Directories = append(d.Directories, Directory{
		Name:   name,
		Parent: d,
	})
}

func (d *Directory) AddFile(name string, size int) {
	for _, file := range d.Files {
		if file.Name == name {
			return
		}
	}
	d.Files = append(d.Files, File{
		Name: name,
		Size: size,
	})
}

func (d Directory) Size() int {
	size := 0
	for _, directory := range d.Directories {
		size += directory.Size()
	}
	for _, file := range d.Files {
		size += file.Size
	}
	return size
}

func (d Directory) Sizes() []int {
	output := []int{}
	for _, directory := range d.Directories {
		output = append(output, directory.Sizes()...)
	}
	output = append(output, d.Size())
	return output
}

func (d Directory) FindFoldersUnderSize(folderSize int) int {
	score := 0
	for _, directory := range d.Directories {
		score += directory.FindFoldersUnderSize(folderSize)
	}
	size := d.Size()
	if size <= folderSize {
		score += size
	}
	return score
}

func (d Directory) FindFolderToDelete(totalSpace, neededSpace int) int {
	spaceToClear := totalSpace - neededSpace - d.Size()

	// Find smallest folder that fits the needed space
	sizes := d.Sizes()
	smallest := math.MaxInt
	for _, size := range sizes {
		if size >= spaceToClear && size < smallest {
			smallest = size
		}
	}
	return smallest
}

type File struct {
	Name string
	Size int
}

func (s Solver) GetSolver() (solver.SolverInterface, error) {
	directory, err := parseInput()
	if err != nil {
		return Solver{}, err
	}

	return Solver{input: directory}, nil
}

func (s Solver) GetDay() int {
	return 7
}

func (s Solver) Part1() (string, error) {
	return strconv.Itoa(s.input.FindFoldersUnderSize(100000)), nil
}

func (s Solver) Part2() (string, error) {
	return strconv.Itoa(s.input.FindFolderToDelete(70000000, 30000000)), nil
}

func parseInput() (Directory, error) {
	var (
		lines         = strings.Split(inputFile, "\n")
		output        = NewRootDirectory()
		currentFolder = &output
		err           error
	)

	for i := 0; i < len(lines); i++ {
		switch {
		case lines[i] == "$ cd /":
			currentFolder = &output

		case lines[i] == "$ ls":
			for j := 1; j < len(lines)-i; j++ {
				if strings.HasPrefix(lines[i+j], "$") {
					i += j - 1
					break // Stop the loop on the next command
				}
				parts := strings.Split(lines[i+j], " ")
				switch parts[0] {
				case "dir":
					currentFolder.AddDir(parts[1])
				default:
					size, err := strconv.Atoi(parts[0])
					if err != nil {
						return Directory{}, err
					}
					currentFolder.AddFile(parts[1], size)
				}
			}
		case strings.HasPrefix(lines[i], "$ cd"):
			parts := strings.Split(lines[i], " ")
			if parts[2] == ".." {
				currentFolder = currentFolder.Parent
			} else {
				currentFolder, err = currentFolder.FindDir(parts[2])
				if err != nil {
					return Directory{}, err
				}
			}

		}
	}

	return output, nil
}
