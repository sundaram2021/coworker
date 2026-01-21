package utils

import (
	"os"
	"strings"
)

func LoadFiles(dir string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return []os.DirEntry{}, err
	}
	var cleanFiles []os.DirEntry
	for _, entry := range entries {
		if !strings.HasPrefix(entry.Name(), ".") {
			cleanFiles = append(cleanFiles, entry)
		}
	}
	return cleanFiles, nil
}
