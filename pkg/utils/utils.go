package utils

import "golang.org/x/exp/constraints"

func Min[T constraints.Integer | constraints.Float](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Integer | constraints.Float](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Limit[T constraints.Integer | constraints.Float](value, min, max T) T {
	return Min(Max(value, min), max)
}

func Abs[T constraints.Integer | constraints.Float](value T) T {
	if value < 0 {
		return -value
	}
	return value
}
