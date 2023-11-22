package scheduler

import (
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	s := NewScheduler()

	t1 := NewTask("1", "Task 1", time.Now().Add(1*time.Second), func() {
		t.Log("Executing Task 1")
	})

	s.AddTask(t1)

	if _, ok := s.tasks["1"]; !ok {
		t.Error("Task 1 was not added to the scheduler")
	}

	s.RemoveTask("1")

	if _, ok := s.tasks["1"]; ok {
		t.Error("Task 1 was not removed from the scheduler")
	}
}
