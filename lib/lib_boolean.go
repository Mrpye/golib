package lib

// Not returns the opposite of a bool
func NOT(a bool) bool {
	return !a
}

// OR returns the OR of two bools
func OR(a bool, b bool) bool {
	return a || b
}

// AND returns the AND of two bools
func AND(a bool, b bool) bool {
	return a && b
}
