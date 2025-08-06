package main

func main() {

}

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	var dfs func(i int, sum int, path []int)
	dfs = func(i int, sum int, path []int) {
		if sum == n && len(path) == k {
			ans = append(ans, append([]int{}, path...))
			return
		}
		if sum > n || len(path) > k {
			return
		}
		for j := i; j <= 9; j++ {
			path = append(path, j)
			dfs(j+1, sum+j, path)
			path = path[:len(path)-1]
		}
	}
	dfs(1, 0, []int{})
	return ans
}
