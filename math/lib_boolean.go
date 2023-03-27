// package math provides basic math functions for use with the template engine.
package math

// The NOT function returns the opposite of the boolean value passed to it.
// - a: the boolean value to negate
// - returns: the opposite of the boolean value passed to it
func NOT(a bool) bool {
	return !a
}

// `OR` the logical	or of its two boolean arguments
// - a: the first boolean value
// - b: the second boolean value
// - returns: `true` if either `a` or `b` is `true`, otherwise it returns `false`
func OR(a bool, b bool) bool {
	return a || b
}

// `AND` the logical AND of its two boolean arguments
// - a: the first boolean value
// - b: the second boolean value
// - returns: `true` if both `a` and `b` are `true`, otherwise it returns `false`
func AND(a bool, b bool) bool {
	return a && b
}
