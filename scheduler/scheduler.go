package scheduler

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Scheduler represents a struct definition for the Scheduler type.
//
// Example Usage:
//     s := Scheduler{
//         tasks: make(map[string]*Task),
//         mu:    sync.Mutex{},
//     }
//
// Inputs:
// - tasks: A map that stores tasks with their IDs as keys and pointers to Task objects as values.
// - mu: A mutex used for synchronization.
//
// Flow:
// 1. The Scheduler struct has two fields: tasks and mu.
// 2. The tasks field is a map that stores tasks with their IDs as keys and pointers to Task objects as values.
// 3. The mu field is a mutex that is used for synchronization when accessing or modifying the tasks map.
//
// Outputs:
// - None. 
type Scheduler struct {
	tasks map[string]*Task
	mu    sync.Mutex
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make(map[string]*Task),
	}
}

// AddTask adds a new task to the scheduler's task list, starts a new goroutine to execute the task, and returns an error if the task ID already exists.
//
// Parameters:
// - t (*Task): The task to be added to the scheduler's task list.
//
// Returns:
// - error: An error indicating that the task ID already exists if the task ID is already present in the task list. Otherwise, it returns nil.
func (s *Scheduler) AddTask(t *Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[t.ID]; exists {
		return errors.New("error: Task ID already exists")
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cancel = cancel
	s.tasks[t.ID] = t

	go s.execute(ctx, t)

	return nil
}

// RemoveTask removes a task from the scheduler's task list by canceling the task's execution and deleting it from the task map.
//
// Parameters:
// - id (string): The ID of the task to be removed.
func (s *Scheduler) RemoveTask(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if task, ok := s.tasks[id]; ok {
		task.Cancel()
		delete(s.tasks, id)
	}
}

// execute executes a task either after a specified duration or when the context is canceled.
//
// Inputs:
//   - ctx (context.Context): The context that allows the execution to be canceled.
//   - t (*Task): The task to be executed.
//
// Flow:
//  1. The method starts a select statement to wait for either the specified duration to elapse or the context to be canceled.
//  2. If the specified duration elapses before the context is canceled, the `time.After` channel will receive a value, and the code will proceed to the next case.
//  3. In the next case, the `TaskFunc` method of the task `t` will be executed.
//  4. If the context is canceled before the specified duration elapses, the `ctx.Done()` channel will receive a value, and the code will return, exiting the method.
//
// Outputs:
//   - None. The method does not return any value.
func (s *Scheduler) execute(ctx context.Context, t *Task) {
	select {
	case <-time.After(time.Until(t.ExecutionTime)):
		t.TaskFunc()
	case <-ctx.Done():
		return
	}
}
