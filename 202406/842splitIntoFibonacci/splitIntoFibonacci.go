package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(splitIntoFibonacci("1101111"))
	fmt.Println(splitIntoFibonacci("539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511"))
	fmt.Println(splitIntoFibonacci("121474836462147483647"))
}

func splitIntoFibonacci(num string) []int {
	nums := make([]int, len(num))
	for i, ch := range num {
		nums[i] = int(ch - '0')
	}
	ans := make([]int, 0)
	dfs(nums, 0, []int{}, &ans)
	return ans
}

func dfs(nums []int, start int, path []int, ans *[]int) {
	if len(*ans) > 0 {
		return
	}
	if start >= len(nums) {
		if len(path) >= 3 {
			*ans = append(*ans, path...)
		}
		return
	}

	for i := start; i < len(nums); i++ {
		// 不能有多个前导0
		if nums[start] == 0 && i > start {
			continue
		}
		cur := gen(nums, start, i)

		if len(path) >= 2 && !check(path, cur) {
			continue
		}
		if cur > math.MaxInt32 {
			continue
		}
		path = append(path, cur)
		dfs(nums, i+1, path, ans)
		path = path[:len(path)-1]
	}
}

func gen(nums []int, start, end int) int {
	ans := 0
	for i := start; i <= end; i++ {
		ans = ans*10 + nums[i]
	}
	return ans
}

func check(path []int, thr int) bool {
	fir, sec := path[len(path)-2], path[len(path)-1]
	if fir+sec != thr {
		return false
	}
	return true
}
