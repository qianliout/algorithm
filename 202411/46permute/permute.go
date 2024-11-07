package main

func main() {

}

func permute2(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	var dfs func(path []int)
	visit := make([]bool, n)
	dfs = func(path []int) {
		if len(path) == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < n; i++ {
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
	return ans
}

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	var dfs func(j int)
	visit := make([]bool, n)
	path := make([]int, n) // 直接覆盖的方式
	dfs = func(j int) {
		if j == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < n; i++ {
			if visit[i] {
				continue
			}
			visit[i] = true
			path[j] = nums[i]
			dfs(j + 1) // 去更新下一个数
			visit[i] = false
		}
	}
	dfs(0)
	return ans
}
