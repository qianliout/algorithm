package main

import (
	"fmt"
)

func main() {
	fmt.Println(distributeCandies(5, 2))
	fmt.Println(distributeCandies(3, 3))
}

// 数据量小，直接枚举
func distributeCandies1(n int, limit int) int {
	ans := 0
	for i := 0; i <= limit; i++ {
		for j := 0; j <= limit; j++ {
			if i+j > n {
				break
			}
			if n-i-j <= limit {
				ans++
			}
		}
	}
	return ans
}

/*
枚举第一个小朋友分得 xxx 颗糖果，那么还剩下 n−xn-xn−x 颗糖果，此时有两种情况：

	n−x>limit×2，至少有一个小朋友会分得大于 limit 颗糖果，此时不存在合法方案。

	n−x≤limit×2,对于第二个小朋友来说，至少得分得 max(0,n−x−limit) 颗糖果，才能保证第三个小朋友分得的糖果不超过 limit颗。同时至多能拿到 min(limit,n−x)颗糖果。

对于第二种情况计算出所有的合法方案即可。
*/
func distributeCandies(n int, limit int) int {
	// 先保证第一个小朋友是正确的
	ans := 0
	for i := 0; i <= min(limit, n); i++ {
		if n-i > 2*limit {
			continue
		}
		// 最后加1是指第一个小朋友分得 i 个糖果
		// min(n-i,limit)是指第二小朋友可以的分法
		ans += min(n-i, limit) - max(0, n-i-limit) + 1
	}
	return ans
}
