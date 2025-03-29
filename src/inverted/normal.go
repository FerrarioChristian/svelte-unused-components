package inverted

import (
	"bufio"
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
		unusedFiles = append(unusedFiles, getUnusedFiles(updatedFiles)...)
	}

	return unusedFiles
}

func getUnusedFiles(files []string) []string {
	usedMap := make(map[string]bool)

	for _, file := range files {
		if strings.HasPrefix(filepath.Base(file), "+") {
			usedMap[file] = true
		}
		lines := readFileLines(file)

		for _, line := range lines {
			for _, candidate := range files {
				name := filepath.Base(candidate)
				if strings.Contains(line, name) {
					usedMap[candidate] = true
				}
			}
		}
	}

	var unusedFiles []string
	for _, file := range files {
		if !usedMap[file] {
			unusedFiles = append(unusedFiles, file)
		}
	}
	return unusedFiles
}

func readFileLines(file string) []string {
	var lines []string
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "</script>") {
			break
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
