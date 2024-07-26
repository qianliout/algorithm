package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumSum([]int{10, 12, 19, 14}))
}

func maximumSum(nums []int) int {
	cnt := make(map[int]int) // 数组和中的最大值
	ans := -1
	for _, ch := range nums {
		a := cal(ch)
		// 1 <= nums.length <= 105
		// 1 <= nums[i] <= 109
		// 数据范围决定了可以这样做
		if cnt[a] == 0 {
			cnt[a] = ch
			continue
		}
		pre := cnt[a]
		ans = max(ans, pre+ch)
		cnt[a] = max(pre, ch)
	}
	return ans
}

func cal(n int) int {
	ans := 0
	for n > 0 {
		ans += n % 10
		n = n / 10
	}
	return ans
}
