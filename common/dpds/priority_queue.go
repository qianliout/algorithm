package dpds

import (
	"sync"
)

type PriorityQueue struct {
	mu    sync.Mutex
	nodes []*taskNode
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		nodes: make([]*taskNode, 0),
	}
}

func (pq *PriorityQueue) Len() int {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return len(pq.nodes)
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.nodes[i], pq.nodes[j] = pq.nodes[j], pq.nodes[i]
	pq.nodes[i].index = i
	pq.nodes[j].index = j
}

func (pq *PriorityQueue) Push(task *Task) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	node := &taskNode{
		task:     task,
		priority: -task.Priority,
		index:    len(pq.nodes),
	}
	pq.nodes = append(pq.nodes, node)
	pq.bubbleUp(len(pq.nodes) - 1)
}

func (pq *PriorityQueue) Pop() *Task {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	if len(pq.nodes) == 0 {
		return nil
	}

	node := pq.nodes[0]
	last := len(pq.nodes) - 1
	pq.nodes[0] = pq.nodes[last]
	pq.nodes[0].index = 0
	pq.nodes = pq.nodes[:last]

	if len(pq.nodes) > 0 {
		pq.bubbleDown(0)
	}

	return node.task
}

func (pq *PriorityQueue) Remove(taskID string) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	for i, node := range pq.nodes {
		if node.task.ID == taskID {
			pq.removeAt(i)
			return
		}
	}
}

func (pq *PriorityQueue) removeAt(i int) {
	last := len(pq.nodes) - 1
	if i >= last {
		pq.nodes = pq.nodes[:last]
		return
	}

	pq.nodes[i] = pq.nodes[last]
	pq.nodes[i].index = i
	pq.nodes = pq.nodes[:last]

	pq.bubbleUp(i)
	pq.bubbleDown(i)
}

func (pq *PriorityQueue) bubbleUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if pq.nodes[i].priority <= pq.nodes[parent].priority {
			break
		}
		pq.Swap(i, parent)
		i = parent
	}
}

func (pq *PriorityQueue) bubbleDown(i int) {
	length := len(pq.nodes)
	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i

		if left < length && pq.nodes[left].priority > pq.nodes[largest].priority {
			largest = left
		}
		if right < length && pq.nodes[right].priority > pq.nodes[largest].priority {
			largest = right
		}
		if largest == i {
			break
		}
		pq.Swap(i, largest)
		i = largest
	}
}

func (pq *PriorityQueue) IsEmpty() bool {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return len(pq.nodes) == 0
}