package aocstrconv

import "strconv"

// MustAtoi converts a string to an integer. It panics if the conversion fails.
func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// MustAtoiSlice converts a slice of strings to a slice of integers. It panics if the conversion fails.
func MustAtoiSlice(s []string) []int {
	var i []int
	for _, v := range s {
		i = append(i, MustAtoi(v))
	}
	return i
}
