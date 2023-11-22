package scheduler

import (
	"context"
	"sync"
	"time"
)

type Scheduler struct {
	tasks map[string]*Task
	mu    sync.Mutex
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make(map[string]*Task),
	}
}

func (s *Scheduler) AddTask(t *Task) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ctx, cancel := context.WithCancel(context.Background())
	t.Cancel = cancel
	s.tasks[t.ID] = t

	go s.execute(ctx, t)
}

func (s *Scheduler) RemoveTask(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if task, ok := s.tasks[id]; ok {
		task.Cancel()
		delete(s.tasks, id)
	}
}

func (s *Scheduler) execute(ctx context.Context, t *Task) {
	select {
	case <-time.After(time.Until(t.ExecutionTime)):
		t.TaskFunc()
	case <-ctx.Done():
		return
	}
}
