package main

import (
	"flag"
	"fmt"
	"svelte-unused-components/implementations/indexed"
	"svelte-unused-components/utils"
)

const Red = "\033[31m"
const Green = "\033[32m"
const Reset = "\033[0m"
const Yellow = "\033[33m"

var output = flag.String("o", "unused_files.txt", "Specifies the output file for the list of unused files. Defaults to `unused_files.txt`.")
var directory = flag.String("d", "./src", "Specifies the directory to search for `.svelte` files. Defaults to the `/src` directory")
var verbose = flag.Bool("v", false, "Enables verbose output.")
var ignored = flag.String("i", "", "Specifies the input file containing a list of files to ignore. Defaults to `ignore_files.txt`.")
var recursive = flag.Bool("r", false, "Enables recursive search")

func main() {
	flag.Parse()

	svelte_files := utils.GetSvelteFilesInDirecory(*directory)

	fmt.Println("\nFound"+Green, len(svelte_files), Reset+"svelte files in", *directory)
	fmt.Println("Analyzing files...")

	var unusedFiles []string
	unusedFiles = indexed.FindUnusedConcurrent(svelte_files, *recursive)

	fmt.Println("\n\n"+Red, len(unusedFiles), Reset+"unused files found.")
	utils.WriteResultsToFile(*output, unusedFiles)

	fmt.Println("\nFile list in " + Yellow + *output + "\n" + Reset)
}
