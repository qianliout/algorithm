package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestWPI([]int{9, 9, 6, 0, 6, 6, 9}))
	fmt.Println(longestWPI([]int{6, 6, 9}))
}

/*
「劳累天数大于不劳累天数」等价于「劳累天数减去不劳累天数大于 0」。
那么把劳累的一天视作 nums[i]=1，不劳累的一天视作 nums[i]=−1，则问题变为：
计算 nums的最长子数组，其元素和大于 0。
既然说到了「子数组的元素和」，那么利用前缀和 sum，将问题变为：
找到两个下标 i 和 j，满足 j<i 且 sum[j]<sum[i],最大化 i−j的值
*/
func longestWPI(hours []int) int {
	sum := make([]int, len(hours)+1)
	// stark保存的下标
	stark := make([]int, 0)
	stark = append(stark, 0) // 单调递减栈，初始化0

	for i, ch := range hours {
		if ch > 8 {
			sum[i+1] = sum[i] + 1
		} else {
			sum[i+1] = sum[i] - 1
		}
		// 这一步是重点，当 sum[i+1]>
		if sum[i+1] < sum[stark[len(stark)-1]] {
			stark = append(stark, i+1)
		}
	}
	ans := 0
	for i := len(hours); i > 0; i-- {
		for len(stark) > 0 && sum[i] > sum[stark[len(stark)-1]] {
			ans = max(ans, i-stark[len(stark)-1])
			stark = stark[:len(stark)-1]
		}
	}
	return ans
}
