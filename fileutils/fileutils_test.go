package fileutils_test

import (
	"svelte-unused-components/fileutils"
	"testing"
)

func TestWriteResultsToFile(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		filename string
		lines    []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileutils.WriteResultsToFile(tt.filename, tt.lines)
		})
	}
}
