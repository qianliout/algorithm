package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestTimeFromDigits([]int{1, 2, 3, 4}))
}

func largestTimeFromDigits(arr []int) string {
	n := len(arr)
	visit := make([]bool, n)
	path := make([]int, 0)
	mx := make([]int, 0)
	dfs(arr, visit, path, &mx)
	if len(mx) == 0 {
		return ""
	}
	return gen(mx)
}

func dfs(nums []int, visit []bool, path []int, mx *[]int) {
	if len(path) == len(nums) {
		if check(path) {
			if len(*mx) == 0 || compare(path, *mx) {
				*mx = (*mx)[:0]
				*mx = append(*mx, path...)
			}
		}
		return
	}

	for i := 0; i < len(nums); i++ {
		if visit[i] {
			continue
		}
		path = append(path, nums[i])
		visit[i] = true
		dfs(nums, visit, path, mx)
		path = path[:len(path)-1]
		visit[i] = false
	}
}

func compare(a, b []int) bool {
	i := 0
	for i < len(a) && i < len(b) {
		if a[i] > b[i] {
			return true
		} else {
			if a[i] < b[i] {
				return false
			}
		}
		i++
	}
	return false
}

func gen(a []int) string {
	return fmt.Sprintf("%d%d:%d%d", a[0], a[1], a[2], a[3])
}

func check(a []int) bool {
	if a[0]*10+a[1] >= 24 {
		return false
	}

	if a[2]*10+a[3] >= 60 {
		return false
	}

	return true
}
