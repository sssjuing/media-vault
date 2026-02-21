package store

import (
	"path/filepath"

	"github.com/sssjuing/media-vault/internal/pkg/config"
	"github.com/sssjuing/media-vault/internal/pkg/downloader"
)

type Store interface {
	FindAll() []*downloader.Resource
	FindByID(id string) *downloader.Resource
	FindByName(name string) *downloader.Resource
	Add(r *downloader.Resource) error
	Delete(id string) error
	UpdateStatus(id string, index int, status int8) error
}

func New() Store {
	workPath := config.GetConfig().GetString("server.work_path")
	fs, _ := NewFileStore(filepath.Join(workPath, "store/data"), filepath.Join(workPath, "store/wal"))
	return fs
}
