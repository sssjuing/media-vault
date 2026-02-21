package store

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/samber/lo"
	"github.com/sssjuing/media-vault/internal/pkg/downloader"
)

// WALEntry 表示 WAL 日志条目
type WALEntry struct {
	ResourceID string
	Status     int8 // 用于 SegmentRow 的 Status
	Timestamp  int64
}

// FileStore 实现 Store 接口
type FileStore struct {
	dir       string        // 数据文件目录
	walDir    string        // WAL 日志目录
	resources sync.Map      // 缓存: map[string]*Resource
	nameIndex sync.Map      // 索引: map[string]string (Filename -> ID)
	mu        sync.RWMutex  // 保护文件和 WAL 写入
	updateCh  chan update   // 更新队列
	closed    chan struct{} // 关闭信号
}

// update 表示一次更新
type update struct {
	resourceID string
	index      int
	status     int8
}

// NewFileStore 初始化存储
func NewFileStore(dataDir, walDir string) (*FileStore, error) {
	// 确保目录存在
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("创建数据目录失败: %v", err)
	}
	if err := os.MkdirAll(walDir, 0755); err != nil {
		return nil, fmt.Errorf("创建 WAL 目录失败: %v", err)
	}

	s := &FileStore{
		dir:      dataDir,
		walDir:   walDir,
		updateCh: make(chan update, 100),
		closed:   make(chan struct{}),
	}

	// 加载现有数据文件
	if err := s.loadExistingFiles(); err != nil {
		return nil, err
	}

	// 重放 WAL 日志
	if err := s.replayWAL(); err != nil {
		return nil, err
	}

	// 启动异步持久化
	go s.persister()

	return s, nil
}

// loadExistingFiles 加载数据文件到内存
func (s *FileStore) loadExistingFiles() error {
	entries, err := os.ReadDir(s.dir)
	if err != nil {
		return fmt.Errorf("读取数据目录失败: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		path := filepath.Join(s.dir, entry.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("读取文件 %s 失败: %v", path, err)
		}

		var res downloader.Resource
		if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&res); err != nil {
			return fmt.Errorf("解码文件 %s 失败: %v", path, err)
		}

		if res.ID != entry.Name() {
			return fmt.Errorf("文件 %s ID 不匹配: %s", path, res.ID)
		}

		res.Downloading = false
		res.Success = lo.EveryBy(res.SegmentList, func(sr downloader.SegmentRow) bool {
			return sr.Status == 1
		})

		s.resources.Store(res.ID, &res)
		s.nameIndex.Store(res.Filename, res.ID)
	}

	return nil
}

// replayWAL 重放 WAL 日志
func (s *FileStore) replayWAL() error {
	entries, err := os.ReadDir(s.walDir)
	if err != nil {
		return fmt.Errorf("读取 WAL 目录失败: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		path := filepath.Join(s.walDir, entry.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("读取 WAL 文件 %s 失败: %v", path, err)
		}

		var entries []WALEntry
		if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&entries); err != nil {
			return fmt.Errorf("解码 WAL 文件 %s 失败: %v", path, err)
		}

		for _, e := range entries {
			res, ok := s.resources.Load(e.ResourceID)
			if !ok {
				continue
			}
			r := res.(*downloader.Resource)
			s.resources.Store(e.ResourceID, r)
		}

		// 删除已重放的 WAL 文件
		if err := os.Remove(path); err != nil {
			return fmt.Errorf("删除 WAL 文件 %s 失败: %v", path, err)
		}
	}

	return nil
}

// FindAll 返回所有 Resource
func (s *FileStore) FindAll() []*downloader.Resource {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []*downloader.Resource
	s.resources.Range(func(_, value interface{}) bool {
		r := value.(*downloader.Resource)
		result = append(result, r)
		return true
	})
	return result
}

// FindByID 按 ID 查找 Resource
func (s *FileStore) FindByID(id string) *downloader.Resource {
	s.mu.RLock()
	defer s.mu.RUnlock()

	res, ok := s.resources.Load(id)
	if !ok {
		return nil
	}
	return res.(*downloader.Resource)
}

// FindByName 按 Filename 查找 Resource
func (s *FileStore) FindByName(name string) *downloader.Resource {
	s.mu.RLock()
	defer s.mu.RUnlock()

	id, ok := s.nameIndex.Load(name)
	if !ok {
		return nil
	}
	res, ok := s.resources.Load(id)
	if !ok {
		return nil
	}
	return res.(*downloader.Resource)
}

// Add 添加或更新 Resource
func (s *FileStore) Add(r *downloader.Resource) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 更新内存缓存
	s.resources.Store(r.ID, r)
	s.nameIndex.Store(r.Filename, r.ID)

	// 写入文件
	s.saveResource(r.ID)
	return nil
}

// Delete 删除 Resource
func (s *FileStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 删除内存缓存
	res, ok := s.resources.LoadAndDelete(id)
	if !ok {
		return fmt.Errorf("resource %s not exist", id)
	}
	r := res.(*downloader.Resource)
	s.nameIndex.Delete(r.Filename)

	// 删除切片存放目录
	os.RemoveAll(r.TempDir)

	// 删除文件
	path := filepath.Join(s.dir, id)
	os.Remove(path) // 忽略错误（如文件不存在）

	return nil
}

// UpdateStatus 更新 SegmentRow 的 Status
func (s *FileStore) UpdateStatus(resourceID string, index int, status int8) error {
	res, ok := s.resources.Load(resourceID)
	if !ok {
		return fmt.Errorf("资源 %s 不存在", resourceID)
	}
	r := res.(*downloader.Resource)
	if index < 0 || index >= len(r.SegmentList) {
		return fmt.Errorf("索引越界: %d", index)
	}

	// 更新内存缓存
	r.SegmentList[index].Status = status
	s.resources.Store(resourceID, r)

	// 写入 WAL 和异步持久化
	select {
	case s.updateCh <- update{resourceID: resourceID, index: index, status: status}:
	default:
		// 队列满，立即持久化
		s.mu.Lock()
		if err := s.writeWAL(resourceID, index, status); err != nil {
			s.mu.Unlock()
			return err
		}
		if err := s.saveResource(resourceID); err != nil {
			s.mu.Unlock()
			return err
		}
		s.mu.Unlock()
	}

	return nil
}

// // GetSegment 获取单个 SegmentRow
// func (s *FileStore) GetSegment(resourceID string, index int) (downloader.SegmentRow, error) {
// 	s.mu.RLock()
// 	defer s.mu.RUnlock()

// 	res, ok := s.resources.Load(resourceID)
// 	if !ok {
// 		return downloader.SegmentRow{}, fmt.Errorf("资源 %s 不存在", resourceID)
// 	}
// 	r := res.(*downloader.Resource)
// 	if index < 0 || index >= len(r.SegmentList) {
// 		return downloader.SegmentRow{}, fmt.Errorf("索引越界: %d", index)
// 	}
// 	return r.SegmentList[index], nil
// }

// writeWAL 写入 WAL 日志
func (s *FileStore) writeWAL(resourceID string, index int, status int8) error {
	entry := WALEntry{
		ResourceID: resourceID,
		Status:     status,
		Timestamp:  time.Now().UnixNano(),
	}

	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode([]WALEntry{entry}); err != nil {
		return fmt.Errorf("编码 WAL 失败: %v", err)
	}

	path := filepath.Join(s.walDir, fmt.Sprintf("%d.wal", entry.Timestamp))
	return os.WriteFile(path, buf.Bytes(), 0644)
}

// saveResource 持久化 Resource 到文件
func (s *FileStore) saveResource(resourceID string) error {
	res, ok := s.resources.Load(resourceID)
	if !ok {
		return fmt.Errorf("资源 %s 不存在", resourceID)
	}
	r := res.(*downloader.Resource)

	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(r); err != nil {
		return fmt.Errorf("编码资源失败: %v", err)
	}

	path := filepath.Join(s.dir, resourceID)
	return os.WriteFile(path, buf.Bytes(), 0644)
}

// persister 异步持久化
func (s *FileStore) persister() {
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	count := 0
	const batchSize = 50
	updates := make(map[string]map[int]WALEntry) // resourceID -> index -> WALEntry

	for {
		select {
		case <-s.closed:
			s.mu.Lock()
			for resourceID := range updates {
				s.saveResource(resourceID)
			}
			s.mu.Unlock()
			return
		case <-ticker.C:
			if count > 0 {
				s.mu.Lock()
				for resourceID := range updates {
					s.saveResource(resourceID)
				}
				s.mu.Unlock()
				count = 0
				updates = make(map[string]map[int]WALEntry)
			}
		case u := <-s.updateCh:
			count++
			if _, ok := updates[u.resourceID]; !ok {
				updates[u.resourceID] = make(map[int]WALEntry)
			}
			updates[u.resourceID][u.index] = WALEntry{
				ResourceID: u.resourceID,
				Status:     u.status,
				Timestamp:  time.Now().UnixNano(),
			}

			s.mu.Lock()
			if err := s.writeWAL(u.resourceID, u.index, u.status); err != nil {
				s.mu.Unlock()
				continue
			}
			s.mu.Unlock()

			if count >= batchSize {
				s.mu.Lock()
				for resourceID := range updates {
					s.saveResource(resourceID)
				}
				s.mu.Unlock()
				count = 0
				updates = make(map[string]map[int]WALEntry)
			}
		}
	}
}

// Close 关闭存储
func (s *FileStore) Close() error {
	close(s.closed)
	return nil
}
