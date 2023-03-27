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

// ZipFolder zips a folder and all its contents
// - ZipFile: the name of the zip file to create
// - zipFolder: the folder to zip
// - returns: an error if there is one
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

//	MakeDirAll creates a directory and all its parent directories if they don't exist
//
// - filename: the name of the file to create
// - returns: an error if there is one
func MakeDirAll(filename string) error {
	file := path.Dir(filename)
	return os.MkdirAll(file, os.ModePerm)
}

// DirExists If the path exists, return true, otherwise return false
// - path: the path to check
// - returns: true if the path exists, otherwise false
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
// - dirPath: the path to the directory to create
// - perm: the permissions to set on the directory
// - returns: a function to remove the directory and an error if there is one
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

// RemoveContents removes all the contents of a directory
// - dir: the directory to remove the contents of
// - returns: an error if there is one
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
