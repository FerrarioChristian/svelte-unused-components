package bench

import (
	"fmt"
	"time"
)

func Benchmark(name string, f func() []string) {
	start := time.Now()
	result := f()
	duration := time.Since(start)
	fmt.Printf("== %s ==\nUnused: %d\nTime spent: %v\n\n", name, len(result), duration)
}
