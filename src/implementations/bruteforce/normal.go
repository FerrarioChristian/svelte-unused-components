package bruteforce

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func getUnusedFilesRecursive(files []string) []string {
	unusedFiles := getUnusedFiles(files)

	if len(unusedFiles) > 0 {
		var updatedFiles []string
		for _, file := range files {
			if !slices.Contains(unusedFiles, file) {
				updatedFiles = append(updatedFiles, file)
			}
		}
		unusedFiles = append(unusedFiles, getUnusedFilesRecursive(updatedFiles)...)
	}

	return unusedFiles
}

func getUnusedFiles(files []string) []string {
	var unusedFiles []string

	for _, file := range files {
		isUsed := isFileUsed(file, files)
		if !isUsed {
			unusedFiles = append(unusedFiles, file)
		}
	}

	return unusedFiles
}

func isFileUsed(file string, files []string) bool {
	currentFile := filepath.Base(file)
	if strings.HasPrefix(currentFile, "+") {
		return true
	}
	for _, f := range files {

		f, err := os.Open(f)
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, currentFile) {
				f.Close()
				return true
			}
			if strings.Contains(line, "</script>") {
				break
			}
		}
		f.Close()

		if scanner.Err() != nil {
			fmt.Println("Error while reading file:", scanner.Err())
		}
	}
	return false
}
