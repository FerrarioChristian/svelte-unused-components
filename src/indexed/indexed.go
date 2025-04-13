package indexed

func FindUnusedNormal(svelte_files []string, recursive bool) []string {
	if recursive {
		return getUnusedFilesRecursive(svelte_files)
	}
	return getUnusedFiles(svelte_files)
}

func FindUnusedConcurrent(svelte_files []string, recursive bool) []string {
	if recursive {
		return getUnusedFilesRecursiveConcurrent(svelte_files)
	}
	return getUnusedFilesConcurrent(svelte_files)
}
