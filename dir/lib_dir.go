// Package for working with directories
package dir

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"syscall"
)

// It takes a folder and zips it into a zip file
func ZipFolder(ZipFile string, zipFolder string) error {
	file, err := os.Create(ZipFile)
	if err != nil {
		return err
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
		return err
	}

	//defer file.Close()
	return nil
}

// > MakeDirAll creates a directory and all its parent directories if they don't exist
func MakeDirAll(filename string) error {
	file := path.Dir(filename)
	return os.MkdirAll(file, os.ModePerm)
}

// If the path exists, return true, otherwise return false
func DirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// MakeDir creates a directory if it does not exist and returns a function to remove it
/*
	Example:
	undoDir, err := MakeDirAllWithRemove(outputFilePath, 0755)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			undoDir()
		}
	}()
*/
func MakeDirAllWithRemove(dirPath string, perm os.FileMode) (func(), error) {
	var undoDir string

	for p := dirPath; ; p = path.Dir(p) {
		f_info, err := os.Stat(p)

		if err == nil {
			if f_info.IsDir() {
				break
			}

			f_info, err = os.Lstat(p)
			if err != nil {
				return nil, err
			}

			if f_info.IsDir() {
				break
			}

			return nil, &os.PathError{Op: "mkdirAll", Path: p, Err: syscall.ENOTDIR}
		}

		if os.IsNotExist(err) {
			undoDir = p
		} else {
			return nil, err
		}
	}

	if undoDir == "" {
		return func() {}, nil
	}

	if err := os.MkdirAll(dirPath, perm); err != nil {
		return nil, err
	}

	return func() { os.RemoveAll(undoDir) }, nil
}

// It opens the directory, reads the names of all the files in the directory, and then deletes them
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
