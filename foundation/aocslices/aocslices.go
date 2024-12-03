package aocslices

func Remove[T []E, E any](slice T, index int) T {
	return append(slice[:index], slice[index+1:]...)
}
