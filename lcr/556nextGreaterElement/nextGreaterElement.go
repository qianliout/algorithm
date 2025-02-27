package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(nextGreaterElement(1221))
	fmt.Println(nextGreaterElement(230241))     // 230412
	fmt.Println(nextGreaterElement(2147483476)) // 230412
}

func nextGreaterElement(num int) int {
	ss := []byte(fmt.Sprintf("%d", num))

	n := len(ss)
	pre := -1
	next := -1
	for i := n - 1; i >= 0; i-- {
		if i-1 >= 0 && ss[i-1] < ss[i] {
			pre = i - 1
			next = i
			break
		}
	}
	if pre == -1 {
		return -1
	}
	// 排序
	sec := ss[next:]
	sort.Slice(sec, func(i, j int) bool { return sec[i] < sec[j] })
	for j := 0; j < len(sec); j++ {
		if sec[j] > ss[pre] {
			ss[pre], sec[j] = sec[j], ss[pre]
			ss = append(ss[:pre+1], sec...)
			ans, _ := strconv.Atoi(string(ss))
			if ans > math.MaxInt32 {
				return -1
			}
			return ans
		}
	}
	return -1
}
