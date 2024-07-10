package main

import (
	"sort"
)

func main() {

}

func checkIfCanBreak(s1 string, s2 string) bool {
	return help(s1, s2)
}

func help(s1, s2 string) bool {
	ss1, ss2 := []byte(s1), []byte(s2)
	sort.Slice(ss1, func(i, j int) bool { return ss1[i] < ss1[j] })
	sort.Slice(ss2, func(i, j int) bool { return ss2[i] < ss2[j] })
	// 检查小于
	less := true
	for i := range ss1 {
		if ss1[i] > ss2[i] {
			less = false
		}
	}
	more := true

	for i := range ss1 {
		if ss1[i] < ss2[i] {
			more = false
		}
	}
	return less || more
}
