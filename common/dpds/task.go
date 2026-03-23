package dpds

import (
	"sync/atomic"
)

type TaskStatus int

const (
	StatusPending TaskStatus = iota
	StatusReady
	StatusRunning
	StatusCompleted
)

func (s TaskStatus) String() string {
	switch s {
	case StatusPending:
		return "PENDING"
	case StatusReady:
		return "READY"
	case StatusRunning:
		return "RUNNING"
	case StatusCompleted:
		return "COMPLETED"
	default:
		return "UNKNOWN"
	}
}

type Task struct {
	ID          string
	Priority    int
	Duration    int
	Dependencies []string
	TaskFunc    func() error

	status atomic.Value
}

func NewTask(id string, priority int, duration int, deps []string, fn func() error) *Task {
	t := &Task{
		ID:          id,
		Priority:    priority,
		Duration:    duration,
		Dependencies: deps,
		TaskFunc:    fn,
	}
	t.status.Store(StatusPending)
	return t
}

func (t *Task) GetStatus() TaskStatus {
	return t.status.Load().(TaskStatus)
}

func (t *Task) setStatus(status TaskStatus) {
	t.status.Store(status)
}

type taskNode struct {
	task     *Task
	priority int
	index    int
}