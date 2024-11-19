package main

func main() {

}

func permute2(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	path := make([]int, n)
	var dfs func(i int)
	used := make([]bool, n)
	dfs = func(i int) {
		if i >= n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		// 枚举选那一个的思想
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
