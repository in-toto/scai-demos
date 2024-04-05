package fileio

import (
	"os"
	"path/filepath"
)

func ReadFileIntoMap(filename string, fileMap map[string][]byte) error {
	name := filepath.Base(filename)
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	fileMap[name] = content
	return nil
}
