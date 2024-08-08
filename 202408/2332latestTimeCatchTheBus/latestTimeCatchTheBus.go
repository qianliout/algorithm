package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(latestTimeCatchTheBus([]int{3}, []int{4}, 1))
}

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)
	j := 0
	n := len(passengers)
	c := 0
	for _, b := range buses {
		c = capacity
		for c > 0 && j < n && passengers[j] <= b {
			c--
			j++
		}
	}
	j = j - 1 // buses 只有一个值时，这里可能会是 -1

	ans := buses[len(buses)-1]
	if c == 0 { // 这里先判断再读取 j
		ans = passengers[j]
	}

	for j >= 0 && ans == passengers[j] {
		ans--
		j--
	}
	return ans
}
