package lib

import (
	"strconv"
	"strings"
)

func PassFail(value bool) string {
	if value {
		return "Pass"
	}
	return "Fail"
}

// Convert to bool
func ConvertToBool(value interface{}) bool {
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
		}
		b, _ := strconv.ParseBool(val)
		return b
	case int:
		return val > 0
	default:
		return false
	}
}

func ConvertToString(value interface{}) string {

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
