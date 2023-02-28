// package math provides basic math functions for use with the template engine.
package math

// The NOT function returns the opposite of the boolean value passed to it.
func NOT(a bool) bool {
	return !a
}

// `OR` returns `true` if either `a` or `b` is `true`, otherwise it returns `false`
func OR(a bool, b bool) bool {
	return a || b
}

// `AND` returns the logical AND of its two boolean arguments
func AND(a bool, b bool) bool {
	return a && b
}
