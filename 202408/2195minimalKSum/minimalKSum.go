package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minimalKSum([]int{1, 4, 25, 10, 25}, 2))
}

// 会超时
func minimalKSum1(nums []int, k int) int64 {
	cnt := make(map[int]int)
	for _, ch := range nums {
		cnt[ch]++
	}
	ans := 1
	sum := 0
	for k > 0 {
		if cnt[ans] <= 0 {
			k--
			sum += ans
			ans++
		} else {
			ans++
		}
	}
	return int64(sum)
}

func minimalKSum(nums []int, k int) int64 {
	nums = append(nums, 0, math.MaxInt)
	sort.Ints(nums)
	n := len(nums)
	var ans int64
	for i := 1; i < n; i++ {
		fill := nums[i] - nums[i-1] - 1
		if fill <= 0 {
			continue
		}
		if fill >= k {
			return ans + int64((nums[i-1]*2+1+k)*k/2)
		}
		ans += int64((nums[i-1] + nums[i]) * fill / 2) // 填充 fill 个数：等差数列求和
		k -= fill                                      // 更新剩余要填充的数字个数
	}
	return ans
}
