package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetUtilsContent() string {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("Could not get caller info")
	}
	sourceDir := filepath.Dir(currentFile)

	utilsPath := filepath.Join(sourceDir, "utils.prs")
	utilsContent, err := os.ReadFile(utilsPath)
	if err != nil {
		panic(err)
	}
	return string(utilsContent)
}
