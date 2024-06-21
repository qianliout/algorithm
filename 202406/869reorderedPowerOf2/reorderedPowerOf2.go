package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reorderedPowerOf2(10))
	fmt.Println(reorderedPowerOf2(8))
	fmt.Println(reorderedPowerOf2(16))
	fmt.Println(reorderedPowerOf2(1024))
}

func reorderedPowerOf2(n int) bool {
	table := find(int(math.Pow10(9) + 10))
	nums := split(n)
	var res bool
	var dfs func(path []int)
	visit := make([]bool, len(nums))
	dfs = func(path []int) {
		if res {
			return
		}
		if len(path) == len(nums) {
			_, yes := cal(path, table)
			if yes {
				res = true
			}
			return
		}

		for i := 0; i < len(nums); i++ {
			if visit[i] {
				continue
			}
			visit[i] = true
			path = append(path, nums[i])
			dfs(path)
			path = path[:len(path)-1]
			visit[i] = false
		}
	}
	dfs([]int{})
	return res
}

func cal(path []int, table map[int]bool) (int, bool) {
	if len(path) == 0 || path[0] == 0 {
		return 0, false
	}
	ans := 0
	for i := 0; i < len(path); i++ {
		ans = ans*10 + path[i]
	}
	if table[ans] {
		return ans, true
	}
	return ans, false
}

func split(n int) []int {
	ans := make([]int, 0)
	for n > 0 {
		ans = append(ans, n%10)
		n = n / 10
	}
	return ans
}

// 打表，找 n 以内2的幂数
func find(n int) map[int]bool {
	res := make(map[int]bool)
	for i := 1; i <= n; i = i * 2 {
		res[i] = true
	}
	return res
}
