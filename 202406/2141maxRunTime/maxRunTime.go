package main

import (
	"sort"
)

func main() {

}

func maxRunTime(n int, batteries []int) int64 {
	sort.Slice(batteries, func(i, j int) bool { return batteries[i] > batteries[j] })
	sum := 0
	for _, ch := range batteries {
		sum += ch
	}
	for _, b := range batteries {
		if b <= sum/n {
			return int64(sum / n)
		}
		sum -= b
		n -= 1
	}
	return 0
}
