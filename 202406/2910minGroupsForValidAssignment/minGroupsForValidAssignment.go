package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minGroupsForValidAssignment([]int{10, 10, 10, 3, 1, 1}))
}

func minGroupsForValidAssignment(balls []int) int {
	cnt := make(map[int]int)
	mn := math.MaxInt
	for _, ch := range balls {
		cnt[ch]++
	}
	for _, v := range cnt {
		mn = min(mn, v)
	}
	for k := mn; k >= 1; k-- {
		ans := 0
		for _, fre := range cnt {
			// 最少的桶里放 k 个元素，其他的桶可以放 k+1个素
			q := fre / k // 对于元素元素个数是 fre 的元素，可以有 q个桶
			r := fre % k // 也就是说剩下的有r 个元素，这个元素可以放在上面的 q 个桶里
			if q < r {
				ans = 0
				break // 尝试更小的一个 k
			}

			// 如果里面的 for 是break退出的，那么ans 会被重新等于0
			ans += int(math.Ceil(float64(fre) / float64(k+1)))
			// 这样写更显功力
			// ans += (fre + k) / (k + 1)
		}
		if ans > 0 {
			return ans
		}
	}
	return len(balls)
}
