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

var output = flag.String("o", "unused_files.txt", "Specifies the output file for the list of unused files. Defaults to `unused_files.txt`.")
var directory = flag.String("d", "./src", "Specifies the directory to search for `.svelte` files. Defaults to the `/src` directory")
var verbose = flag.Bool("v", false, "Enables verbose output. (This will also disable progress display.)")
var ignored = flag.String("i", "", "Specifies the input file containing a list of files to ignore. Defaults to `ignore_files.txt`.")
var recursive = flag.Bool("r", false, "Enables recursive search")
var noProgressBar = flag.Bool("np", false, "Disables the progress display (useful when redirecting output)")

func main() {
	flag.Parse()

	svelte_files := getSvelteFiles(*directory)

	fmt.Println("\nFound"+Green, len(svelte_files), Reset+"svelte files in", *directory)
	fmt.Println("Analyzing files...")
	if !*noProgressBar && !*verbose && *recursive {
		initProgressTracker(len(svelte_files))
		fmt.Println(Green + "O" + Reset + " = used file")
		fmt.Println(Red + "X" + Reset + " = unused file")
	}

	var unusedFiles []string
	if *recursive {
		unusedFiles = getUnusedFilesRecursive(svelte_files)
	} else {
		fmt.Println("\033[?25l")
		unusedFiles = getUnusedFiles(svelte_files)
	}

	fmt.Println("\n\n"+Red, len(unusedFiles), Reset+"unused files found.")
	writeToFile(*output, unusedFiles)

	fmt.Println("\nFile list in " + Yellow + *output + "\n" + Reset)
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
			if *verbose {
				fmt.Println(Red + file + Reset)
			}
			unusedFiles = append(unusedFiles, file)
		} else if *verbose {
			fmt.Println(Green + file + Reset)
		}
		if !*noProgressBar && !*verbose {
			updateProgress(i, len(files), isUsed)
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
			fmt.Println("Error while reading file:", scanner.Err())
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
