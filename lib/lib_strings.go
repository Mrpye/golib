package lib

import (
	"bytes"
	"compress/gzip"
	b64 "encoding/base64"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

//GzipBase64String gzip and base 64 a string
func GzipBase64String(data string) (string, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(data)); err != nil {
		return "", err
	}
	if err := gz.Close(); err != nil {
		return "", err
	}
	sEnc := b64.StdEncoding.EncodeToString(b.Bytes())
	return sEnc, nil
}

// RemoveSpacesForHyphens removes spaces and replaces them with hyphens
func RemoveSpacesForHyphens(name string) (string, error) {
	//Lets make lowercase and remove spaces
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")
	return name, nil
}

// RemoveSpacesForUnderscores removes spaces and replaces them with underscores
func RemoveSpacesForUnderscores(name string) string {
	//Lets make lowercase and remove spaces
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")
	return name
}

// Concat concatenates strings
func Concat(sep string, strs ...string) string {
	return strings.Join(strs, sep)
}

// IsNumber returns true if a string is a number
func IsNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// GetDomainOrIP returns the domain or IP from a URL
func GetDomainOrIP(url_str string) string {
	u, _ := url.Parse(url_str)
	host, _, _ := net.SplitHostPort(u.Host)
	return host
}

// getPortString returns the port from a URL
func GetPortString(url_str string) string {
	u, _ := url.Parse(url_str)
	_, port, _ := net.SplitHostPort(u.Host)
	return port
}

// getPortInt returns the port from a URL as an int
func GetPortInt(url_str string) int {
	u, _ := url.Parse(url_str)
	_, port, _ := net.SplitHostPort(u.Host)
	intVar, _ := strconv.Atoi(port)
	return intVar
}

// Clean cleans a string
func Clean(str string, replace string) string {
	re, err := regexp.Compile(`[^\w]`)
	if err == nil {
		return re.ReplaceAllString(str, replace)
	}
	return str
}

// InsertString inserts a string into a slice of strings
func InsertString(orig []string, index int, value string) []string {
	if index >= len(orig) {
		return append(orig, value)
	}
	orig = append(orig[:index+1], orig[index:]...)
	orig[index] = value

	return orig
}

// castStringToType casts a string to a type
func CastStringToType(input_type string, value interface{}) interface{} {
	if input_type == "bool" {
		b1, _ := strconv.ParseBool(value.(string))
		return b1
	} else if input_type == "int" {
		b1, _ := strconv.ParseInt(value.(string), 10, 64)
		return b1
	} else if input_type == "float" {
		b1, _ := strconv.ParseFloat(value.(string), 64)
		return b1
	}
	return value
}

// CommaListContainsString returns true if a string contains any of the items in a comma separated list
func CommaListContainsString(string_list string, str string) bool {
	arr := strings.Split(string_list, ",")
	for _, a := range arr {
		if strings.Contains(str, a) {
			return true
		}
	}
	return false
}
