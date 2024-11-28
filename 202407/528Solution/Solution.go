package main

import (
	"math/rand"
	"time"
)

func main() {

}

type Solution struct {
	rd    *rand.Rand
	total int64
	node  []Node
}

type Node struct {
	start, end int64
}

func Constructor(w []int) Solution {
	var total int64
	node := make([]Node, len(w))
	for i := 0; i < len(w); i++ {
		node[i] = Node{total, total + int64(w[i])}
		total += int64(w[i])
	}
	return Solution{rand.New(rand.NewSource(time.Now().UnixNano())), total, node}
}

func (this *Solution) PickIndex() int {
	idx := this.rd.Int63n(this.total) % this.total
	for i := 0; i < len(this.node); i++ {
		if this.node[i].start <= idx && this.node[i].end > idx {
			return i
		}
	}
	return -1
}
