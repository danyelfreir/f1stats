package util

import "strings"

func CleanAndSplitURL(path string) []string {
	newPath := strings.Clone(path)
	newPath = strings.TrimPrefix(newPath, "/")
	newPath = strings.TrimSuffix(newPath, "/")
	pathStubs := strings.Split(newPath, "/")
	return pathStubs
}
