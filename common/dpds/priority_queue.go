package dpds

import "sync"

type PriorityQueue struct {
	mu    sync.Mutex
	items []taskItem
}

type taskItem struct {
	task     *Task
	priority int
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{items: make([]taskItem, 0)}
}

func (pq *PriorityQueue) Len() int {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return len(pq.items)
}

func (pq *PriorityQueue) Push(task *Task) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	item := taskItem{task: task, priority: -task.Priority}
	pq.items = append(pq.items, item)
	pq.bubbleUp(len(pq.items) - 1)
}

func (pq *PriorityQueue) Pop() *Task {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	if len(pq.items) == 0 {
		return nil
	}
	item := pq.items[0]
	last := len(pq.items) - 1
	pq.items[0] = pq.items[last]
	pq.items = pq.items[:last]
	if len(pq.items) > 0 {
		pq.bubbleDown(0)
	}
	return item.task
}

func (pq *PriorityQueue) Remove(taskID string) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	for i, item := range pq.items {
		if item.task.ID == taskID {
			last := len(pq.items) - 1
			pq.items[i] = pq.items[last]
			pq.items = pq.items[:last]
			if i < len(pq.items) {
				pq.bubbleUp(i)
				pq.bubbleDown(i)
			}
			return
		}
	}
}

func (pq *PriorityQueue) IsEmpty() bool {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return len(pq.items) == 0
}

func (pq *PriorityQueue) bubbleUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if pq.items[i].priority <= pq.items[parent].priority {
			break
		}
		pq.items[i], pq.items[parent] = pq.items[parent], pq.items[i]
		i = parent
	}
}

func (pq *PriorityQueue) bubbleDown(i int) {
	n := len(pq.items)
	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i
		if left < n && pq.items[left].priority > pq.items[largest].priority {
			largest = left
		}
		if right < n && pq.items[right].priority > pq.items[largest].priority {
			largest = right
		}
		if largest == i {
			break
		}
		pq.items[i], pq.items[largest] = pq.items[largest], pq.items[i]
		i = largest
	}
}
