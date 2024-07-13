package main

import (
	"fmt"
)

func main() {
	fmt.Println(minOperations([]int{1, 5}))
	fmt.Println(minOperations([]int{3, 2, 2, 4}))
}

func minOperations(nums []int) int {
	cnt := 0
	for {
		for i, ch := range nums {
			// 如果是奇数就先执行一次加一操作
			if ch&1 == 1 {
				cnt++
			}
			// 反着着，如果是奇数，会把末尾的1抹掉，但是上面已经做了cnt++了，所以不会出错
			nums[i] = nums[i] >> 1
		}
		if sum(nums) == 0 {
			break
		}
		cnt++
	}
	return cnt
}

func sum(nums []int) int {
	s := 0
	for _, ch := range nums {
		s += ch
	}
	return s
}
