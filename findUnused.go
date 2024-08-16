package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const Red = "\033[31m"
const Green = "\033[32m"
const Reset = "\033[0m"
const Yellow = "\033[33m"

const PROGRESS_BAR_SIZE = 100

var progressTracker []bool
var isFirstRecursiveCall = true

var recursive = flag.Bool("r", false, "Enables recursive search")
var noProgressBar = flag.Bool("np", false, "Disables the progress display (useful when redirecting output)")

func main() {
	flag.Parse()

	svelte_files := getSvelteFiles("./test/fe")

	fmt.Println("\nAnalisi di"+Green, len(svelte_files), Reset+"file svelte in corso...")

	fmt.Println(Green + "O" + Reset + " = file utilizzato")
	fmt.Println(Red + "X" + Reset + " = file non utilizzato")

	var unusedFiles []string
	if *recursive {
		unusedFiles = getUnusedFilesRecursive(svelte_files)
	} else {
		fmt.Println("\033[?25l")
		unusedFiles = getUnusedFiles(svelte_files)
	}

	fmt.Println("\n\n\n\nSono stati trovati"+Red, len(unusedFiles), Reset+"file non utilizzati.")
	writeToFile("unused_files.txt", unusedFiles)

	fmt.Println("\nLista dei file inutilizzati in " + Yellow + "unused_files.txt\n\n" + Reset)
}

func getSvelteFiles(root string) []string {
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

func getUnusedFilesRecursive(files []string) []string {
	unusedFiles := getUnusedFiles(files)

	if len(unusedFiles) > 0 {
		var updatedFiles []string
		for _, file := range files {
			if !contains(unusedFiles, file) {
				updatedFiles = append(updatedFiles, file)
			}
		}
		isFirstRecursiveCall = false
		unusedFiles = append(unusedFiles, getUnusedFilesRecursive(updatedFiles)...)
	}

	return unusedFiles
}

func getUnusedFiles(files []string) []string {
	var unusedFiles []string

	for i, file := range files {
		isUsed := isFileUsed(file, files)
		if !isUsed {
			unusedFiles = append(unusedFiles, file)
		}
		if !*noProgressBar {
			updateProgress(i, len(files))
		}
	}

	return unusedFiles
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
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
			fmt.Println("Errore durante la lettura del file:", scanner.Err())
		}
	}
	return false
}

func writeToFile(filename string, lines []string) {

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
