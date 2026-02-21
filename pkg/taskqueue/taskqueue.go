package taskqueue

import (
	"context"
	"sync"
)

type Task interface {
	Execute()
}

type TaskQueue struct {
	ctx      context.Context
	taskChan chan Task // 任务通道
	// errorChan chan error     // 错误通知通道
	wg     sync.WaitGroup // 等待组用于等待所有任务完成
	closed bool           // 队列关闭标志
	mutex  sync.Mutex     // 保护 closed 字段
}

func New(ctx context.Context) *TaskQueue {
	d := &TaskQueue{
		ctx:      ctx,
		taskChan: make(chan Task, 100), // 带缓冲的通道
		// errorChan: make(chan error),     // 错误通道
	}
	d.wg.Add(1)
	go d.run() // 启动任务处理协程
	return d
}

// 添加任务到队列, 此方法被多个协程同时调用时, closed 用于保证线程安全
func (tq *TaskQueue) AddTask(r Task) bool {
	tq.mutex.Lock()
	defer tq.mutex.Unlock() // 使用 defer 简化锁管理
	if tq.closed {
		tq.mutex.Unlock()
		return false
	}
	tq.taskChan <- r
	return true
}

// // 返回错误通知通道
// func (tq *TaskQueue) Errors() <-chan error {
// 	return tq.errorChan
// }

// 关闭队列
func (tq *TaskQueue) Close() {
	tq.mutex.Lock()
	defer tq.mutex.Unlock()
	if !tq.closed {
		tq.closed = true
		close(tq.taskChan)
		// close(tq.errorChan) 这里不能 close, 因为 run 中可能还在写入
	}
}

// 等待所有任务完成
func (tq *TaskQueue) Wait() {
	tq.wg.Wait()
}

// 任务处理循环
func (tq *TaskQueue) run() {
	defer tq.wg.Done() // TODO: 可以在这里 close errorChan
	for {
		select {
		case <-tq.ctx.Done():
			return
		case t, ok := <-tq.taskChan:
			if !ok {
				return
			}
			t.Execute()
		}
	}
}
