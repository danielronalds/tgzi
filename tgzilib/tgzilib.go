package tgzilib

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Gets the list of files and directories in a given directory
//
// # Params
//
// - `dir` The directory to look in. i.e. "." looks in the current directory
//
// - `hideDotfiles` Whether to return dotfiles or not
//
// # Returns
//
// Returns an error if the directory is unreadable, otherwise a slice of filenames as strings
func GetFiles(dir string, hideDotfiles bool) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	files := make([]string, 0)

	for _, entry := range entries {
		filename := entry.Name()

		if filename[0] == '.' && hideDotfiles {
			continue
		}

		if entry.IsDir() {
			filename = fmt.Sprintf("%v/", filename)
		}
		files = append(files, filename)
	}

	return files, nil
}

// Compress the slice of files into a .tar.gz archive
//
// # Params
//
// - `archive` The name of the archive to compress into
//
// - `files` The slice of files to compress
//
// # Returns
//
// An error if the archive already exists, or if  one occurs during either running the command
func CompressFiles(archive string, files []string) error {
	if len(files) == 0 {
		return nil
	}

	// Checking if the file already exits
	exists, err := fileExists(archive)
	if err != nil {
		return err
	}
	if exists {
		return os.ErrExist
	}

	args := []string{"-czvf", archive}
	for _, file := range files {
		args = append(args, file)
	}

	cmd := exec.Command("tar", args...)

	return cmd.Run()
}

// Returns whether a file exists or not
//
// # Params
//
// - `file` The name of the file to check
func fileExists(file string) (bool, error) {
	_, err := os.Stat(file)

	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, err
}

// Ensures the archive has the file extensions `.tar.gz`
//
// # Params
//
// - `archive` The name of the archive
//
// # Returns
//
// The archive name with the appropriate file extension
func NormaliseArchiveName(archive string) string {
	file := strings.Split(archive, ".")
	return fmt.Sprintf("%s.tar.gz", file[0])
}
