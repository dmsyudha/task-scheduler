package scheduler

import (
	"context"
	"time"
)

type Task struct {
	ID            string
	Name          string
	ExecutionTime time.Time
	TaskFunc      func()
	Cancel        context.CancelFunc
}

func NewTask(id string, name string, executionTime time.Time, taskFunc func()) *Task {
	_, cancel := context.WithCancel(context.Background())

	return &Task{
		ID:            id,
		Name:          name,
		ExecutionTime: executionTime,
		TaskFunc:      taskFunc,
		Cancel:        cancel,
	}
}
