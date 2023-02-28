// This package contains helper functions for that are related to converting interfaces to other types.
package convert

import (
	"reflect"
	"strconv"
	"strings"
)

// `ToPassFail` takes a boolean value and returns a string value of either "Pass" or "Fail" depending
// on the boolean value
func ToPassFail(value bool) string {
	if value {
		return "Pass"
	}
	return "Fail"
}

// If the value is a string, it will return true if the string is "t", "true", "yes", "y", "1", "pass",
// or "0". Otherwise, it will return false
func ToBool(value interface{}) bool {
	switch val := value.(type) {
	case bool:
		return val
	case string:
		if strings.ToLower(val) == "t" {
			return true
		} else if strings.ToLower(val) == "f" {
			return false
		} else if strings.ToLower(val) == "yes" {
			return true
		} else if strings.ToLower(val) == "no" {
			return false
		} else if strings.ToLower(val) == "y" {
			return true
		} else if strings.ToLower(val) == "n" {
			return false
		} else if strings.ToLower(val) == "1" {
			return true
		} else if strings.ToLower(val) == "pass" {
			return true
		} else if strings.ToLower(val) == "fail" {
			return false
		} else if strings.ToLower(val) == "0" {
			return false
		}
		b, _ := strconv.ParseBool(val)
		return b
	case int:
		return val > 0
	case int8:
		return val > 0
	case int32:
		return val > 0
	case int64:
		return val > 0
	case float32:
		return val > 0
	case float64:
		return val > 0
	default:
		return false
	}
}

// It takes an interface{} and returns a string
func ToString(value interface{}) string {

	switch val := value.(type) {
	case bool:
		return strconv.FormatBool(val)
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	default:
		return ""
	}
}

// If the value is a string, convert it to an int64, otherwise convert it to an int
func ToInt(value interface{}) int {
	switch val := value.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		//convert the string to a int64
		if val == "" {
			return 0
		}
		if strings.Contains(val, ".") {
			f, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return 0
			}
			return int(f)
		}
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0
		}
		return int(i)
	case int:
		return val
	case float32:
		return int(val)
	case float64:
		return int(val)
	case int64:
		return int(val)
	case int8:
		return int(val)
	default:
		return 0
	}
}

// It converts any value to an int8
func ToInt8(value interface{}) int8 {
	switch val := value.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		//convert the string to a int64
		if val == "" {
			return 0
		}
		if strings.Contains(val, ".") {
			f, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return 0
			}
			return int8(f)
		}
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0
		}
		return int8(i)
	case int:
		return int8(val)
	case float32:
		return int8(val)
	case float64:
		return int8(val)
	case int64:
		return int8(val)
	case int8:
		return val
	default:
		return 0
	}
}

// If the value is a string, convert it to an int32. If the value is a float, convert it to an int32.
// If the value is a bool, convert it to an int32. If the value is an int, convert it to an int32. If
// the value is an int64, convert it to an int32. If the value is an int32, return it. If the value is
// anything else, return 0
func ToInt32(value interface{}) int32 {
	switch val := value.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		//convert the string to a int64
		if val == "" {
			return 0
		}
		if strings.Contains(val, ".") {
			f, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return 0
			}
			return int32(f)
		}
		i, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return 0
		}
		return int32(i)
	case int:
		return int32(val)
	case float32:
		return int32(val)
	case float64:
		return int32(val)
	case int64:
		return int32(val)
	case int32:
		return val
	default:
		return 0
	}
}

// If the value is a string, convert it to an int64. If the value is a float, convert it to an int64.
// If the value is a bool, convert it to an int64. If the value is an int, convert it to an int64. If
// the value is an int64, return it. If the value is anything else, return 0
func ToInt64(value interface{}) int64 {
	switch val := value.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		//convert the string to a int64
		if val == "" {
			return 0
		}
		if strings.Contains(val, ".") {
			f, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return 0
			}
			return int64(f)
		}
		i, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return 0
		}
		return int64(i)
	case int:
		return int64(val)
	case float32:
		return int64(val)
	case float64:
		return int64(val)
	case int64:
		return val
	case int32:
		return int64(val)
	default:
		return 0
	}
}

// If the value is a string, convert it to a float32. Otherwise, return the value as a float32
func ToFloat32(value interface{}) float32 {
	switch val := value.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		//convert the string to a float32
		if val == "" {
			return 0
		}
		f, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return 0
		}
		return float32(f)
	case int:
		return float32(val)
	case float32:
		return val
	case float64:
		return float32(val)
	default:
		return 0
	}
}

// If the value is a string, convert it to a float64. If the value is a bool, convert it to a float64.
// If the value is an int, convert it to a float64. If the value is a float64, return it. Otherwise,
// return 0
func ToFloat64(value interface{}) float64 {
	switch val := value.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		//convert the string to a float64
		if val == "" {
			return 0
		}
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0
		}
		return f
	case int:
		return float64(val)
	case float64:
		return val
	default:
		return 0
	}
}

// If the value is a map[string]interface{}, return it, otherwise return nil
func ToMapInterface(value interface{}) map[string]interface{} {
	switch val := value.(type) {
	case map[string]interface{}:
		return val
	default:
		return nil
	}
}

// It takes a map of strings to interfaces and returns a map of strings to strings
func ToMapString(value interface{}) map[string]string {
	switch val := value.(type) {
	case map[string]interface{}:
		//convert the map to a map of strings
		m := make(map[string]string)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.String {
				m[k] = ToString(v)
			}
		}
		return m
	case map[string]bool:
		//convert the map to a map of strings
		m := make(map[string]string)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Int {
				m[k] = ToString(v)
			}
		}
		return m
	case map[string]float64:
		//convert the map to a map of strings
		m := make(map[string]string)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Int {
				m[k] = ToString(v)
			}
		}
		return m
	case map[string]int:
		//convert the map to a map of strings
		m := make(map[string]string)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Int {
				m[k] = ToString(v)
			}
		}
		return m
	case map[string]int64:
		//convert the map to a map of strings
		m := make(map[string]string)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Int {
				m[k] = ToString(v)
			}
		}
		return m
	case map[string]string:
		return val
	default:
		return nil
	}
}

// It takes a map of strings to interfaces, and returns a map of strings to ints
func ToMapInt(value interface{}) map[string]int {
	switch val := value.(type) {
	case map[string]interface{}:
		//convert the map to a map of strings
		m := make(map[string]int)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Int {
				m[k] = ToInt(v)
			}
		}
		return m
	case map[string]int:
		//convert the map to a map of strings
		return val
	case map[string]int64:
		//convert the map to a map of strings
		m := make(map[string]int)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Int {
				m[k] = ToInt(v)
			}
		}
		return m
	case map[string]float64:
		//convert the map to a map of strings
		m := make(map[string]int)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Int {
				m[k] = ToInt(v)
			}
		}
		return m
	case map[string]bool:
		//convert the map to a map of strings
		m := make(map[string]int)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Int {
				m[k] = ToInt(v)
			}
		}
		return m
	default:
		return nil
	}
}

// It takes a map of strings to interfaces, and returns a map of strings to float32s
func ToMapFloat32(value interface{}) map[string]float32 {
	switch val := value.(type) {
	case map[string]interface{}:
		//convert the map to a map of strings
		m := make(map[string]float32)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Float32 {
				m[k] = ToFloat32(v)
			}
		}
		return m
	case map[string]int:
		//convert the map to a map of strings
		m := make(map[string]float32)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Float32 {
				m[k] = ToFloat32(v)
			}
		}
		return m
	case map[string]string:
		//convert the map to a map of strings
		m := make(map[string]float32)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Float32 {
				m[k] = ToFloat32(v)
			}
		}
		return m
	case map[string]float64:
		//convert the map to a map of strings
		m := make(map[string]float32)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Float32 {
				m[k] = ToFloat32(v)
			}
		}
		return m
	case map[string]bool:
		//convert the map to a map of strings
		m := make(map[string]float32)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Float32 {
				m[k] = ToFloat32(v)
			}
		}
		return m
	case map[string]float32:
		return val
	default:
		return nil
	}
}

// It takes a map of strings to interfaces, and returns a map of strings to floats
func ToMapFloat64(value interface{}) map[string]float64 {
	switch val := value.(type) {
	case map[string]interface{}:
		//convert the map to a map of strings
		m := make(map[string]float64)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Float64 {
				m[k] = ToFloat64(v)
			}
		}
		return m
	case map[string]int:
		//convert the map to a map of strings
		m := make(map[string]float64)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Float32 {
				m[k] = ToFloat64(v)
			}
		}
		return m
	case map[string]string:
		//convert the map to a map of strings
		m := make(map[string]float64)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Float32 {
				m[k] = ToFloat64(v)
			}
		}
		return m
	case map[string]float32:
		//convert the map to a map of strings
		m := make(map[string]float64)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Float32 {
				m[k] = ToFloat64(v)
			}
		}
		return m
	case map[string]bool:
		//convert the map to a map of strings
		m := make(map[string]float64)
		for k, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.Float32 {
				m[k] = ToFloat64(v)
			}
		}
		return m
	case map[string]float64:
		return val
	default:
		return nil
	}
}

// If the value is an array, return it as an array of interfaces
func ToArrayInterface(value interface{}) []interface{} {

	switch val := value.(type) {
	case []interface{}:
		return val
	case []string:
		//convert the array to an array of interfaces
		a := make([]interface{}, len(val))
		for i, v := range val {
			a[i] = v
		}
		return a
	case []int:
		//convert the array to an array of interfaces
		a := make([]interface{}, len(val))
		for i, v := range val {
			a[i] = v
		}
		return a
	case []float32:
		//convert the array to an array of interfaces
		a := make([]interface{}, len(val))
		for i, v := range val {
			a[i] = v
		}
		return a
	case []float64:
		//convert the array to an array of interfaces
		a := make([]interface{}, len(val))
		for i, v := range val {
			a[i] = v
		}
		return a
	case []bool:
		//convert the array to an array of interfaces
		a := make([]interface{}, len(val))
		for i, v := range val {
			a[i] = v
		}
		return a
	default:

		return nil
	}
}

// If the value is an array of strings, return the array, otherwise return nil
func ToArrayString(value interface{}) []string {
	switch val := value.(type) {
	case []interface{}:
		//convert the array to an array of strings
		a := make([]string, len(val))
		for i, v := range val {
			if reflect.TypeOf(v).Kind() == reflect.String {
				a[i] = ToString(v)
			}
		}
		return a
	case []string:
		return val
	case []int:
		//convert the array to an array of interfaces
		a := make([]string, len(val))
		for i, v := range val {
			a[i] = ToString(v)
		}
		return a
	case []float32:
		//convert the array to an array of interfaces
		a := make([]string, len(val))
		for i, v := range val {
			a[i] = ToString(v)
		}
		return a
	case []float64:
		//convert the array to an array of interfaces
		a := make([]string, len(val))
		for i, v := range val {
			a[i] = ToString(v)
		}
		return a
	case []bool:
		//convert the array to an array of interfaces
		a := make([]string, len(val))
		for i, v := range val {
			a[i] = ToString(v)
		}
		return a
	default:

		return nil
	}
}

// If the value is an array of ints, return it, otherwise return nil
func ToArrayInt(value interface{}) []int {
	switch val := value.(type) {
	case []interface{}:
		//convert the array to an array of strings
		a := make([]int, len(val))
		for i, v := range val {
			a[i] = ToInt(v)
		}
		return a
	case []string:
		//convert the array to an array of interfaces
		a := make([]int, len(val))
		for i, v := range val {
			a[i] = ToInt(v)
		}
		return a

	case []int:
		return val
	case []float32:
		//convert the array to an array of interfaces
		a := make([]int, len(val))
		for i, v := range val {
			a[i] = ToInt(v)
		}
		return a
	case []float64:
		//convert the array to an array of interfaces
		a := make([]int, len(val))
		for i, v := range val {
			a[i] = ToInt(v)
		}
		return a
	case []bool:
		//convert the array to an array of interfaces
		a := make([]int, len(val))
		for i, v := range val {
			a[i] = ToInt(v)
		}
		return a
	default:

		return nil
	}
}

// If the value is an array of float32, return the value as an array of float32
func ToArrayFloat32(value interface{}) []float32 {
	switch val := value.(type) {
	case []interface{}:
		//convert the array to an array of strings
		a := make([]float32, len(val))
		for i, v := range val {
			a[i] = ToFloat32(v)
		}
		return a
	case []string:
		//convert the array to an array of interfaces
		a := make([]float32, len(val))
		for i, v := range val {
			a[i] = ToFloat32(v)
		}
		return a

	case []int:
		//convert the array to an array of interfaces
		a := make([]float32, len(val))
		for i, v := range val {
			a[i] = ToFloat32(v)
		}
		return a

	case []float32:
		return val
	case []float64:
		//convert the array to an array of interfaces
		a := make([]float32, len(val))
		for i, v := range val {
			a[i] = ToFloat32(v)
		}
		return a
	case []bool:
		//convert the array to an array of interfaces
		a := make([]float32, len(val))
		for i, v := range val {
			a[i] = ToFloat32(v)
		}
		return a
	default:

		return nil
	}
}

// If the value is an array of float64, return it, otherwise return nil
func ToArrayFloat64(value interface{}) []float64 {
	switch val := value.(type) {
	case []interface{}:
		//convert the array to an array of strings
		a := make([]float64, len(val))
		for i, v := range val {

			a[i] = ToFloat64(v)

		}
		return a
	case []string:
		//convert the array to an array of interfaces
		a := make([]float64, len(val))
		for i, v := range val {
			a[i] = ToFloat64(v)
		}
		return a

	case []int:
		//convert the array to an array of interfaces
		a := make([]float64, len(val))
		for i, v := range val {
			a[i] = ToFloat64(v)
		}
		return a

	case []float32:
		//convert the array to an array of interfaces
		a := make([]float64, len(val))
		for i, v := range val {
			a[i] = ToFloat64(v)
		}
		return a

	case []float64:
		return val
	case []bool:
		//convert the array to an array of interfaces
		a := make([]float64, len(val))
		for i, v := range val {
			a[i] = ToFloat64(v)
		}
		return a
	default:

		return nil
	}
}

// If the value is an array of strings, return the array of strings, otherwise return nil
func ToArrayBool(value interface{}) []bool {
	switch val := value.(type) {
	case []interface{}:
		//convert the array to an array of strings
		a := make([]bool, len(val))
		for i, v := range val {

			a[i] = ToBool(v)

		}
		return a
	case []string:
		//convert the array to an array of interfaces
		a := make([]bool, len(val))
		for i, v := range val {
			a[i] = ToBool(v)
		}
		return a

	case []int:
		//convert the array to an array of interfaces
		a := make([]bool, len(val))
		for i, v := range val {
			a[i] = ToBool(v)
		}
		return a

	case []float32:
		//convert the array to an array of interfaces
		a := make([]bool, len(val))
		for i, v := range val {
			a[i] = ToBool(v)
		}
		return a

	case []float64:
		//convert the array to an array of interfaces
		a := make([]bool, len(val))
		for i, v := range val {
			a[i] = ToBool(v)
		}
		return a
	case []bool:
		return val
	default:

		return nil
	}
}

// It returns the type of the value passed to it as a string
func GetType(value interface{}) string {
	//Return the type as a string
	return reflect.TypeOf(value).Kind().String()
}
