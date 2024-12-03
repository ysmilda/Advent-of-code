package char

// IsDigit checks if a character is a digit.
func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// ToInt converts a character to an integer. It assumes that the character is a digit.
func ToInt(c byte) int {
	return int(c - '0')
}
