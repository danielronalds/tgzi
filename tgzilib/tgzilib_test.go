package tgzilib

import (
	"errors"
	"os"
	"reflect"
	"strings"
	"testing"
)

// Tests if GetFiles works with the expected conditions
func TestGetFiles(t *testing.T) {
	inputDir := "."
	want := []string{"tgzilib.go", "tgzilib_test.go"}

	result, err := GetFiles(inputDir, true)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Got %v, wanted %v", result, want)
	}
}

// Tests if CompressFiles works with the expected conditions
func TestCompressFiles(t *testing.T) {
	archive := "test_archive.tar.gz"

	files, err := GetFiles(".", true)
	if err != nil {
		t.Fatalf("Got error getting files: %s", err.Error())
	}

	err = CompressFiles(archive, files)
	if err != nil {
		t.Fatalf("Got error: %s", err.Error())
	}

	exists, err := fileExists(archive)
	if !exists {
		t.Fatalf("Archive appears to not have been created!")
	}

	os.Remove(archive)
}

// Tests if CompressFiles can recognise if the file already exists
func TestCompressFilesReturningIfFileExists(t *testing.T) {
	archive := "test_archive.tar.gz"

	// Creating a dummy file
	f, err := os.Create(archive)
	if err != nil {
		t.Fatalf("Got error creating fake archive: %s", err.Error())
	}
	f.Close()                // no defer needed, want to close immediatly
	defer os.Remove(archive) // Want this file removed no matter what

	err = CompressFiles(archive, []string{"DummyFile.lol"})

	if err == nil {
		t.Fatalf("Expected an error, received nil!")
	}

	if !errors.Is(err, os.ErrExist) {
		t.Fatalf("Expected ErrExist, got %v", err)
	}
}

// Testing if NormaliseArchiveName appends the file extension
func TestNormaliseArchiveName(t *testing.T) {
	unformatedArchive := "archive"

	want := "archive.tar.gz"

	result := NormaliseArchiveName(unformatedArchive)

	if strings.Compare(want, result) != 0 {
		t.Fatalf("Wanted %s, got %s", want, result)
	}
}

// Testing how NormaliseArchiveName behaves if the file already has the file extension
func TestNormaliseArchiveNameWithNormalisedName(t *testing.T) {
	unformatedArchive := "archive.tar.gz"

	want := "archive.tar.gz"

	result := NormaliseArchiveName(unformatedArchive)

	if strings.Compare(want, result) != 0 {
		t.Fatalf("Wanted %s, got %s", want, result)
	}
}

// Testing how NormaliseArchiveName behaves with a blank file name
//
// NOTE: Expected behaviour is to supply a default name
func TestNormaliseArchiveNameWithBlankName(t *testing.T) {
	unformatedArchive := ""

	want := "archive.tar.gz"

	result := NormaliseArchiveName(unformatedArchive)

	if strings.Compare(want, result) != 0 {
		t.Fatalf("Wanted %s, got %s", want, result)
	}
}
