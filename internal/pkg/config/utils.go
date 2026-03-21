package config

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func findModuleRoot(path string) string {
	if path == "/" {
		return ""
	}
	modFilePath := filepath.Join(path, "go.mod")
	if _, err := os.Stat(modFilePath); err == nil {
		return path
	}
	return findModuleRoot(filepath.Dir(path))
}

func GetRootPath() string {
	_, filename, _, _ := runtime.Caller(0)
	localPath := path.Dir(filename)
	return findModuleRoot(localPath)
}
