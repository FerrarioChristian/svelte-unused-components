package inverted

func FindUnusedNormal(files []string, recursive bool) []string {
	if recursive {
		return getUnusedFilesRecursive(files)
	}
	return getUnusedFiles(files)
}
