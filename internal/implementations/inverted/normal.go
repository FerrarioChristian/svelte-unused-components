package inverted

import (
	"github.com/FerrarioChristian/svelte-unused-components/internal/utils"
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
	usedMap := make(map[string]bool)

	for _, file := range files {
		if strings.HasPrefix(filepath.Base(file), "+") {
			usedMap[file] = true
		}
		lines := utils.ReadFileLines(file)

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
