package main

func main() {

}

func rob(nums []int) int {
	n := len(nums)
	var dfs func(i int, pre int) int
	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, 2)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i int, pre int) int {
		if i < 0 {
			return 0
		}
		if mem[i][pre] != -1 {
			return mem[i][pre]
		}
		ans := 0

		if pre == 1 {
			ans += dfs(i-1, 0)
		} else {
			a := dfs(i-1, 1) + nums[i]
			b := dfs(i-1, 0)
			ans = ans + max(a, b)
		}
		mem[i][pre] = ans
		return ans
	}
	return dfs(n-1, 0)
}
