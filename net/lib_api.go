// This package contains helper functions for that are related to API calls.
package net

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Header struct for api headers
type Header struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

// It takes a username and password, concatenates them with a colon, encodes the resulting string in
// base64, and returns the encoded string
// - username: the username
// - password: the password
// - returns: the encoded string
func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// CallApi() is a function that takes a URL, a method, a list of headers, a payload, and a boolean to
// ignore SSL errors. It returns a byte array of result data, a boolean, and an error
// - url: the URL to call
// - method: the HTTP method to use POST GET PATH etc
// - headers: a list of headers to send with the request
// - payload: the data to send with the request
// - ignore_ssl: a boolean to ignore SSL errors
// - returns: a byte array of result data, a boolean is success, and an error if there is one
func CallApi(url string, method string, headers []Header, payload io.Reader, ignore_ssl bool) ([]byte, bool, error) {
	//****************************
	//Create the connection string
	//****************************
	req, err := http.NewRequest(method, url, payload) //New Request
	if err != nil {
		return nil, false, err
	}
	//***************
	//Set the headers
	//***************
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: ignore_ssl}
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()
	//*************
	//Read the data
	//*************
	responseData, err := ioutil.ReadAll(resp.Body)
	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	if err != nil {
		return responseData, false, err
	} else {
		if !statusOK {
			fmt.Println("Non-OK HTTP status:", resp.StatusCode)
			return responseData, false, err
		}
	}
	return responseData, true, nil
}

// CheckValidUrl Check If the host string contains either "http://" or "https://", return nil. Otherwise, return an error
// - host: the host string to check
// - returns: an error if the host string does not contain either "http://" or "https://"
func CheckValidUrl(host string) error {
	//Validate is a URL
	if strings.Contains(host, "http://") || strings.Contains(host, "https://") {
		return nil
	}
	return fmt.Errorf("Invalid URL: %s", host)
}

// RemoveURL removes http:// or https:// from a string
// - str: the string to remove http:// or https:// from
// - returns: the string without http:// or https://
func RemoveURL(str string) string {
	str = strings.ReplaceAll(str, "https://", "")
	str = strings.ReplaceAll(str, "http://", "")
	return str
}
