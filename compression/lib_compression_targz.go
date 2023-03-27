// Package for working with tar.gz files
package compression

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Mrpye/golib/dir"
	"github.com/Mrpye/golib/path"
)

// Compress creates a archive from the folder inputFilePath points to in the file outputFilePath points to.
// Only adds the last directory in inputFilePath to the archive, not the whole path.
// It tries to create the directory structure outputFilePath contains if it doesn't exist.
// It returns potential errors to be checked or nil if everything works.
// - inputFilePath: the path to the folder to compress
// - outputFilePath: the path to the archive to create
// - returns: an error if there is one
func Compress(inputFilePath, outputFilePath string) (err error) {
	inputFilePath = path.StripTrailingSlashes(inputFilePath)
	inputFilePath, outputFilePath, err = makeAbsolute(inputFilePath, outputFilePath)
	if err != nil {
		return err
	}
	undoDir, err := dir.MakeDirAllWithRemove(filepath.Dir(outputFilePath), 0755)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			undoDir()
		}
	}()

	//err = compress(inputFilePath, outputFilePath, filepath.Dir(inputFilePath))
	err = compress(inputFilePath, outputFilePath, inputFilePath)
	if err != nil {
		return err
	}

	return nil
}

// Extract extracts a archive from the file inputFilePath points to in the directory outputFilePath points to.
// It tries to create the directory structure outputFilePath contains if it doesn't exist.
// It returns potential errors to be checked or nil if everything works.
// - inputFilePath: the path to the archive to extract
// - outputFilePath: the path to the folder to create
// - returns: an error if there is one
func Extract(inputFilePath, outputFilePath string) (err error) {
	outputFilePath = path.StripTrailingSlashes(outputFilePath)
	inputFilePath, outputFilePath, err = makeAbsolute(inputFilePath, outputFilePath)
	if err != nil {
		return err
	}
	undoDir, err := dir.MakeDirAllWithRemove(outputFilePath, 0755)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			undoDir()
		}
	}()

	return extract(inputFilePath, outputFilePath)
}

// Make input and output paths absolute.
// Returns the absolute paths and an error if there is one.
// - inputFilePath: the path to the folder to compress
// - outputFilePath: the path to the archive to create
// - returns: the absolute paths and an error if there is one
func makeAbsolute(inputFilePath, outputFilePath string) (string, string, error) {
	inputFilePath, err := filepath.Abs(inputFilePath)
	if err == nil {
		outputFilePath, err = filepath.Abs(outputFilePath)
	}

	return inputFilePath, outputFilePath, err
}

// compress The main interaction with tar and gzip. Creates a archive and recursivly adds all files in the directory.
// The finished archive contains just the directory added, not any parents.
// This is possible by giving the whole path exept the final directory in subPath.
// - inPath: the path to the folder to compress
// - outFilePath: the path to the archive to create
// - subPath: the path to the folder to compress without the last directory
// - returns: an error if there is one
func compress(inPath, outFilePath, subPath string) (err error) {
	files, err := ioutil.ReadDir(inPath)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errors.New("targz: input directory is empty")
	}

	file, err := os.Create(outFilePath)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			os.Remove(outFilePath)
		}
	}()

	gzipWriter := gzip.NewWriter(file)
	tarWriter := tar.NewWriter(gzipWriter)

	err = writeDirectory(inPath, tarWriter, subPath)
	if err != nil {
		return err
	}

	err = tarWriter.Close()
	if err != nil {
		return err
	}

	err = gzipWriter.Close()
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

// writeDirectory Read a directory and write it to the tar writer. Recursive function that writes all sub folders.
// - directory: the path to the folder to compress
// - tarWriter: the tar writer to write to
// - subPath: the path to the folder to compress without the last directory
// - returns: an error if there is one
func writeDirectory(directory string, tarWriter *tar.Writer, subPath string) error {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, file := range files {
		currentPath := filepath.Join(directory, file.Name())
		if file.IsDir() {
			err := writeDirectory(currentPath, tarWriter, subPath)
			if err != nil {
				return err
			}
		} else {
			err = writeTarGz(currentPath, tarWriter, file, subPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// writeTarGz Write path without the prefix in subPath to tar writer.
// - path: the path to the file to compress
// - tarWriter: the tar writer to write to
// - fileInfo: the file info of the file to compress
// - subPath: the path to the folder to compress without the last directory
// - returns: an error if there is one
func writeTarGz(path string, tarWriter *tar.Writer, fileInfo os.FileInfo, subPath string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	evaledPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		return err
	}

	subPath, err = filepath.EvalSymlinks(subPath)
	if err != nil {
		return err
	}

	link := ""
	if evaledPath != path {
		link = evaledPath
	}

	header, err := tar.FileInfoHeader(fileInfo, link)
	if err != nil {
		return err
	}
	file_name := strings.ReplaceAll(evaledPath[len(subPath):], "\\", "/")
	header.Name = strings.Replace(file_name, "/", "", 1)

	err = tarWriter.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tarWriter, file)
	if err != nil {
		return err
	}

	return err
}

// extract the file in filePath to directory.
// - filePath: the path to the archive to extract
// - directory: the path to the folder to extract to
// - returns: an error if there is one
func extract(filePath string, directory string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(bufio.NewReader(file))
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		fileInfo := header.FileInfo()
		dir := filepath.Join(directory, filepath.Dir(header.Name))
		filename := filepath.Join(dir, fileInfo.Name())

		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}

		file, err := os.Create(filename)
		if err != nil {
			return err
		}

		writer := bufio.NewWriter(file)

		buffer := make([]byte, 4096)
		for {
			n, err := tarReader.Read(buffer)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if n == 0 {
				break
			}

			_, err = writer.Write(buffer[:n])
			if err != nil {
				return err
			}
		}

		err = writer.Flush()
		if err != nil {
			return err
		}

		err = file.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

// ExtractFile It opens a gzip file, reads it as a tar file, and returns the contents of the file with the name you
// specify
// - filePath: the path to the archive to extract
// - file_name: the name of the file to extract
// - returns: the contents of the file with the name you specify
// - returns: an error if there is one
func ExtractFile(filePath string, file_name string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(bufio.NewReader(file))
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		fileInfo := header.FileInfo()

		if fileInfo.Name() == file_name {
			buf := bytes.NewBuffer(nil)
			writer := bufio.NewWriter(buf)

			buffer := make([]byte, 4096)
			for {
				n, err := tarReader.Read(buffer)
				if err != nil && err != io.EOF {
					panic(err)
				}
				if n == 0 {
					break
				}

				_, err = writer.Write(buffer[:n])
				if err != nil {
					return nil, err
				}
			}

			err = writer.Flush()
			if err != nil {
				return nil, err
			}

			err = file.Close()
			if err != nil {
				return nil, err
			}

			return buf.Bytes(), nil
		}

		err = file.Close()
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
