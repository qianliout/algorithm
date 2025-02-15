package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(getMinSwaps("5489355142", 4))
	fmt.Println(getMinSwaps("059", 5))
}

func getMinSwaps(num string, k int) int {
	n := len(num)
	ss := []byte(num)
	nums := make([]int, n)
	for i, ch := range num {
		nums[i] = int(ch) - int('0')
	}
	// k = k % n // 这一步不能加
	for k > 0 {
		nextPermutation(nums)
		k--
	}
	ans := 0
	for i := 0; i < n; i++ {
		if int(ss[i])-int('0') == nums[i] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if nums[i] == int(ss[j])-int('0') {
				for m := j - 1; m >= i; m-- {
					ans++
					ss[m], ss[m+1] = ss[m+1], ss[m]
				}
				break
			}
		}
	}
	return ans
}

func nextPermutation(nums []int) {
	n := len(nums)
	start, end, sub := -1, -1, math.MaxInt64
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			start = i - 1
			break
		}
	}
	if start == -1 {
		// 说明已经是最大的排列，也就是降序了，此时按题目的意思，应该换成升序
		// 如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）
		sort.Ints(nums)
		// le, ri := 0, n-1
		// for le < ri {
		// 	nums[le], nums[ri] = nums[ri], nums[le]
		// 	le++
		// 	ri--
		// }
		return
	}
	end = start

	for j := start + 1; j < n; j++ {
		if nums[j] > nums[start] && nums[j]-nums[start] < sub {
			end = j
			sub = nums[j] - nums[start]
		}
	}
	// 把后面的排序
	nums[start], nums[end] = nums[end], nums[start]
	if start+1 < n {
		sort.Ints(nums[start+1:])
	}
}

func nextPermutation2(nums []int) {
	n := len(nums)
	start, end, sub := -1, -1, math.MaxInt64
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			start = i - 1
			break
		}
	}

	if start == -1 {
		le, ri := 0, len(nums)-1
		for le < ri {
			nums[le], nums[ri] = nums[ri], nums[le]
			le++
			ri--
		}
		return
	}

	end = start

	for j := start + 1; j < len(nums); j++ {
		if nums[j] > nums[start] && nums[j]-nums[start] < sub {
			end = j
			sub = nums[j] - nums[start]
		}
	}
	nums[start], nums[end] = nums[end], nums[start]
	if start+1 < n {
		sort.Ints(nums[start+1:])
	}
}
