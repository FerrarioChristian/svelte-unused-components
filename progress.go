package main

import "fmt"

func initProgressTracker(size int) {
	progressTracker = make([]bool, size)
	for i := 0; i < size; i++ {
		progressTracker[i] = true
	}
}

func updateProgress(current int, total int, isUsed bool) {

	if !*recursive {
		printProgressBar(current, total)
	} else {
		printProgressMatrix(current, isUsed, total)
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

func printProgressMatrix(current int, isUsed bool, total int) {
	if isFirstRecursiveCall {
		if current%PROGRESS_BAR_SIZE == 0 {
			fmt.Println()
		}
		if isUsed {
			fmt.Printf(Green + "O" + Reset)
			progressTracker[current] = true
		} else {
			fmt.Printf(Red + "X" + Reset)
			progressTracker[current] = false
		}
	} else {
		if current == 0 {
			rowsPrinted := (total / PROGRESS_BAR_SIZE)
			fmt.Printf("\033[%dA\r", rowsPrinted+1)
		}

		realIndex := findNthOPosition(current)

		if realIndex == -1 {
			realIndex = total - 1
		}

		if realIndex%PROGRESS_BAR_SIZE == 0 {
			fmt.Println()
		}

		if isUsed && progressTracker[realIndex] {
			fmt.Printf("\033[C")
		}

		if !isUsed && progressTracker[realIndex] {
			fmt.Printf(Red + "X" + Reset)
			progressTracker[realIndex] = false
		}
	}
}

func findNthOPosition(index int) int {
	occurrencesFound := 0
	for i, value := range progressTracker {
		if value {
			if occurrencesFound == index {
				checkMissedNewline(i)
				return i
			}
			occurrencesFound++
		}
	}
	return -1
}

func checkMissedNewline(currentIndex int) {
	// Trova l'ultimo valore `true` prima del currentIndex
	lastTrueIndex := -1
	for i := currentIndex - 1; i >= 0; i-- {
		if progressTracker[i] {
			lastTrueIndex = i
			break
		}
	}

	if lastTrueIndex == -1 {
		return
	}

	for i := lastTrueIndex + 1; i < currentIndex; i++ {
		if !progressTracker[i] && i%PROGRESS_BAR_SIZE == 0 {
			fmt.Print("\n")
		}
	}
}
