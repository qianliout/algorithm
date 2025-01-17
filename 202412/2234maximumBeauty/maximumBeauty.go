package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(maximumBeauty([]int{1, 3, 1, 1}, 7, 6, 12, 1)) // 14
	fmt.Println(maximumBeauty([]int{2, 4, 5, 3}, 10, 5, 2, 6)) // 30
}

func maximumBeauty(flowers []int, newFlowers int64, target, full, partial int) int64 {
	sort.Ints(flowers)
	n := len(flowers)
	if flowers[0] >= target { // 剪枝，此时所有花园都是完善的
		return int64(n * full)
	}

	leftFlowers := int(newFlowers) - target*n // 填充后缀后，剩余可以种植的花
	for i, f := range flowers {
		flowers[i] = min(f, target) // 去掉多余的花
		leftFlowers += flowers[i]   // 补上已有的花
	}

	ans := 0
	for i, x, sumFlowers := 0, 0, 0; i <= n; i++ { // 枚举后缀长度 n-i
		if leftFlowers >= 0 {
			// 计算最长前缀的长度
			for ; x < i && flowers[x]*x-sumFlowers <= leftFlowers; x++ {
				sumFlowers += flowers[x] // 注意 x 只增不减，二重循环的时间复杂度为 O(n)
			}
			beauty := (n - i) * full // 计算总美丽值
			if x > 0 {
				beauty += min((leftFlowers+sumFlowers)/x, target-1) * partial
			}
			ans = max(ans, beauty)
		}
		if i < n {
			leftFlowers += target - flowers[i]
		}
	}
	return int64(ans)
}
