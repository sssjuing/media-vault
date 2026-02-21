package oss

import (
	"context"
	"fmt"
	"testing"
)

func TestGetFiles(t *testing.T) {
	result, _ := GetFiles("pictures/covers")
	for _, item := range result {
		fmt.Println(*item)
	}
}

func TestUploadLocalFile(t *testing.T) {
	info, err := UploadLocalFile(context.Background(), "/tmp/sss.ts", "/data/zq3-jx3/zq3-jx3.ts")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}
