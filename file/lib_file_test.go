package file

import (
	"os"
	"testing"
)

// TestFileExists tests the FileExists function
func TestFileExists(t *testing.T) {

	err := DownloadFile("https://raw.githubusercontent.com/Mrpye/golib/main/README.md", "./test_file.txt")
	if err != nil {
		t.Errorf("DownloadFile() error = %v", err)
		return
	}
	if FileExists("./test_file.txt") != true {
		t.Errorf("FileExists() error = %v", err)
		return
	}

	l, err := CopyFile("./test_file.txt", "./test_file_copy.txt")
	if err != nil {
		t.Errorf("CopyFile() error = %v", err)
		return
	}
	if l <= 0 {
		t.Errorf("CopyFile() no data copied = %v", l)
		return
	}

	data, err := ReadFileToString("./test_file_copy.txt")
	if err != nil {
		t.Errorf("ReadFileToString() error = %v", err)
		return
	}

	if len(data) <= 0 {
		t.Errorf("ReadFileToString() no data = %v", data)
		return
	}

	err = SaveStringToFile("./test_file2.txt", data)
	if err != nil {
		t.Errorf("ReadFileToString() error = %v", err)
		return
	}
	//Clean up
	err = os.Remove("./test_file.txt")
	if err != nil {
		t.Errorf("os.Remove() error = %v", err)
		return
	}
	err = os.Remove("./test_file_copy.txt")
	if err != nil {
		t.Errorf("os.Remove() error = %v", err)
		return
	}
	err = os.Remove("./test_file2.txt")
	if err != nil {
		t.Errorf("os.Remove() error = %v", err)
		return
	}

	if FileExists("./test_file2.txt") {
		t.Errorf("FileExists() error = %v", err)
		return
	}

	if FileExists("./test_file2.txt") {
		t.Errorf("FileExists() error = %v", err)
		return
	}

	if FileExists("./test_file2.txt") {
		t.Errorf("FileExists() error = %v", err)
		return
	}

}
