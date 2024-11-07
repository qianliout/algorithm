package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [-1,-1,2],[-1,0,1]]
}

// 会超时
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	n := len(nums)
	var dfs func(path []int, j int, s int)
	visit := make([]bool, n)
	dfs = func(path []int, j int, s int) {
		if len(path) == 3 && s == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := j; i < n; i++ {
			if visit[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !visit[i-1] {
				continue
			}
			path = append(path, nums[i])
			visit[i] = true
			s += nums[i]
			dfs(path, i+1, s)
			path = path[:len(path)-1]
			s -= nums[i]
			visit[i] = false
		}
	}
	dfs([]int{}, 0, 0)
	return ans
}

func threeSum3(nums []int) [][]int {
	sort.Ints(nums)
	cnt := make(map[int][]int)
	for i, ch := range nums {
		cnt[ch] = append(cnt[ch], i)
	}
	n := len(nums)
	ans := make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			s := nums[i] + nums[j]
			next := cnt[-s]
			for _, k := range next {
				if k > j {
					ans = append(ans, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return ans
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	i := 0
	ans := make([][]int, 0)
	for i < n {
		if i > 0 && nums[i] == nums[i-1] {
			i++
			continue
		}
		left, right := i+1, n-1
		for left < right {
			s := nums[i] + nums[left] + nums[right]
			if s == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if s > 0 {
				right--
			} else if s < 0 {
				left++
			}
		}
		i++
	}
	return ans
}
