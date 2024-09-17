package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(maximumGood([][]int{{2, 0, 2, 2, 0}, {2, 2, 2, 1, 2}, {2, 2, 2, 1, 2}, {1, 2, 0, 2, 2}, {1, 0, 2, 1, 2}}))
}

func maximumGood(statements [][]int) int {
	ans := 0
	n := len(statements)
	for i := 1; i < 1<<n; i++ {
		ans = max(ans, check(i, statements))
	}
	return ans
}

func check(set int, sta [][]int) int {
	cnt := 0
	for j, row := range sta {
		if (set>>j)&1 == 0 {
			continue
		}
		flag := true
		// 假设j是好人
		for k, v := range row {
			if v < 2 && v != (set>>k)&1 {
				flag = false
				break
			}
		}
		if !flag {
			return 0
		}
		cnt++
	}
	return cnt
}

func check2(set int, sta [][]int) int {
	for j, row := range sta {
		if (set>>j)&1 == 0 {
			continue
		}
		flag := true
		// j是好人,然后去验证，全通过了就说明j 是真的好人，如果有一个没有通过，就说明 j 可能不是好人
		for k, v := range row {
			// 假设j是好人，那么他说的话就都是真的
			if v < 2 && v != (set>>k)&1 {
				flag = false
				break
			}
		}
		if !flag {
			return 0
		}
	}
	return bits.OnesCount(uint(set))
}
