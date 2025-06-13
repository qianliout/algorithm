package main

func main() {

}

func permute2(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	path := make([]int, n)
	used := make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i < 0 || i >= n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for j := 0; j < n; j++ {
			if used[j] {
				continue
			}
			used[j] = true
			path[i] = nums[j]
			dfs(i + 1)
			used[j] = false
		}
	}
	dfs(0)
	return ans
}

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	path := make([]int, n)
	used := make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i < 0 || i >= n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for j := 0; j < n; j++ {
			if used[j] {
				continue
			}
			used[j] = true
			path[i] = nums[j]
			dfs(i + 1)
			used[j] = false
		}
	}
	dfs(0)
	return ans
}

// nums 中的所有整数 互不相同
