package tgzilib

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

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
