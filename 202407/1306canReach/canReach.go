package main

func main() {

}

func canReach(arr []int, start int) bool {
	n := len(arr)
	var dfs func(i int) bool
	visit := make([]bool, n)
	dfs = func(i int) bool {
		if i < 0 || i >= n {
			return false
		}
		if arr[i] == 0 {
			return true
		}
		if visit[i] {
			return false
		}

		visit[i] = true
		return dfs(i+arr[i]) || dfs(i-arr[i])
	}
	return dfs(start)
}
