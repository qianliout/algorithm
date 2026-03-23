package dpds

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestScheduler_Basic(t *testing.T) {
	s := NewScheduler(2)
	s.Start()

	var counter int
	var mu sync.Mutex

	for i := 0; i < 5; i++ {
		task := NewTask(
			fmt.Sprintf("task-%d", i),
			i%3,
			100,
			nil,
			func() error {
				mu.Lock()
				counter++
				mu.Unlock()
				time.Sleep(50 * time.Millisecond)
				return nil
			},
		)
		if err := s.Submit(task); err != nil {
			t.Fatal(err)
		}
	}

	time.Sleep(500 * time.Millisecond)

	mu.Lock()
	defer mu.Unlock()
	if counter != 5 {
		t.Errorf("expected 5 tasks completed, got %d", counter)
	}

	s.Shutdown()
}

func TestScheduler_Priority(t *testing.T) {
	s := NewScheduler(1)
	s.Start()

	var mu sync.Mutex
	order := make([]string, 0, 3)

	task0 := NewTask("p0", 0, 0, nil, func() error {
		time.Sleep(10 * time.Millisecond)
		mu.Lock()
		order = append(order, "p0")
		mu.Unlock()
		return nil
	})
	task1 := NewTask("p1", 1, 0, nil, func() error {
		time.Sleep(10 * time.Millisecond)
		mu.Lock()
		order = append(order, "p1")
		mu.Unlock()
		return nil
	})
	task2 := NewTask("p2", 2, 0, nil, func() error {
		time.Sleep(10 * time.Millisecond)
		mu.Lock()
		order = append(order, "p2")
		mu.Unlock()
		return nil
	})

	s.Submit(task0)
	s.Submit(task1)
	s.Submit(task2)

	time.Sleep(100 * time.Millisecond)

	mu.Lock()
	defer mu.Unlock()
	if len(order) != 3 {
		t.Fatalf("expected 3 tasks completed, got %d, order: %v", len(order), order)
	}
	if order[0] != "p0" || order[1] != "p1" || order[2] != "p2" {
		t.Errorf("expected order [p0 p1 p2], got %v", order)
	}

	s.Shutdown()
}

func TestScheduler_Dependency(t *testing.T) {
	s := NewScheduler(2)
	s.Start()

	var mu sync.Mutex
	order := make([]string, 0, 3)

	taskB := NewTask("B", 0, 0, []string{"A"}, func() error {
		time.Sleep(10 * time.Millisecond)
		mu.Lock()
		order = append(order, "B")
		mu.Unlock()
		return nil
	})
	taskA := NewTask("A", 0, 0, nil, func() error {
		time.Sleep(10 * time.Millisecond)
		mu.Lock()
		order = append(order, "A")
		mu.Unlock()
		return nil
	})

	s.Submit(taskB)
	s.Submit(taskA)

	time.Sleep(200 * time.Millisecond)

	mu.Lock()
	defer mu.Unlock()
	if len(order) != 2 || order[0] != "A" || order[1] != "B" {
		t.Errorf("expected order [A B], got %v", order)
	}

	s.Shutdown()
}

func TestScheduler_AddDependency(t *testing.T) {
	s := NewScheduler(1)
	s.Start()

	var mu sync.Mutex
	order := make([]string, 0, 3)

	taskA := NewTask("A", 0, 0, nil, func() error {
		time.Sleep(100 * time.Millisecond)
		mu.Lock()
		order = append(order, "A")
		mu.Unlock()
		return nil
	})
	taskC := NewTask("C", 0, 0, nil, func() error {
		time.Sleep(10 * time.Millisecond)
		mu.Lock()
		order = append(order, "C")
		mu.Unlock()
		return nil
	})

	s.Submit(taskA)
	s.Submit(taskC)

	time.Sleep(5 * time.Millisecond)

	if err := s.AddDependency("C", "A"); err != nil {
		t.Fatalf("AddDependency failed: %v", err)
	}

	time.Sleep(200 * time.Millisecond)

	mu.Lock()
	defer mu.Unlock()
	if len(order) != 2 {
		t.Fatalf("expected 2 tasks completed, got %d, order: %v", len(order), order)
	}
	if order[0] != "A" || order[1] != "C" {
		t.Errorf("expected order [A C], got order %v", order)
	}

	s.Shutdown()
}

func TestScheduler_CycleDetection(t *testing.T) {
	s := NewScheduler(2)
	s.Start()

	taskA := NewTask("A", 0, 0, []string{"B"}, nil)
	taskB := NewTask("B", 0, 0, []string{"A"}, nil)

	if err := s.Submit(taskA); err != nil {
		t.Fatal(err)
	}
	if err := s.Submit(taskB); err != nil {
		t.Fatal(err)
	}

	time.Sleep(100 * time.Millisecond)

	statusA, _ := s.GetTaskStatus("A")
	statusB, _ := s.GetTaskStatus("B")

	if statusA == StatusCompleted || statusB == StatusCompleted {
		t.Error("cyclic tasks should not complete")
	}

	s.Shutdown()
}

func TestScheduler_ConcurrentControl(t *testing.T) {
	s := NewScheduler(2)
	s.Start()

	maxRunning := 0
	var mu sync.Mutex
	var counter int

	for i := 0; i < 10; i++ {
		task := NewTask(
			fmt.Sprintf("task-%d", i),
			0,
			0,
			nil,
			func() error {
				mu.Lock()
				maxRunning++
				if maxRunning > 2 {
					t.Errorf("exceeded max concurrent: %d", maxRunning)
				}
				counter++
				mu.Unlock()
				time.Sleep(50 * time.Millisecond)
				mu.Lock()
				maxRunning--
				mu.Unlock()
				return nil
			},
		)
		s.Submit(task)
	}

	time.Sleep(500 * time.Millisecond)

	mu.Lock()
	defer mu.Unlock()
	if counter != 10 {
		t.Errorf("expected 10 tasks completed, got %d", counter)
	}

	s.Shutdown()
}

func TestScheduler_Shutdown(t *testing.T) {
	s := NewScheduler(2)
	s.Start()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		task := NewTask(
			fmt.Sprintf("task-%d", i),
			0,
			0,
			nil,
			func() error {
				time.Sleep(100 * time.Millisecond)
				wg.Done()
				return nil
			},
		)
		s.Submit(task)
	}

	go func() {
		time.Sleep(50 * time.Millisecond)
		s.Shutdown()
	}()

	wg.Wait()
}