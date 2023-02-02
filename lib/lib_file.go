package lib

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// ZipFolder zips a folder
func ZipFolder(ZipFile string, zipFolder string) {
	file, err := os.Create(ZipFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		fmt.Printf("Crawling: %#v\n", path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.
		/*log.Print(path)
		parts := strings.Split(path, string(os.PathSeparator))
		cleaned := parts[1:]
		path = strings.Join(cleaned, string(os.PathSeparator))*/
		f, err := w.Create(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	}
	err = filepath.Walk(zipFolder, walker)
	if err != nil {
		panic(err)
	}

	defer file.Close()

}

// MakeDir creates a directory if it does not exist
func MakeDir(filename string) error {
	file := path.Dir(filename)
	return os.MkdirAll(file, os.ModePerm)
}

// Download a file from a url
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file := path.Dir(filepath)
	err = os.MkdirAll(file, os.ModePerm)
	if err != nil {
		return err
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// CopyFile copies a file from src to dst.
func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// SaveStringToFile save string to file
func SaveStringToFile(filename string, content string) error {
	if content != "" {
		file := path.Dir(filename)
		err := os.MkdirAll(file, os.ModePerm)

		if err != nil {
			return err
		}
		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		if err != nil {
			return err
		}
		f.WriteString(string(content))
	}
	return nil
}

// ReadFileToString read file to string
func ReadFileToString(sourceFile string) (string, error) {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return "", err
	}
	return string(input), nil
}

// Check if File Exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

//Check if Dir Exists
func DirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
