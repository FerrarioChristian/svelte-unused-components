package main

import "fmt"

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

func updateProgress(current int, total int, used bool) {

	if !*recursive {
		printProgressBar(current, total)
	}
	printProgressMatrix(current, total, used)
}

func initProgressTracker(size int) {
	progressTracker = make([]bool, size)
}

func printProgressMatrix(current int, total int, used bool) {
	if used {
		progressTracker[current] = true
		fmt.Printf("\r[")
		for i := 0; i < current-calcProgressOffset(current); i++ {
			fmt.Printf(Green + "O" + Reset)
		}
		for i := current - calcProgressOffset(current); i < total; i++ {
			if progressTracker[i] {
				fmt.Printf(Green + "O" + Reset)
			} else {
				fmt.Printf(Red + "X" + Reset)
			}
		}
		fmt.Printf("] %d/%d", current+1, total)
	} else {
		fmt.Printf("\r[")
		for i := 0; i < current-calcProgressOffset(current); i++ {
			fmt.Printf(Green + "O" + Reset)
		}
		for i := current - calcProgressOffset(current); i < total; i++ {
			if progressTracker[i] {
				fmt.Printf(Green + "O" + Reset)
			} else {
				fmt.Printf(Red + "X" + Reset)
			}
		}
		fmt.Printf("] %d/%d", current+1, total)
	}
}

func calcProgressOffset(current int) int {
	offset := 0
	for i := 0; i < current; i++ {
		if !progressTracker[i] {
			offset++
		}
	}
	return offset
}
