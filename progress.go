package main

import "fmt"

func updateProgress(current int, total int) {

	if !*recursive {
		printProgressBar(current, total)
	}
}

func printProgressBar(current int, total int) {
	progress := (current * PROGRESS_BAR_SIZE) / total
	fmt.Printf("\r[")
	for i := 0; i < progress; i++ {
		fmt.Printf("=")
	}
	for i := progress; i < PROGRESS_BAR_SIZE; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("] %d/%d", current+1, total)
}
