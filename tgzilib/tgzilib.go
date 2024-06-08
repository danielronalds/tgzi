package tgzilib

import (
	"fmt"
	"os"
)

// Gets the list of files and directories in a given directory
//
// # Params
//
// - `dir` The directory to look in. i.e. "." looks in the current directory
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
