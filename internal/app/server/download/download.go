package download

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	downloadstore "github.com/sssjuing/media-vault/internal/app/server/download/store"
	"github.com/sssjuing/media-vault/internal/pkg/config"
	"github.com/sssjuing/media-vault/internal/pkg/downloader"
	"github.com/sssjuing/media-vault/pkg/taskqueue"
)

var hashSet mapset.Set[string] // 任务队列中的资源 id 集合
var tq *taskqueue.TaskQueue
var store downloadstore.Store

func Init(logger *slog.Logger) {
	hashSet = mapset.NewSet[string]()
	tq = taskqueue.New(context.Background())
	store = downloadstore.New()
}

func createDownloader(resource *downloader.Resource, logger *slog.Logger) *downloader.Downloader {
	wsf := downloader.WithSegmentFinish(func(sr *downloader.SegmentRow, index int) {
		logger.Debug(fmt.Sprintf("downloading: %s, %d", sr.Path, sr.Status))
		store.UpdateStatus(resource.ID, index, sr.Status)
	})
	wtf := downloader.WithFinally(func(r *downloader.Resource, err error) {
		logger.Info(fmt.Sprintf("task %s finished, total: %d", r.Filename, len(r.SegmentList)))
		if err != nil {
			logger.Error(fmt.Sprintf("downloading: %s, %d, %s", r.Filename, len(r.SegmentList), err.Error()))
		}
		hashSet.Remove(r.ID)
	})
	wl := downloader.WithLogger(logger)
	return downloader.New(resource, wsf, wtf, wl)
}

func AddTask(url, referer, name string, logger *slog.Logger) error {
	workPath := config.GetConfig().GetString("server.work_path")
	resource := downloader.NewResource(url, referer, name, filepath.Join(workPath, "tmp", name), nil)
	if ok := hashSet.Add(resource.ID); !ok {
		return fmt.Errorf("task %s is already included in the queue, please do not add it repeatedly", name)
	}
	if store.FindByID(resource.ID) == nil {
		store.Add(resource)
	}
	d := createDownloader(resource, logger)
	d.PreExecute()
	if ok := tq.AddTask(d); !ok {
		return fmt.Errorf("downloader has been closed")
	}
	return nil
}

func RetryTask(r *downloader.Resource, logger *slog.Logger) error {
	if r.Success {
		return fmt.Errorf("cannot add already successful tasks to the queue")
	}
	if ok := hashSet.Add(r.ID); !ok {
		return fmt.Errorf("task %s is already included in the queue, please do not add it repeatedly", r.Filename)
	}
	d := createDownloader(r, logger)
	d.PreExecute()
	if ok := tq.AddTask(d); !ok {
		return fmt.Errorf("downloader has been closed")
	}
	return nil
}

// GetHashSet
func GetHashSet() mapset.Set[string] {
	return hashSet
}

// GetStore
func GetStore() downloadstore.Store {
	return store
}
