package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimizeMax([]int{10, 1, 2, 7, 1, 3}, 2))
	fmt.Println(minimizeMax([]int{4, 2, 1, 2}, 1))
	fmt.Println(minimizeMax([]int{3, 4, 2, 3, 2, 1, 2}, 3))
}

func minimizeMax1(nums []int, p int) int {
	sort.Ints(nums)
	var check func(mx int) bool

	var cal func(start int, mx int) int

	n := len(nums)
	mem := make([]int, n)
	// 这里是用动态规划的做法
	cal = func(start int, mx int) int {
		// 选start
		if start >= len(nums)-1 {
			return 0
		}
		if mem[start] >= 0 {
			return mem[start]
		}
		res1 := 0
		if nums[start+1]-nums[start] <= mx {
			res1++
		}
		res1 += cal(start+2, mx)
		res2 := cal(start+1, mx)

		a := max(res1, res2)
		mem[start] = a
		return a
	}

	check = func(mx int) bool {
		mem = make([]int, n)
		for i := 0; i < n; i++ {
			mem[i] = -1
		}
		cnt := cal(0, mx)
		return cnt >= p
	}
	mi := 0
	mx := nums[n-1] - nums[0] + 1
	le, ri := mi, mx
	for le < ri {
		// 左端点
		mid := le + (ri-le)/2
		if mid >= mi && mid < mx && check(mid) {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return le
}

func count(nums []int, start int, mx int) int {
	// 选start
	if start >= len(nums)-1 {
		return 0
	}
	res1 := 0
	if nums[start+1]-nums[start] <= mx {
		res1++
	}
	res1 += count(nums, start+2, mx)
	res2 := count(nums, start+1, mx)

	return max(res1, res2)
}

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	var check func(mx int) bool

	var cal func(start int, mx int) int

	n := len(nums)
	mem := make([]int, n)
	//  贪心的做法，第一个可选，那就直接选
	cal = func(start int, mx int) int {
		res := 0
		i := 0
		for i < n-1 {
			if nums[i+1]-nums[i] <= mx {
				res++
				i += 2
			} else {
				i++
			}
		}
		return res
	}

	check = func(mx int) bool {
		mem = make([]int, n)
		for i := 0; i < n; i++ {
			mem[i] = -1
		}
		cnt := cal(0, mx)
		return cnt >= p
	}
	mi := 0
	mx := nums[n-1] - nums[0] + 1
	le, ri := mi, mx
	for le < ri {
		// 左端点
		mid := le + (ri-le)/2
		if mid >= mi && mid < mx && check(mid) {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return le
}
