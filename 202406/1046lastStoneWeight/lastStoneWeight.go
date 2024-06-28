package main

import (
	"container/heap"
	"fmt"

	. "outback/algorithm/common/commonHeap"
)

func main() {
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeDuplicates("aaaaaaaa"))
	fmt.Println(removeDuplicates("aaaaaaaaa"))
}

func lastStoneWeight(stones []int) int {
	hp := make(MaxHeap, 0)
	for _, ch := range stones {
		heap.Push(&hp, ch)
	}
	for len(hp) > 0 {
		if len(hp) == 1 {
			return hp[0]
		}
		m1, m2 := heap.Pop(&hp).(int), heap.Pop(&hp).(int)
		if m1 == m2 {
			continue
		}
		heap.Push(&hp, abs(m1-m2))
	}
	if hp.Len() == 0 {
		return 0
	}
	return hp[0]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func removeDuplicates(s string) string {
	stark := make([]byte, 0)
	n := len(s)
	i := 0
	for i < n {
		if len(stark) == 0 || s[i] != stark[len(stark)-1] {
			stark = append(stark, s[i])
			i++
			continue
		}
		if i < n && len(stark) > 0 && s[i] == stark[len(stark)-1] {
			i++
		}
		stark = stark[:len(stark)-1]
	}
	return string(stark)
}
