package bruteforce

import (
	"runtime"
	"sync"
)

func getUnusedFilesRecursiveConcurrent(files []string) []string {
	unused := getUnusedFilesConcurrent(files)

	if len(unused) == 0 {
		return unused
	}

	// Rimuoviamo gli inutilizzati per la chiamata successiva
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

	// Ricorsione sui rimanenti
	moreUnused := getUnusedFilesRecursiveConcurrent(remaining)
	return append(unused, moreUnused...)
}

func getUnusedFilesConcurrent(files []string) []string {
	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()

	fileChan := make(chan string)
	resultChan := make(chan string)

	// Start workers
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range fileChan {
				if !isFileUsed(file, files) {
					resultChan <- file
				}
			}
		}()
	}

	// Feed jobs
	go func() {
		for _, file := range files {
			fileChan <- file
		}
		close(fileChan)
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var unused []string
	for file := range resultChan {
		unused = append(unused, file)
	}

	return unused
}
