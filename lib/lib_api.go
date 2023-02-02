package lib

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

//BasicAuth function for creating the basic auth header
func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// CallApi function for calling an api
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

// CheckValidUrl checks if a string is a valid URL
func CheckValidUrl(host string) error {
	if strings.Contains(strings.ToLower(host), "http://") || strings.Contains(strings.ToLower(host), "https://") {
		return nil
	} else {
		return fmt.Errorf("host must be a valid URL")
	}
}

// RemoveURL removes http:// or https:// from a string
func RemoveURL(str string) string {
	str = strings.ReplaceAll(str, "https://", "")
	str = strings.ReplaceAll(str, "http://", "")
	return str
}
