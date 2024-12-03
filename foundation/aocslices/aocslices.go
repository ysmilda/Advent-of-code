package aocslices

// Remove removes the element at index from a slice.
func Remove[T []E, E any](slice T, index int) T {
	return append(slice[:index], slice[index+1:]...)
}
