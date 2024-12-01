package main

import (
	"fmt"
)

func main() {
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5))
}

func canReach1(arr []int, start int) bool {
	n := len(arr)
	q := []int{start}
	used := make([]bool, n)
	used[start] = true
	for len(q) > 0 {
		no := q[0]
		q = q[1:]
		nex := []int{no - arr[no], no + arr[no]}

		for _, i := range nex {
			if i >= n || i < 0 {
				continue
			}
			if arr[i] == 0 {
				return true
			}
			if !used[i] {
				q = append(q, i)
				used[i] = true
			}
		}
	}

	return false
}

func canReach(arr []int, start int) bool {
	n := len(arr)

	used := make([]bool, n)
	var dfs func(i int) bool
	dfs = func(i int) bool {
		if arr[i] == 0 {
			return true
		}
		nex := []int{i + arr[i], i - arr[i]}
		ans := false

		used[i] = true
		for _, j := range nex {
			if j < 0 || j >= n {
				continue
			}
			if used[j] {
				continue
			}
			ans = ans || dfs(j)
		}
		return ans
	}

	ans := dfs(start)

	return ans
}
