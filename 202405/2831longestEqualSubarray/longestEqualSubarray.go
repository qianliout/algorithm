package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestEqualSubarray([]int{1, 3, 2, 3, 1, 3}, 3))
	fmt.Println(longestEqualSubarray([]int{1, 1, 2, 2, 1, 1}, 2))
	fmt.Println(longestEqualSubarray2([]int{1, 1, 2, 2, 1, 1}, 2))
}

func longestEqualSubarray2(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	// 这个条件很重要 1 <= nums[i] <= nums.length
	pos := make([][]int, len(nums)+1)
	for i, ch := range nums {
		pos[ch] = append(pos[ch], i-len(pos[ch]))
	}
	ans := 1
	for _, po := range pos {
		if len(po) <= ans {
			continue
		}
		le := 0
		for ri, p := range po {
			for p-po[le] > k {
				le++
			}
			ans = max(ans, ri-le+1)
		}
	}
	return ans
}

func longestEqualSubarray(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	pos := make(map[int][]int)
	for i, ch := range nums {
		if pos[ch] == nil {
			pos[ch] = make([]int, 0)
		}
		pos[ch] = append(pos[ch], i)
	}
	ans := 1
	for _, ids := range pos {
		if len(ids) <= ans {
			continue
		}
		le := 0
		for ri := range ids {
			// 我们枚举每个元素作为等值元素，我们从哈希表中取出这个元素的下标列表 ids，然后我们定义两个指针 le 和 ri，用于维护一个窗口，
			// 使得窗口内的元素个数减去等值元素的个数，结果不超过 k。那么我们只需要求出最大的满足条件的窗口即可。
			// 窗口：ids[ri]-ids[le]+1 // 注意这里所说的窗口是上面数组所生成的窗口，而不由这个 ids 列表生成的窗口，ids 列表生成的窗口表示等值元素的个数
			// 等值元素的个数 (ri-le+1)
			for le <= ri && (ids[ri]-ids[le]+1)-(ri-le+1) > k {
				le++
			}
			// 这样写也是可以的
			// for ids[ri]-ids[le]-(ri-le) > k {
			// 	le++
			// }
			ans = max(ans, ri-le+1)
		}
	}
	return ans
}
