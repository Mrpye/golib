package dir

import (
	"os"
	"testing"

	"github.com/Mrpye/golib/file"
)

// TestFileExists tests the FileExists function
func TestDir(t *testing.T) {

	if DirExists("./test/test/") {
		t.Error("DirExists returned true for a non-existing directory")
		err := os.RemoveAll("./test/")
		if err != nil {
			t.Error("ZipFolder returned an error")
			return
		}
		return
	}

	err := MakeDirAll("./test/test/")
	if err != nil {
		t.Error(err)
		return
	}

	if !DirExists("./test/test/") {
		t.Error("DirExists returned false for an existing directory")
		return
	}

	err = file.SaveStringToFile("./test/test/test.txt", "test")
	if err != nil {
		t.Error("SaveStringToFile returned an error")
		return
	}
	err = file.SaveStringToFile("./test/test/test2.txt", "test")
	if err != nil {
		t.Error("SaveStringToFile returned an error")
		return
	}

	err = RemoveContents("./test/test/")
	if err != nil {
		t.Error("RemoveContents returned an error")
		return
	}

	if file.FileExists("./test/test/test2.txt") {
		t.Error("RemoveContents did not remove the files")
		return
	}

	err = file.SaveStringToFile("./test/test/test.txt", "test")
	if err != nil {
		t.Error("SaveStringToFile returned an error")
		return
	}
	err = file.SaveStringToFile("./test/test/test2.txt", "test")
	if err != nil {
		t.Error("SaveStringToFile returned an error")
		return
	}

	err = ZipFolder("./test.zip", "test/test/")
	if err != nil {
		t.Error("ZipFolder returned an error")
		return
	}

	info, err := os.Stat("./test.zip")
	if err != nil {
		t.Error(" os.Stat returned an error")
		return
	}
	x := info.Size()
	if x == 0 {
		t.Error("ZipFolder did not create a zip file")
		return
	}

	if !file.FileExists("./test.zip") {
		t.Error("ZipFolder did not create the zip file")
		return
	}

	err = os.RemoveAll("./test/")
	if err != nil {
		t.Error("ZipFolder returned an error")
		return
	}

	err = os.Remove("./test.zip")
	if err != nil {
		t.Error("ZipFolder returned an error")
		return
	}

}
