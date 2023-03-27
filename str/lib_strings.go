// Package string provides a set of functions for working with strings
package str

import (
	"bytes"
	"compress/gzip"
	b64 "encoding/base64"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/Mrpye/golib/convert"
	uuid "github.com/nu7hatch/gouuid"
)

// const definitions for the input types
const (
	INPUT_TYPE_BOOL   = "bool"
	INPUT_TYPE_INT    = "int"
	INPUT_TYPE_FLOAT  = "float"
	INPUT_TYPE_STRING = "string"
)

// err definitions
var Err_ValueIsBlank = fmt.Errorf("a value is blank")

// CheckNotBlank takes a list of strings and returns an error if any of the strings are blank
// - value: a list of strings
// - returns: an error if any of the strings are blank
func CheckNotBlank(value ...string) error {
	for _, v := range value {
		if v == "" {
			return Err_ValueIsBlank
		}
	}
	return nil
}

// GzipBase64String takes a string or []byte, compresses it, and then base64 encodes it as a string
// - data: the data to compress and encode
// - returns: the compressed and encoded data as a string
//	and an error if there was a problem compressing or encoding the data
/*func GzipBase64String(data interface{}) (string, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	switch v := data.(type) {
	case []byte:
		if _, err := gz.Write(v); err != nil {
			return "", err
		}
	case string:
		if _, err := gz.Write([]byte(v)); err != nil {
			return "", err
		}
	}

	if err := gz.Close(); err != nil {
		return "", err
	}
	sEnc := b64.StdEncoding.EncodeToString(b.Bytes())
	return sEnc, nil
}*/

// GzipBase64 It takes a string, compresses it, and then base64 encodes it
//   - data: the data to compress and encode as a string or []byte
//   - returns: the compressed and encoded data as a string
//     and an error if there was a problem compressing or encoding the data
func GzipBase64(data interface{}) (string, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	switch v := data.(type) {
	case []byte:
		if _, err := gz.Write(v); err != nil {
			return "", err
		}
	case string:
		if _, err := gz.Write([]byte(v)); err != nil {
			return "", err
		}
	}
	if err := gz.Close(); err != nil {
		return "", err
	}
	sEnc := b64.StdEncoding.EncodeToString(b.Bytes())
	return sEnc, nil
}

// RemoveSpacesForHyphens It takes a string, converts it to lowercase, replaces all spaces with hyphens, and returns the
// string
// - value: the string to convert
// - returns: the converted string and an error if there was a problem converting the string
func RemoveSpacesForHyphens(value string) (string, error) {
	value = strings.ToLower(value)
	value = strings.ReplaceAll(value, " ", "-")
	return value, nil
}

// RemoveSpacesForUnderscores It takes a string, converts it to lowercase, and replaces all spaces with underscores
// - value: the string to convert
// - returns: the converted string
func RemoveSpacesForUnderscores(value string) string {
	value = strings.ToLower(value)
	value = strings.ReplaceAll(value, " ", "_")
	return value
}

// Concat takes a string and a list of strings and returns a string.
// - sep: the separator to use between the strings
// - string_list: a list of strings to concatenate
// - returns: a concatenated string
func Concat(sep string, string_list ...any) string {
	var str []string
	for _, s := range string_list {
		str = append(str, convert.ToString(s))
	}
	return strings.Join(str, sep)
}

// IsNumber returns true if a string is a number
// - s: the string to check
// - returns: true if the string is a number
func IsNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// GetDomainOrIP returns the domain or IP from a URL
// - url_str: the URL to parse
// - returns: the domain or IP from the URL
func GetDomainOrIP(url_str string) string {
	u, _ := url.Parse(url_str)
	host, _, err := net.SplitHostPort(u.Host)
	if err != nil {
		host = u.Host
	}
	return host
}

// GetPortString returns the port from a URL
// - url_str: the URL to parse
// - returns: the port from the URL
func GetPortString(url_str string) string {
	u, _ := url.Parse(url_str)
	_, port, _ := net.SplitHostPort(u.Host)
	return port
}

// GetPortInt returns the port from a URL as an int
// - url_str: the URL to parse
// - returns: the port from the URL as an int
func GetPortInt(url_str string) int {
	u, _ := url.Parse(url_str)
	_, port, _ := net.SplitHostPort(u.Host)
	intVar, _ := strconv.Atoi(port)
	return intVar
}

// Clean It takes a string and replaces all non-alphanumeric characters with the string passed in as the
// second argument
// - str: the string to clean
// - replace: the string to replace non-alphanumeric characters with
// - returns: the cleaned string
func Clean(str string, replace string) string {
	re, err := regexp.Compile(`[^\w]`)
	if err == nil {
		return re.ReplaceAllString(str, replace)
	}
	return str
}

// InsertString inserts a string into a slice of strings
// - orig: the original slice of strings
// - index: the index to insert the string at
// - value: the string to insert
// - returns: the new slice of strings
func InsertString(orig []string, index int, value string) []string {
	if index >= len(orig) {
		return append(orig, value)
	}
	orig = append(orig[:index+1], orig[index:]...)
	orig[index] = value

	return orig
}

// CastStringToType casts a string to a type
// - input_type: the type to cast to
//   - INPUT_TYPE_BOOL   = "bool"
//   - INPUT_TYPE_INT    = "int"
//   - INPUT_TYPE_FLOAT  = "float"
//
// - value: the string to cast
// - returns: the cast value
func CastStringToType(input_type string, value interface{}) interface{} {
	if input_type == INPUT_TYPE_BOOL {
		b1, _ := strconv.ParseBool(value.(string))
		return b1
	} else if input_type == INPUT_TYPE_INT {
		b1, _ := strconv.ParseInt(value.(string), 10, 64)
		return b1
	} else if input_type == INPUT_TYPE_FLOAT {
		b1, _ := strconv.ParseFloat(value.(string), 64)
		return b1
	}
	return value
}

// CommaListContainsString returns true if a string contains any of the items in a comma separated list
// - string_list: the comma separated list of strings
// - str: the string to check
// - returns: true if the string contains any of the items in the comma separated list
func CommaListContainsString(string_list string, str string) bool {
	arr := strings.Split(string_list, ",")
	for _, a := range arr {
		if strings.Contains(str, a) {
			return true
		}
	}
	return false
}

// CreateKey takes a string and returns a string that can be used as a key
// - value: the string to convert
// - replace_with: the string to replace spaces with
// - returns: the converted string
func CreateKey(value string, replace_with string) string {
	//Lets make lowercase and remove spaces
	value = strings.ToLower(value)
	value = strings.ReplaceAll(value, " ", replace_with)
	value = Clean(value, replace_with)
	return value
}

// CreateKeyOrGUID takes a string and returns a string that can be used as a key or a GUID if the string is empty
// - value: the string to convert
// - replace_with: the string to replace spaces with
// - returns: the converted string
// - error: any error that occurs
func CreateKeyOrGUID(value string, replace_with string) (string, error) {
	//**************************************
	//If no name is passed just generate one
	//**************************************
	if value == "" {
		u, err := uuid.NewV4()
		if err != nil {
			return value, err
		}
		value = u.String()
		return value, nil
	}
	//Lets make lowercase and remove spaces
	value = CreateKey(value, replace_with)
	return value, nil
}
