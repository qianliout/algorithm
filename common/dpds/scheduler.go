package dpds

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Scheduler struct {
	maxConcurrent int32
	runningCount  int32

	tasks     map[string]*Task
	taskChan  chan *Task
	stopChan  chan struct{}
	waitGroup sync.WaitGroup

	depGraph   *DependencyGraph
	readyQueue *PriorityQueue
}

func NewScheduler(maxConcurrent int) *Scheduler {
	s := &Scheduler{
		maxConcurrent: int32(maxConcurrent),
		runningCount:  0,
		tasks:         make(map[string]*Task),
		taskChan:      make(chan *Task, 1024),
		stopChan:      make(chan struct{}),
		depGraph:      NewDependencyGraph(),
		readyQueue:    NewPriorityQueue(),
	}
	return s
}

func (s *Scheduler) Start() {
	go s.schedulerLoop()
}

func (s *Scheduler) schedulerLoop() {
	for {
		select {
		case task := <-s.taskChan:
			s.enqueueTask(task)
		case <-s.stopChan:
			return
		}
	}
}

func (s *Scheduler) enqueueTask(task *Task) {
	if s.depGraph.HasDependencies(task.ID) {
		task.setStatus(StatusPending)
	} else {
		task.setStatus(StatusReady)
		s.readyQueue.Push(task)
	}
	s.schedule()
}

func (s *Scheduler) schedule() {
	for atomic.LoadInt32(&s.runningCount) < s.maxConcurrent {
		task := s.readyQueue.Pop()
		if task == nil {
			return
		}
		s.startTask(task)
	}
}

func (s *Scheduler) startTask(task *Task) {
	task.setStatus(StatusRunning)
	atomic.AddInt32(&s.runningCount, 1)
	s.waitGroup.Add(1)

	go func(t *Task) {
		defer func() {
			s.handleTaskDone(t.ID)
			s.waitGroup.Done()
		}()
		if t.TaskFunc != nil {
			t.TaskFunc()
		}
	}(task)
}

func (s *Scheduler) handleTaskDone(taskID string) {
	atomic.AddInt32(&s.runningCount, -1)

	task := s.tasks[taskID]
	if task != nil {
		task.setStatus(StatusCompleted)
	}

	for _, dependentID := range s.depGraph.GetDependents(taskID) {
		s.checkAndReady(dependentID)
	}

	s.schedule()
}

func (s *Scheduler) checkAndReady(taskID string) {
	task := s.tasks[taskID]
	if task == nil || task.GetStatus() != StatusPending {
		return
	}

	deps := s.depGraph.GetDependencies(taskID)
	for _, depID := range deps {
		depTask := s.tasks[depID]
		if depTask == nil || depTask.GetStatus() != StatusCompleted {
			return
		}
	}

	task.setStatus(StatusReady)
	s.readyQueue.Push(task)
}

func (s *Scheduler) Submit(task *Task) error {
	if task.ID == "" {
		return fmt.Errorf("task ID cannot be empty")
	}

	s.tasks[task.ID] = task

	for _, depID := range task.Dependencies {
		if err := s.depGraph.AddDependency(task.ID, depID); err != nil {
			return err
		}
	}

	s.taskChan <- task
	return nil
}

func (s *Scheduler) AddDependency(taskID, dependsOnID string) error {
	if _, ok := s.tasks[taskID]; !ok {
		return fmt.Errorf("task %s not found", taskID)
	}
	if _, ok := s.tasks[dependsOnID]; !ok {
		return fmt.Errorf("dependency task %s not found", dependsOnID)
	}

	task := s.tasks[taskID]
	status := task.GetStatus()
	if status == StatusRunning {
		return fmt.Errorf("cannot add dependency to running task %s", taskID)
	}
	if status == StatusCompleted {
		return fmt.Errorf("cannot add dependency to completed task %s", taskID)
	}

	if err := s.depGraph.AddDependency(taskID, dependsOnID); err != nil {
		return err
	}

	if status == StatusReady {
		s.readyQueue.Remove(taskID)
		task.setStatus(StatusPending)
	}

	s.checkAndReady(taskID)
	s.schedule()
	return nil
}

func (s *Scheduler) RemoveDependency(taskID, dependsOnID string) error {
	s.depGraph.RemoveDependency(taskID, dependsOnID)
	s.checkAndReady(taskID)
	s.schedule()
	return nil
}

func (s *Scheduler) GetTaskStatus(taskID string) (TaskStatus, error) {
	task, ok := s.tasks[taskID]
	if !ok {
		return StatusPending, fmt.Errorf("task %s not found", taskID)
	}
	return task.GetStatus(), nil
}

func (s *Scheduler) Shutdown() {
	close(s.stopChan)
	s.waitGroup.Wait()
}

func (s *Scheduler) GetRunningCount() int32 {
	return atomic.LoadInt32(&s.runningCount)
}

func (s *Scheduler) GetPendingCount() int {
	return s.readyQueue.Len()
}

func (s *Scheduler) GetReadyCount() int {
	return s.readyQueue.Len()
}