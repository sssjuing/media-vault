package utils

import "testing"

func TestListFilesInDir(t *testing.T) {
	files, err := ListFilesInDir("/data")
	if err != nil {
		t.Error(err)
		return
	}
	for _, f := range files {
		t.Logf("path: %s, size: %d MB", f.Path, f.Size/1024/1024)
	}
}
