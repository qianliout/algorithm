package main

import (
	"sort"
)

func main() {

}

func orderlyQueue(s string, k int) string {
	ss := []byte(s)
	if k >= 2 {
		// k>2 时一定可以两两相交换，模拟冒泡排序，
		sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })
		return string(ss)
	}
	ans := s
	for i := 0; i < len(s); i++ {
		ss = append(ss, ss[0])
		ss = ss[1:]
		ans = min(ans, string(ss))
	}
	return ans
}
