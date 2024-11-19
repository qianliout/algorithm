package main

func main() {

}

func stoneGameVIII(stones []int) int {
	n := len(stones)
	sum := make([]int, n+1)
	for i, ch := range stones {
		sum[i+1] = sum[i] + ch
	}
	mem := make(map[int]int)
	// 初值，这里是最容易出错的
	mem[n] = sum[n]
	var dfs func(i int) int
	dfs = func(i int) int {
		if v, ok := mem[i]; ok {
			return v
		}
		// 这个判断可有可无
		if i > n {
			return 0
		}
		ans := max(dfs(i+1), sum[i]-dfs(i+1))
		mem[i] = ans
		return ans
	}
	ans := dfs(2)
	return ans
}
