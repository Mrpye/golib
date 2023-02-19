package lib

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"syscall"
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
func MakeDirAll(filename string) error {
	file := path.Dir(filename)
	return os.MkdirAll(file, os.ModePerm)
}

//Check if Dir Exists
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
