package indexed

import (
	"os"
	"path/filepath"
	"strings"
)

// Costruisce un indice: filename => contenuto
func buildFileContentIndex(files []string) map[string]string {
	index := make(map[string]string)
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}
		index[file] = string(content)
	}
	return index
}

func getUnusedFiles(files []string) []string {
	index := buildFileContentIndex(files)

	var unused []string
	for _, file := range files {
		if !isFileUsedIndexed(file, index) {
			unused = append(unused, file)
		}
	}
	return unused
}

func isFileUsedIndexed(file string, index map[string]string) bool {
	currentFile := filepath.Base(file)
	if strings.HasPrefix(currentFile, "+") {
		return true
	}
	for name, content := range index {
		if name == currentFile {
			continue
		}
		if strings.Contains(content, currentFile) {
			return true
		}
	}
	return false
}

func getUnusedFilesRecursive(files []string) []string {
	unused := getUnusedFiles(files)

	if len(unused) == 0 {
		return unused
	}

	// Escludiamo i file inutilizzati e ricominciamo
	var remaining []string
	unusedSet := make(map[string]bool)
	for _, f := range unused {
		unusedSet[f] = true
	}
	for _, f := range files {
		if !unusedSet[f] {
			remaining = append(remaining, f)
		}
	}

	// Ricorsione
	moreUnused := getUnusedFilesRecursive(remaining)
	return append(unused, moreUnused...)
}
