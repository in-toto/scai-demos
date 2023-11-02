package fileio

import (
	"os"
	"path/filepath"
	"strings"
)

func HasJSONExt(filename string) bool {
	return strings.HasSuffix(filename, ".json")
}

func CreateOutDir(filename string) error {
	outDir := filepath.Dir(filename)

	return os.MkdirAll(outDir, 0755)
}
