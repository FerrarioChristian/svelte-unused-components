package inverted

func FindUnusedNormal(files []string, recursive bool) []string {
	if recursive {
		return getUnusedFilesRecursive(files)
	}
	return getUnusedFiles(files)
}

func FindUnusedConcurrent(svelte_files []string, recursive bool) []string {
	if recursive {
		return getUnusedFilesRecursiveConcurrent(svelte_files)
	}
	return getUnusedFilesConcurrent(svelte_files)
}

func FindUnusedWorkers(svelte_files []string, recursive bool) []string {
	if recursive {
		return getUnusedFilesRecursiveWorkers(svelte_files)
	}
	return getUnusedFilesWorkers(svelte_files)
}
