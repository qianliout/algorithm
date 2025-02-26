package main

import (
	"fmt"
)

func main() {
	fmt.Println(validateBookSequences([]int{6, 7, 8, 9, 10, 11}, []int{11, 9, 8, 10, 6, 7}))
}

func validateBookSequences(putIn []int, takeOut []int) bool {
	st := make([]int, 0)
	le, ri, n := 0, 0, len(putIn)
	for ri < n {
		for le < n && (len(st) == 0 || st[len(st)-1] != takeOut[ri]) {
			st = append(st, putIn[le])
			le++
		}
		if len(st) == 0 || st[len(st)-1] != takeOut[ri] {
			return false
		}
		st = st[:len(st)-1]
		ri++
	}
	return true
}
