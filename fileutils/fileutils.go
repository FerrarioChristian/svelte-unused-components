package fileutils

import (
	"io/fs"
	"os"
	"path/filepath"
)

func GetSvelteFilesInDirecory(root string) []string {
	var files []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".svelte" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return files
}

func WriteResultsToFile(filename string, lines []string) {

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}
}
