package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var Red = "\033[31m"
var Green = "\033[32m"
var Reset = "\033[0m"
var Yellow = "\033[33m"

func main() {

	svelte_files, err := getSvelteFiles("./test/app")
	if err != nil {
		panic(err)
	}

	fmt.Println("\nAnalisi di"+Green, len(svelte_files), Reset+"file svelte in corso...")
	fmt.Println(Green + "O" + Reset + " = file utilizzato")
	fmt.Println(Red + "X" + Reset + " = file non utilizzato\n")

	unusedFiles := getUnusedFiles(svelte_files)

	fmt.Println("\n\nSono stati trovati"+Red, len(unusedFiles), Reset+"file non utilizzati. \n")

	writeErr := writeToFile("unused_files.txt", unusedFiles)
	if writeErr != nil {
		fmt.Println("Errore:", err)
	}

	fmt.Println("Lista dei file inutilizzati in " + Yellow + "unused_files.txt\n\n" + Reset)

}

func getSvelteFiles(root string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
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
		return nil, err
	}

	return files, nil
}

func getUnusedFiles(files []string) []string {
	var unusedFiles []string

	for i, file := range files {
		if !isFileUsed(file, files) {
			unusedFiles = append(unusedFiles, file)
		}
		if (i+1)%100 == 0 {
			fmt.Print("\n")
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
				fmt.Print(Green + "O" + Reset)
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
	fmt.Print(Red + "X" + Reset)
	return false
}

func writeToFile(filename string, lines []string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
