// This package contains helper functions for that are related to converting interfaces to other types and
// checking the type of an interface
package convert

// IsString takes an interface and returns true if the interface is a string
// - value: the interface to check
// - returns: true if the interface is a string
func IsString(value interface{}) bool {
	switch value.(type) {
	case string:
		return true
	default:
		return false
	}
}

// IsBool takes an interface and returns true if the interface is a bool
// - value: the interface to check
// - returns: true if the interface is a bool
func IsBool(value interface{}) bool {
	switch value.(type) {
	case bool:
		return true
	default:
		return false
	}
}

// IsInt takes an interface and returns true if the interface is an int
// - value: the interface to check
// - returns: true if the interface is an int
func IsInt(value interface{}) bool {
	switch value.(type) {
	case int8, int16, int32, int64, int:
		return true
	default:
		return false
	}
}

// IsFloat takes an interface and returns true if the interface is a float
// - value: the interface to check
// - returns: true if the interface is a float
func IsFloat(value interface{}) bool {
	switch value.(type) {
	case float32, float64:
		return true
	default:
		return false
	}
}

// IsMapInterface takes an interface and returns true if the interface is a map[string]interface{}
// - value: the interface to check
// - returns: true if the interface is a map[string]interface{}
func IsMapInterface(value interface{}) bool {
	switch value.(type) {
	case map[string]interface{}:
		return true
	default:
		return false
	}
}
