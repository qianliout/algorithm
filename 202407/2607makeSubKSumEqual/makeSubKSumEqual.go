package main

import (
	"sort"
)

func main() {

}

func makeSubKSumEqual(arr []int, k int) int64 {
	n := len(arr)
	k = gcd(n, k)
	// 分组,裴蜀定理 这里是难点
	g := make([][]int, k)
	for i := 0; i < n; i++ {
		g[i%k] = append(g[i%k], arr[i])
	}

	ans := 0
	for i := 0; i < k; i++ {
		nums := g[i]
		sort.Ints(nums)
		ans += cal(nums)
	}
	return int64(ans)
}

func cal(nums []int) int {
	// 中位数贪心
	mid := nums[len(nums)/2]
	ans := 0
	for _, ch := range nums {
		ans += abs(ch - mid)
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
