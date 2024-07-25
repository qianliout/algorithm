package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoEggDrop(100))
}

func twoEggDrop(n int) int {
	dp1 := make([]int, n+1) // 表示有 1个蛋，验证 j 层楼需要的操作交数
	dp2 := make([]int, n+1) // 表示有 2个蛋，验证 j 层楼需要的操作交数
	for j := 0; j <= n; j++ {
		dp1[j] = j
	}
	for j := 1; j <= n; j++ {
		dp2[j] = n + 1 // 初值:也是最大值，下面会取最小值
		for k := 1; k <= j; k++ {
			// 假如一个鸡蛋在 k 楼碎了，那么在0到 k-1只能是一层一层的试了，k到 n 层用两个鸡蛋二分相似的方式
			dp2[j] = min(dp2[j], max(dp1[k-1]+1, dp2[j-k]+1))
		}
	}
	return dp2[n]
}
