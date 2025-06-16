package inverted

import (
	"github.com/FerrarioChristian/svelte-unused-components/internal/utils"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"sync"
)

func getUnusedFilesRecursiveConcurrent(files []string) []string {
	unusedFiles := getUnusedFilesConcurrent(files)

	if len(unusedFiles) > 0 {
		var updatedFiles []string
		for _, file := range files {
			if !slices.Contains(unusedFiles, file) {
				updatedFiles = append(updatedFiles, file)
			}
		}
		unusedFiles = append(unusedFiles, getUnusedFilesRecursiveConcurrent(updatedFiles)...)
	}

	return unusedFiles
}

func getUnusedFilesConcurrent(files []string) []string {
	var usedMap sync.Map
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			if strings.HasPrefix(filepath.Base(file), "+") {
				usedMap.Store(file, true)
			}

			lines := utils.ReadFileLines(file)

			for _, line := range lines {
				for _, candidate := range files {
					name := filepath.Base(candidate)
					if strings.Contains(line, name) {
						usedMap.Store(candidate, true)
					}
				}
			}
		}(file)
	}

	wg.Wait()

	var unusedFiles []string
	for _, file := range files {
		if _, used := usedMap.Load(file); !used {
			unusedFiles = append(unusedFiles, file)
		}
	}
	return unusedFiles
}

func getUnusedFilesRecursiveWorkers(files []string) []string {
	unusedFiles := getUnusedFilesWorkers(files)

	if len(unusedFiles) > 0 {
		var updatedFiles []string
		for _, file := range files {
			if !slices.Contains(unusedFiles, file) {
				updatedFiles = append(updatedFiles, file)
			}
		}
		unusedFiles = append(unusedFiles, getUnusedFilesRecursiveWorkers(updatedFiles)...)
	}

	return unusedFiles
}

func getUnusedFilesWorkers(files []string) []string {
	var usedMap sync.Map
	var wg sync.WaitGroup

	fileChan := make(chan string)

	// Numero di worker goroutine attive
	numWorkers := runtime.NumCPU()
	// Avvia i worker
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range fileChan {
				if strings.HasPrefix(filepath.Base(file), "+") {
					usedMap.Store(file, true)
				}

				lines := utils.ReadFileLines(file)

				for _, line := range lines {
					for _, candidate := range files {
						name := filepath.Base(candidate)
						if strings.Contains(line, name) {
							usedMap.Store(candidate, true)
						}
					}
				}
			}
		}()
	}

	// Invia tutti i file al canale
	go func() {
		for _, file := range files {
			fileChan <- file
		}
		close(fileChan)
	}()

	wg.Wait()

	// Raccoglie i file non usati
	var unusedFiles []string
	for _, file := range files {
		if _, used := usedMap.Load(file); !used {
			unusedFiles = append(unusedFiles, file)
		}
	}
	return unusedFiles
}
