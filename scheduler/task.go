package scheduler

import (
	"context"
	"time"
)

// Task represents a task with an ID, name, execution time, task function, and cancel function.
type Task struct {
	ID            string             // the unique identifier for the task
	Name          string             // the name of the task
	ExecutionTime time.Time          // the time at which the task should be executed
	TaskFunc      func()             // the function to be executed when the task is triggered
	Cancel        context.CancelFunc // the function to cancel the task if it is removed during waiting time
}

// NewTask creates a new instance of the Task struct.
// It takes in four parameters: id (string), name (string), executionTime (time.Time), and taskFunc (function).
// It returns a pointer to the newly created Task object.
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
