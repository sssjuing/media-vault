package taskqueue

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type MyTask struct{}

func (mt *MyTask) Execute() {
	fmt.Println("execute task")
	time.Sleep(2 * time.Second)
}

func TestTaskQueue(t *testing.T) {
	tq := New(context.Background())
	for range 5 {
		tq.AddTask(&MyTask{})
	}
	tq.Close()
	tq.Wait()
}

func TestTaskQueueWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 4500*time.Millisecond)
	defer cancel()
	tq := New(ctx)
	for range 5 {
		tq.AddTask(&MyTask{})
	}
	tq.Close()
	tq.Wait()
}
