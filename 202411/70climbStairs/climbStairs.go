package main

func main() {

}

func climbStairs(n int) int {
	mem := make([]int, n)
	for i := range mem {
		mem[i] = -1
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i <= 2 {
			return i
		}
		if mem[i] != -1 {
			return mem[i]
		}
		a := dfs(i-1) + dfs(i-2)
		mem[i] = a
		return a
	}
	return dfs(n)
}
