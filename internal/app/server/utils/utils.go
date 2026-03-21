package utils

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
)

func ParseUint(str string) (uint, error) {
	actressId, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(actressId), nil
}

func ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type File struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
}

func ListFilesInDir(targetDir string) ([]File, error) {
	files := make([]File, 0, 10)
	if err := filepath.WalkDir(targetDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}
			files = append(files, File{path, info.Size()})
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return files, nil
}
