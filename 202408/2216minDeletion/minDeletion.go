package main

import (
	"fmt"
)

func main() {
	fmt.Println(minDeletion([]int{1, 1, 2, 2, 3, 3}))
}

func minDeletion(nums []int) int {
	cnt := 0
	n := len(nums)
	// 使用变量 cnt 代表已删除的元素个数，由于每次删除元素，剩余元素都会往前移动，因此遍历过程中，当前下标为 i−cnt。
	for i := 0; i < n; i++ {
		if (i-cnt)&1 == 0 && i+1 < n && nums[i+1] == nums[i] {
			cnt++
		}

	}
	if (n-cnt)%2 == 1 {
		cnt++
	}
	return cnt
}

// nums.length 为偶数
// 对所有满足 i % 2 == 0 的下标 i ，nums[i] != nums[i + 1] 均成立
