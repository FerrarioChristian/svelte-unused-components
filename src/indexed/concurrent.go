package indexed

import (
	"runtime"
	"sync"
)

func getUnusedFilesConcurrent(files []string) []string {
	index := buildFileContentIndex(files)
	tasks := make(chan string, len(files))
	results := make(chan string, len(files))

	workers := runtime.NumCPU()

	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range tasks {
				if !isFileUsedIndexed(file, index) {
					results <- file
				}
			}
		}()
	}

	for _, file := range files {
		tasks <- file
	}
	close(tasks)

	wg.Wait()
	close(results)

	var unused []string
	for f := range results {
		unused = append(unused, f)
	}
	return unused
}

func getUnusedFilesRecursiveConcurrent(files []string) []string {
	workers := runtime.NumCPU()

	index := buildFileContentIndex(files)
	unused := getUnusedFilesConcurrentWithIndex(files, index, workers)

	if len(unused) == 0 {
		return unused
	}

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

	// Ricorsione sui rimanenti, ricostruendo l'indice
	moreUnused := getUnusedFilesRecursiveConcurrent(remaining)
	return append(unused, moreUnused...)
}

// Variante interna riutilizzabile con indice pre-costruito
func getUnusedFilesConcurrentWithIndex(files []string, index map[string]string, workers int) []string {
	tasks := make(chan string, len(files))
	results := make(chan string, len(files))

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range tasks {
				if !isFileUsedIndexed(file, index) {
					results <- file
				}
			}
		}()
	}

	for _, file := range files {
		tasks <- file
	}
	close(tasks)

	wg.Wait()
	close(results)

	var unused []string
	for f := range results {
		unused = append(unused, f)
	}
	return unused
}
