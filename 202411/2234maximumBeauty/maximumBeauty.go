package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumBeauty([]int{1, 3, 1, 1}, 7, 6, 12, 1))
	// fmt.Println(maximumBeauty([]int{2, 4, 5, 31}, 10, 15, 2, 6))
}

func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	sort.Ints(flowers)
	fullCnt := 0
	for _, ch := range flowers {
		if ch >= target {
			fullCnt++
		}
	}
	flowers = flowers[:len(flowers)-fullCnt]
	n := len(flowers)
	if len(flowers) == 0 {
		return int64(full * fullCnt)
	}
	sum := make([]int, len(flowers)+1)
	for i, ch := range flowers {
		sum[i+1] = sum[i] + ch
	}

	mx := flowers[len(flowers)-1]
	add := int(newFlowers)
	var dfs func(i, add int) int

	return int64(ans)
}

// 如果一个花园有 至少 target 朵花，那么这个花园称为 完善的 ，花园的 总美丽值 为以下分数之 和 ：
// 完善 花园数目乘以 full.
// 剩余 不完善 花园里，花的 最少数目 乘以 partial 。如果没有不完善花园，那么这一部分的值为 0 。
// 增加add，求最小值的最大值
func help(nums []int, pre []int, mx, add int) int {
	// 最大的最小值是 mid,能不能行
	check := func(i int) bool {
		j := sort.SearchInts(nums, i+1) - 1
		if add+pre[j+1] >= (j+1)*i {
			return true
		}
		return false
	}

	le, ri := 0, mx+1
	for le < ri {
		mid := le + (ri-le+1)/2
		if mid < mx+1 && check(mid) {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	return le
}
